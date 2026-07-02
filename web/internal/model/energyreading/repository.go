package energyreading

import (
	"context"
	"energy-monitor-server/internal/model/appliance"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type readingRepo struct {
	db *sqlx.DB
}

type ReadingRepository interface {
	Create(ctx context.Context, r *EnergyReading) error
	List(ctx context.Context, userID int64, applianceID *int64, from, to *time.Time) ([]EnergyReading, error)

	GetSummary(
		ctx context.Context,
		userID int64,
		applianceID *int64,
		rangeType string,
		month, year *int,
	) (*ReadingSummary, error)

	UpdateApplianceLastReading(
		ctx context.Context,
		applianceID int64,
	) error

	GetEnergyChart(
		ctx context.Context,
		userID int64,
		rangeType string,
		month, year *int,
	) ([]ChartPoint, error)

	GetAnalyticsEnergyChart(
		ctx context.Context,
		userID int64,
		applianceID *int64,
		rangeType string,
		month, year *int,
	) ([]ChartPoint, error)

	GetVoltageCurrentChart(
		ctx context.Context,
		userID int64,
		applianceID *int64,
		rangeType string,
		month, year *int,
	) ([]VoltageCurrentPoint, error)

	GetAnalyticsSummary(
		ctx context.Context,
		userID int64,
		applianceID *int64,
		rangeType string,
		month, year *int,
	) (*AnalyticsSummary, error)

	GetDetailedReadings(
		ctx context.Context,
		userID int64,
		applianceID *int64,
		rangeType string,
		month, year *int,
		limit int,
		offset int,
	) ([]EnergyReading, int, error)
}

func NewReadingRepository(db *sqlx.DB) ReadingRepository {
	return &readingRepo{db: db}
}

func (r *readingRepo) Create(ctx context.Context, reading *EnergyReading) error {
	query := `
		INSERT INTO energy_readings (
			appliance_id,
			ts,
			voltage,
			current,
			power,
			energy_kwh,
			frequency_hz
		)
		VALUES ($1, NOW(), $2, $3, $4, $5, $6)
		RETURNING id, ts
	`

	return r.db.QueryRowxContext(ctx, query,
		reading.ApplianceID,
		reading.Voltage,
		reading.Current,
		reading.Power,
		reading.EnergyKWh,
		reading.FrequencyHz,
	).Scan(&reading.ID, &reading.Timestamp)
}

func (r *readingRepo) List(
	ctx context.Context,
	userID int64,
	applianceID *int64,
	from, to *time.Time,
) ([]EnergyReading, error) {
	query := `
		SELECT er.*
		FROM energy_readings er
		JOIN appliances a ON er.appliance_id = a.id
		WHERE a.user_id = $1
	`
	args := []any{userID}
	i := 2

	if applianceID != nil {
		query += fmt.Sprintf(" AND er.appliance_id = $%d", i)
		args = append(args, *applianceID)
		i++
	}

	if from != nil {
		query += fmt.Sprintf(" AND er.timestamp >= $%d", i)
		args = append(args, *from)
		i++
	}

	if to != nil {
		query += fmt.Sprintf(" AND er.timestamp <= $%d", i)
		args = append(args, *to)
		i++
	}

	query += " ORDER BY er.timestamp DESC LIMIT 500"

	var readings []EnergyReading
	err := r.db.SelectContext(ctx, &readings, query, args...)
	return readings, err
}

func (r *readingRepo) GetSummary(
	ctx context.Context,
	userID int64,
	applianceID *int64,
	rangeType string,
	month, year *int,
) (*ReadingSummary, error) {

	args := []any{userID}
	nextParam := 2

	applianceFilter := ""
	if applianceID != nil {
		applianceFilter = fmt.Sprintf("AND er.appliance_id = $%d", nextParam)
		args = append(args, *applianceID)
		nextParam++
	}

	condition, rangeArgs, err := buildRangeCondition(
		rangeType,
		month,
		year,
		nextParam,
	)
	if err != nil {
		return nil, err
	}

	args = append(args, rangeArgs...)

	query := fmt.Sprintf(`
		SELECT
			COALESCE(SUM(
				(LEAST(EXTRACT(EPOCH FROM (ts - prev_ts)), 5) * power) / 3600000.0
			), 0) AS total_energy_kwh,

			COALESCE(MAX(power), 0) AS peak_power,

			(
				SELECT COUNT(*)
				FROM appliances
				WHERE user_id = $1 AND status = 'online'
			) AS active_devices,

			(
				SELECT COUNT(*)
				FROM appliances WHERE user_id = $1
			) AS device_count,

			(
				SELECT rate_per_kwh
				FROM settings
				WHERE user_id = $1
			) AS billing_rate,

			(
				SELECT COUNT(*)
				FROM alerts al
				JOIN appliances a ON a.id = al.appliance_id
				WHERE a.user_id = $1 AND al.resolved_at IS NULL
			) AS active_alerts

		FROM (
			SELECT
				er.ts,
				er.power,
				LAG(er.ts) OVER (
					PARTITION BY er.appliance_id
					ORDER BY er.ts
				) AS prev_ts
			FROM energy_readings er
			JOIN appliances a ON a.id = er.appliance_id
			WHERE
				a.user_id = $1
				AND %s
				%s
		) t
		WHERE prev_ts IS NOT NULL;
	`, condition, applianceFilter)

	var summary ReadingSummary

	err = r.db.QueryRowContext(ctx, query, args...).Scan(
		&summary.TotalEnergyKWh,
		&summary.PeakPower,
		&summary.ActiveDevices,
		&summary.DeviceCount,
		&summary.BillingRate,
		&summary.ActiveAlerts,
	)

	if err != nil {
		return nil, err
	}

	return &summary, nil
}
func (r *readingRepo) UpdateApplianceLastReading(
	ctx context.Context,
	applianceID int64,
) error {
	query := `
		UPDATE appliances
		SET
			last_reading = NOW(),
			status = $1,
			updated_at = NOW()
		WHERE id = $2
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		appliance.ApplianceStatusOnline,
		applianceID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return appliance.ErrApplianceNotFound
	}

	return nil
}

func (r *readingRepo) GetEnergyChart(
	ctx context.Context,
	userID int64,
	rangeType string,
	month, year *int,
) ([]ChartPoint, error) {

	monthStart := time.Now()

	if month != nil {
		m := time.Month(*month)

		y := monthStart.Year()
		if year != nil {
			y = *year
		}

		monthStart = time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	} else {
		monthStart = time.Date(
			monthStart.Year(),
			monthStart.Month(),
			1,
			0, 0, 0, 0,
			time.UTC,
		)
	}

	var query string

	switch rangeType {

	case "today":
		query = `
		WITH hours AS (
			SELECT generate_series(0, 23) AS hour
		),
		energy AS (
			SELECT
				EXTRACT(HOUR FROM t.ts)::int AS hour,
				SUM(
					(LEAST(EXTRACT(EPOCH FROM (t.ts - t.prev_ts)), 5) * t.power) / 3600000.0
				) AS value
			FROM (
				SELECT
					er.ts,
					er.power,
					LAG(er.ts) OVER (
						PARTITION BY er.appliance_id
						ORDER BY er.ts
					) AS prev_ts
				FROM energy_readings er
				JOIN appliances a ON a.id = er.appliance_id
				WHERE
					a.user_id = $1
					AND er.ts >= date_trunc('day', NOW())
			) t
			WHERE prev_ts IS NOT NULL
			GROUP BY hour
		)

		SELECT
			LPAD(hours.hour::text, 2, '0') || ':00' AS label,
			COALESCE(e.value, 0) AS value
		FROM hours
		LEFT JOIN energy e ON e.hour = hours.hour
		ORDER BY hours.hour;
		`

	case "7d":
		query = `
			WITH days AS (
				SELECT generate_series(
					NOW() - INTERVAL '6 days',
					NOW(),
					INTERVAL '1 day'
				)::date AS day
			),
			energy AS (
				SELECT
					DATE(t.ts) AS day,
					SUM(
						(LEAST(EXTRACT(EPOCH FROM (t.ts - t.prev_ts)), 5) * t.power) / 3600000.0
					) AS value
				FROM (
					SELECT
						er.ts,
						er.power,
						LAG(er.ts) OVER (
							PARTITION BY er.appliance_id
							ORDER BY er.ts
						) AS prev_ts
					FROM energy_readings er
					JOIN appliances a ON a.id = er.appliance_id
					WHERE
						a.user_id = $1
						AND er.ts >= NOW() - INTERVAL '7 days'
				) t
				WHERE prev_ts IS NOT NULL
				GROUP BY day
			)
			SELECT
				TO_CHAR(d.day, 'MM-DD') AS label,
				COALESCE(e.value, 0) AS value
			FROM days d
			LEFT JOIN energy e ON e.day = d.day
			ORDER BY d.day;
		`

	default: // month
		query = `
			WITH days AS (
				SELECT generate_series(
					date_trunc('month', $2::timestamp),
					date_trunc('month', $2::timestamp) + INTERVAL '1 month - 1 day',
					INTERVAL '1 day'
				)::date AS day
			),
			energy AS (
				SELECT
					DATE(t.ts) AS day,
					SUM(
						(LEAST(EXTRACT(EPOCH FROM (t.ts - t.prev_ts)), 5) * t.power)
						/ 3600000.0
					) AS value
				FROM (
					SELECT
						er.ts,
						er.power,
						LAG(er.ts) OVER (
							PARTITION BY er.appliance_id
							ORDER BY er.ts
						) AS prev_ts
					FROM energy_readings er
					JOIN appliances a ON a.id = er.appliance_id
					WHERE
						a.user_id = $1
						AND er.ts >= date_trunc('month', $2::timestamp)
						AND er.ts < date_trunc('month', $2::timestamp) + INTERVAL '1 month'
				) t
				WHERE prev_ts IS NOT NULL
				GROUP BY day
			)

			SELECT
				TO_CHAR(d.day, 'DD') AS label,
				COALESCE(e.value, 0) AS value
			FROM days d
			LEFT JOIN energy e ON e.day = d.day
			ORDER BY d.day;
		`
	}

	rows, err := r.db.QueryContext(ctx, query, userID, monthStart)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []ChartPoint

	for rows.Next() {
		var p ChartPoint
		if err := rows.Scan(&p.Label, &p.Value); err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}

func buildRangeQuery(rangeType string) (interval string, labelFormat string) {
	switch rangeType {
	case "today":
		return "INTERVAL '1 day'", "HH24:00"
	case "7d":
		return "INTERVAL '7 days'", "YYYY-MM-DD"
	case "month":
		return "INTERVAL '30 days'", "YYYY-MM-DD"
	default:
		return "INTERVAL '1 day'", "HH24:00"
	}
}

func buildRangeCondition(
	rangeType string,
	month, year *int,
	startParam int,
) (string, []any, error) {

	switch rangeType {
	case "today":
		return "er.ts >= date_trunc('day', NOW())", nil, nil

	case "7d":
		return "er.ts >= NOW() - INTERVAL '7 days'", nil, nil

	case "month":
		now := time.Now()

		m := int(now.Month())
		if month != nil {
			m = *month
		}

		y := now.Year()
		if year != nil {
			y = *year
		}

		if m < 1 || m > 12 {
			return "", nil, fmt.Errorf("invalid month")
		}

		start := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
		end := start.AddDate(0, 1, 0)

		return fmt.Sprintf(
			"er.ts >= $%d AND er.ts < $%d",
			startParam,
			startParam+1,
		), []any{start, end}, nil

	default:
		return "", nil, fmt.Errorf("invalid range type")
	}
}

func (r *readingRepo) GetAnalyticsEnergyChart(
	ctx context.Context,
	userID int64,
	applianceID *int64,
	rangeType string,
	month, year *int,
) ([]ChartPoint, error) {

	groupUnit := groupByUnit(rangeType)
	labelFormat := buildLabelFormat(rangeType)

	args := []any{userID}
	nextParam := 2

	applianceFilter := ""
	if applianceID != nil {
		applianceFilter = fmt.Sprintf("AND er.appliance_id = $%d", nextParam)
		args = append(args, *applianceID)
		nextParam++
	}

	condition, rangeArgs, err := buildRangeCondition(
		rangeType,
		month,
		year,
		nextParam,
	)
	if err != nil {
		return nil, err
	}

	args = append(args, rangeArgs...)

	query := fmt.Sprintf(`
		SELECT
			TO_CHAR(group_ts, '%s') AS label,
			SUM(
				(LEAST(EXTRACT(EPOCH FROM (t.ts - t.prev_ts)), 5) * t.power) / 3600000.0
			) AS value
		FROM (
			SELECT
				DATE_TRUNC('%s', er.ts) AS group_ts,
				er.ts,
				er.power,
				LAG(er.ts) OVER (
					PARTITION BY er.appliance_id
					ORDER BY er.ts
				) AS prev_ts
			FROM energy_readings er
			JOIN appliances a ON a.id = er.appliance_id
			WHERE
				a.user_id = $1
				AND %s
				%s
		) t
		WHERE t.prev_ts IS NOT NULL
		GROUP BY group_ts
		ORDER BY group_ts;
	`, labelFormat, groupUnit, condition, applianceFilter)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []ChartPoint

	for rows.Next() {
		var p ChartPoint
		if err := rows.Scan(&p.Label, &p.Value); err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}
func (r *readingRepo) GetVoltageCurrentChart(
	ctx context.Context,
	userID int64,
	applianceID *int64,
	rangeType string,
	month, year *int,
) ([]VoltageCurrentPoint, error) {

	groupUnit := groupByUnit(rangeType)
	labelFormat := buildLabelFormat(rangeType)

	args := []any{userID}
	nextParam := 2

	applianceFilter := ""
	if applianceID != nil {
		applianceFilter = fmt.Sprintf("AND er.appliance_id = $%d", nextParam)
		args = append(args, *applianceID)
		nextParam++
	}

	condition, rangeArgs, err := buildRangeCondition(
		rangeType,
		month,
		year,
		nextParam,
	)
	if err != nil {
		return nil, err
	}

	args = append(args, rangeArgs...)

	query := fmt.Sprintf(`
		SELECT
			TO_CHAR(group_ts, '%s') AS label,
			AVG(voltage) AS voltage,
			AVG(current) AS current
		FROM (
			SELECT
				DATE_TRUNC('%s', er.ts) AS group_ts,
				er.voltage,
				er.current
			FROM energy_readings er
			JOIN appliances a ON a.id = er.appliance_id
			WHERE
				a.user_id = $1
				AND %s
				%s
		) t
		GROUP BY group_ts
		ORDER BY group_ts;
	`, labelFormat, groupUnit, condition, applianceFilter)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []VoltageCurrentPoint

	for rows.Next() {
		var p VoltageCurrentPoint
		if err := rows.Scan(&p.Label, &p.Voltage, &p.Current); err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}
func (r *readingRepo) GetAnalyticsSummary(
	ctx context.Context,
	userID int64,
	applianceID *int64,
	rangeType string,
	month, year *int,
) (*AnalyticsSummary, error) {

	args := []any{userID}
	nextParam := 2

	applianceFilter := ""
	if applianceID != nil {
		applianceFilter = fmt.Sprintf("AND er.appliance_id = $%d", nextParam)
		args = append(args, *applianceID)
		nextParam++
	}

	condition, rangeArgs, err := buildRangeCondition(
		rangeType,
		month,
		year,
		nextParam,
	)
	if err != nil {
		return nil, err
	}

	args = append(args, rangeArgs...)

	query := fmt.Sprintf(`
		SELECT
			COALESCE(
				SUM(t.power * t.dt) / NULLIF(SUM(t.dt), 0),
				0
			) AS avg_power,

			COALESCE(AVG(t.voltage), 0) AS avg_voltage,
			COALESCE(AVG(t.current), 0) AS avg_current,

			COALESCE(MAX(t.power), 0) AS peak_power

		FROM (
			SELECT
				er.power,
				er.voltage,
				er.current,

				LEAST(
					EXTRACT(EPOCH FROM (
						er.ts - LAG(er.ts) OVER (
							PARTITION BY er.appliance_id
							ORDER BY er.ts
						)
					)),
					5
				) AS dt

			FROM energy_readings er
			JOIN appliances a ON a.id = er.appliance_id
			WHERE
				a.user_id = $1
				AND %s
				%s
		) t
		WHERE t.dt IS NOT NULL
	`, condition, applianceFilter)

	var summary AnalyticsSummary

	err = r.db.QueryRowContext(ctx, query, args...).Scan(
		&summary.AvgPower,
		&summary.AvgVoltage,
		&summary.AvgCurrent,
		&summary.PeakPower,
	)
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *readingRepo) GetDetailedReadings(
	ctx context.Context,
	userID int64,
	applianceID *int64,
	rangeType string,
	month, year *int,
	limit int,
	offset int,
) ([]EnergyReading, int, error) {

	args := []any{userID}

	nextParam := 2

	applianceFilter := ""

	if applianceID != nil {
		applianceFilter = fmt.Sprintf("AND er.appliance_id = $%d", nextParam)
		args = append(args, *applianceID)
		nextParam++
	}

	condition, rangeArgs, err := buildRangeCondition(
		rangeType,
		month,
		year,
		nextParam,
	)

	if err != nil {
		return nil, 0, err
	}

	args = append(args, rangeArgs...)

	countQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM energy_readings er
		JOIN appliances a ON a.id = er.appliance_id
		WHERE
			a.user_id = $1
			AND %s
			%s
	`, condition, applianceFilter)

	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	dataQuery := fmt.Sprintf(`
		SELECT
			er.ts,
			er.voltage,
			er.current,
			er.power,
			er.energy_kwh
		FROM energy_readings er
		JOIN appliances a ON a.id = er.appliance_id
		WHERE
			a.user_id = $1
			AND %s
			%s
		ORDER BY er.ts DESC
		LIMIT $%d OFFSET $%d
	`, condition, applianceFilter, len(args)+1, len(args)+2)

	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var result []EnergyReading

	for rows.Next() {
		var r EnergyReading
		if err := rows.Scan(
			&r.Timestamp,
			&r.Voltage,
			&r.Current,
			&r.Power,
			&r.EnergyKWh,
		); err != nil {
			return nil, 0, err
		}
		result = append(result, r)
	}

	if result == nil {
		result = []EnergyReading{}
	}

	return result, total, nil
}

func groupByUnit(rangeType string) string {
	switch rangeType {
	case "today":
		return "hour"
	case "7d":
		return "day"
	case "month":
		return "day"
	default:
		return "hour"
	}
}

func buildLabelFormat(rangeType string) string {
	switch rangeType {
	case "today":
		return "HH24:00"
	case "7d":
		return "MM-DD"
	case "month":
		return "DD"
	default:
		return "HH24:00"
	}
}

func buildApplianceFilter(applianceID *int64, paramIndex int) string {
	if applianceID == nil {
		return ""
	}
	return fmt.Sprintf("AND er.appliance_id = $%d", paramIndex)
}
