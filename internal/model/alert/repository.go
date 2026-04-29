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
	args := []interface{}{userID}
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
			severity,
			message,
			triggered_at
		)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, triggered_at
	`

	return r.db.QueryRowxContext(ctx, query,
		a.ApplianceID,
		a.Type,
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
