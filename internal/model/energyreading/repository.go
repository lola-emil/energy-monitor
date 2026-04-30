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
		from,
		to *time.Time,
	) (*ReadingSummary, error)
	UpdateApplianceLastReading(
		ctx context.Context,
		applianceID int64,
	) error

	GetEnergyChart(
		ctx context.Context,
		userID int64,
		rangeType string, // "today", "7d", "month"
	) ([]ChartPoint, error)
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
	args := []interface{}{userID}
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
	from,
	to *time.Time,
) (*ReadingSummary, error) {
	query := `
		SELECT
			COALESCE(SUM(
				(EXTRACT(EPOCH FROM (ts - prev_ts)) * power) / 3600000.0
			), 0) AS total_energy_kwh,

			COALESCE(MAX(power), 0) AS peak_power,

			(
				SELECT COUNT(*)
				FROM appliances
				WHERE user_id = $1 AND status = 'online'
			) AS active_devices,

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
			WHERE a.user_id = $1
		) t
		WHERE prev_ts IS NOT NULL;
	`

	var summary ReadingSummary

	err := r.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&summary.TotalEnergyKWh,
		&summary.PeakPower,
		&summary.ActiveDevices,
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
) ([]ChartPoint, error) {

	query := `
		SELECT
			DATE(t.ts) AS label,
			SUM(
				(LEAST(EXTRACT(EPOCH FROM (t.ts - t.prev_ts)), 5) * t.power) / 3600000.0
			) AS value
		FROM (
			SELECT
				er.ts,
				er.power,
				er.appliance_id,
				LAG(er.ts) OVER (
					PARTITION BY er.appliance_id
					ORDER BY er.ts
				) AS prev_ts
			FROM energy_readings er
			JOIN appliances a ON a.id = er.appliance_id
			WHERE
				a.user_id = $1
				AND er.ts >= NOW() - INTERVAL '30 days'
		) t
		WHERE t.prev_ts IS NOT NULL
		GROUP BY DATE(t.ts)
		ORDER BY DATE(t.ts);
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
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
