package alert

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AlertFilter struct {
	Status      string // "active" or "resolved"
	Severity    string
	ApplianceID *int64
}

type alertRepo struct {
	db *sqlx.DB
}

type AlertRepository interface {
	List(ctx context.Context, userID int64, filter AlertFilter) ([]Alert, error)
	GetByID(ctx context.Context, userID, id int64) (*Alert, error)
	Resolve(ctx context.Context, userID, id int64) error

	Create(ctx context.Context, alert *Alert) error
	HasActiveAlert(
		ctx context.Context,
		applianceID int64,
		alertType AlertType,
	) (bool, error)

	HasRecentUnresolved(
		ctx context.Context,
		applianceID int64,
		alertType AlertType,
		cooldownMinutes int,
	) (bool, error)

	ResolveActiveAlert(
		ctx context.Context,
		applianceID int64,
		alertType AlertType,
	) error

	GetAnalyticsAlerts(
		ctx context.Context,
		userID int64,
		applianceID *int64,
		rangeType string,
		limit int,
	) ([]Alert, error)

	GetRecentByAppliance(
		ctx context.Context,
		userID int64,
		applianceID int64,
		limit int,
	) ([]Alert, error)
}

func NewAlertRepository(db *sqlx.DB) *alertRepo {
	return &alertRepo{db: db}
}

func (r *alertRepo) List(ctx context.Context, userID int64, f AlertFilter) ([]Alert, error) {
	query := `
		SELECT a.*
		FROM alerts a
		LEFT JOIN appliances ap ON a.appliance_id = ap.id
		WHERE ap.user_id = $1
	`
	args := []any{userID}
	i := 2

	switch f.Status {
	case "active":
		query += " AND a.resolved_at IS NULL"
	case "resolved":
		query += " AND a.resolved_at IS NOT NULL"
	}

	if f.Severity != "" {
		query += fmt.Sprintf(" AND a.severity = $%d", i)
		args = append(args, f.Severity)
		i++
	}

	if f.ApplianceID != nil {
		query += fmt.Sprintf(" AND a.appliance_id = $%d", i)
		args = append(args, *f.ApplianceID)
		i++
	}

	query += " ORDER BY a.triggered_at DESC"

	var alerts []Alert
	err := r.db.SelectContext(ctx, &alerts, query, args...)
	return alerts, err
}

func (r *alertRepo) GetByID(ctx context.Context, userID, id int64) (*Alert, error) {
	var a Alert

	err := r.db.GetContext(ctx, &a, `
		SELECT a.*
		FROM alerts a
		LEFT JOIN appliances ap ON a.appliance_id = ap.id
		WHERE a.id = $1 AND ap.user_id = $2
	`, id, userID)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *alertRepo) Resolve(ctx context.Context, userID, id int64) error {
	res, err := r.db.ExecContext(ctx, `
		UPDATE alerts
		SET resolved_at = NOW()
		WHERE id = $1
		AND appliance_id IN (
			SELECT id FROM appliances WHERE user_id = $2
		)
		AND resolved_at IS NULL
	`, id, userID)

	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *alertRepo) Create(ctx context.Context, a *Alert) error {
	query := `
		INSERT INTO alerts (
			appliance_id,
			type,
			name,
			severity,
			message,
			triggered_at
		)
		VALUES ($1, $2, $3, $4, $5, NOW())
		RETURNING id, triggered_at
	`

	return r.db.QueryRowxContext(ctx, query,
		a.ApplianceID,
		a.Type,
		a.Name,
		a.Severity,
		a.Message,
	).Scan(&a.ID, &a.TriggeredAt)
}

func (r *alertRepo) HasActiveAlert(
	ctx context.Context,
	applianceID int64,
	alertType AlertType,
) (bool, error) {
	var exists bool

	err := r.db.GetContext(ctx, &exists, `
		SELECT EXISTS (
			SELECT 1
			FROM alerts
			WHERE appliance_id = $1
			AND type = $2
			AND resolved_at IS NULL
		)
	`, applianceID, alertType)

	return exists, err
}

func (r *alertRepo) HasRecentUnresolved(
	ctx context.Context,
	applianceID int64,
	alertType AlertType,
	cooldownMinutes int,
) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM alerts
			WHERE
				appliance_id = $1
				AND type = $2
				AND resolved_at IS NULL
				AND triggered_at >= NOW() - ($3 * INTERVAL '1 minute')
		)
	`

	var exists bool

	err := r.db.QueryRowContext(
		ctx,
		query,
		applianceID,
		alertType,
		cooldownMinutes,
	).Scan(&exists)

	return exists, err
}

func (r *alertRepo) ResolveActiveAlert(
	ctx context.Context,
	applianceID int64,
	alertType AlertType,
) error {
	query := `
		UPDATE alerts
		SET
			resolved_at = NOW()
		WHERE
			appliance_id = $1
			AND type = $2
			AND resolved_at IS NULL
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		applianceID,
		alertType,
	)

	return err
}

func (r *alertRepo) GetAnalyticsAlerts(
	ctx context.Context,
	userID int64,
	applianceID *int64,
	rangeType string,
	limit int,
) ([]Alert, error) {

	condition := buildRangeCondition(rangeType)

	query := fmt.Sprintf(`
		SELECT
			al.id,
			al.message,
			al.severity,
			al.triggered_at,
			a.id
		FROM alerts al
		JOIN appliances a ON a.id = al.appliance_id
		JOIN energy_readings er ON a.id = er.appliance_id
		WHERE
			a.user_id = $1
			AND %s
			%s
		ORDER BY al.triggered_at DESC
		LIMIT $%d
	`, condition, buildApplianceFilter(applianceID, 2), 2+boolToInt(applianceID != nil))

	args := []any{userID}
	if applianceID != nil {
		args = append(args, *applianceID)
	}
	args = append(args, limit)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Alert

	for rows.Next() {
		var a Alert
		if err := rows.Scan(
			&a.ID,
			&a.Message,
			&a.Severity,
			&a.TriggeredAt,
			&a.ApplianceID,
		); err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	if result == nil {
		result = []Alert{}
	}

	return result, nil
}

func (r *alertRepo) GetRecentByAppliance(
	ctx context.Context,
	userID int64,
	applianceID int64,
	limit int,
) ([]Alert, error) {

	query := `
		SELECT
			al.id,
			al.message,
			al.severity,
			al.triggered_at
		FROM alerts al
		JOIN appliances a ON a.id = al.appliance_id
		WHERE
			a.user_id = $1
			AND a.id = $2
		ORDER BY al.triggered_at DESC
		LIMIT $3;
	`

	rows, err := r.db.QueryContext(ctx, query, userID, applianceID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Alert

	for rows.Next() {
		var a Alert
		if err := rows.Scan(
			&a.ID,
			&a.Message,
			&a.Severity,
			&a.TriggeredAt,
		); err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	if result == nil {
		result = []Alert{}
	}

	return result, nil
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

func buildRangeCondition(rangeType string) string {
	switch rangeType {
	case "today":
		return "er.ts >= date_trunc('day', NOW())"
	case "7d":
		return "er.ts >= NOW() - INTERVAL '7 days'"
	case "month":
		return "er.ts >= date_trunc('month', NOW())"
	default:
		return "er.ts >= date_trunc('day', NOW())"
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
