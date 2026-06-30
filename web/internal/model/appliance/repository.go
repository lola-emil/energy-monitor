package appliance

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	ErrApplianceNotFound = errors.New("appliance not found")
)

type applianceRepo struct {
	db *sqlx.DB
}

type ApplianceRepository interface {
	Create(ctx context.Context, a *Appliance) error
	GetByID(ctx context.Context, id int64) (*Appliance, error)
	List(ctx context.Context, userID int64) ([]Appliance, error)
	Update(ctx context.Context, a *Appliance) error
	Delete(ctx context.Context, userID, id int64) error

	UpdateLastReading(
		ctx context.Context,
		applianceID int64,
		lastReading time.Time,
	) error

	GetByDeviceCode(
		ctx context.Context,
		deviceCode string,
	) (*Appliance, error)

	GetOfflineCandidates(
		ctx context.Context,
		offlineMinutes int,
	) ([]Appliance, error)

	MarkOffline(
		ctx context.Context,
		applianceID int64,
	) error

	GetWithLatestReading(
		ctx context.Context,
		userID int64,
	) ([]ApplianceWithReading, error)
}

func NewApplianceRepo(db *sqlx.DB) ApplianceRepository {
	return &applianceRepo{
		db: db,
	}
}

func (r *applianceRepo) Create(ctx context.Context, a *Appliance) error {
	query := `
		INSERT INTO appliances (
			user_id,
			name,
			location,
			device_code,
			status
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowxContext(ctx, query,
		a.UserID,
		a.Name,
		a.Location,
		a.DeviceCode,
		a.Status,
	).Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt)
}

func (r *applianceRepo) GetByID(ctx context.Context, id int64) (*Appliance, error) {
	var a Appliance

	err := r.db.GetContext(ctx, &a, `
		SELECT * FROM appliances
		WHERE id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *applianceRepo) List(ctx context.Context, userID int64) ([]Appliance, error) {
	var list []Appliance

	err := r.db.SelectContext(ctx, &list, `
		SELECT * FROM appliances
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)

	return list, err
}

func (r *applianceRepo) Update(ctx context.Context, a *Appliance) error {
	query := `
		UPDATE appliances
		SET name = $1,
		    location = $2,
		    updated_at = NOW()
		WHERE id = $3 AND user_id = $4
	`

	res, err := r.db.ExecContext(ctx, query,
		a.Name,
		a.Location,
		a.ID,
		a.UserID,
	)

	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *applianceRepo) Delete(ctx context.Context, userID, id int64) error {
	res, err := r.db.ExecContext(ctx, `
		DELETE FROM appliances
		WHERE id = $1 AND user_id = $2
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

func (r *applianceRepo) GetByDeviceCode(
	ctx context.Context,
	deviceCode string,
) (*Appliance, error) {
	query := `
		SELECT
			id,
			user_id,
			name,
			location,
			device_code,
			status,
			last_reading,
			created_at,
			updated_at
		FROM appliances
		WHERE device_code = $1
		LIMIT 1
	`

	var appliance Appliance

	err := r.db.QueryRowContext(
		ctx,
		query,
		deviceCode,
	).Scan(
		&appliance.ID,
		&appliance.UserID,
		&appliance.Name,
		&appliance.Location,
		&appliance.DeviceCode,
		&appliance.Status,
		&appliance.LastReading,
		&appliance.CreatedAt,
		&appliance.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrApplianceNotFound
		}
		return nil, err
	}

	return &appliance, nil
}

func (r *applianceRepo) UpdateLastReading(
	ctx context.Context,
	applianceID int64,
	lastReading time.Time,
) error {
	query := `
		UPDATE appliances
		SET
			last_reading = $1,
			status = $2,
			updated_at = NOW()
		WHERE id = $3
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		lastReading,
		ApplianceStatusOnline,
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
		return ErrApplianceNotFound
	}

	return nil
}

func (r *applianceRepo) GetOfflineCandidates(
	ctx context.Context,
	offlineMinutes int,
) ([]Appliance, error) {
	query := `
		SELECT
			id,
			user_id,
			name,
			location,
			device_code,
			status,
			last_reading,
			created_at,
			updated_at
		FROM appliances
		WHERE
			last_reading IS NOT NULL
			AND last_reading < NOW() - ($1 * INTERVAL '1 minute')
			AND status != $2
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		offlineMinutes,
		ApplianceStatusOffline,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appliances []Appliance

	for rows.Next() {
		var a Appliance

		err := rows.Scan(
			&a.ID,
			&a.UserID,
			&a.Name,
			&a.Location,
			&a.DeviceCode,
			&a.Status,
			&a.LastReading,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		appliances = append(appliances, a)
	}

	return appliances, nil
}

func (r *applianceRepo) MarkOffline(
	ctx context.Context,
	applianceID int64,
) error {
	query := `
		UPDATE appliances
		SET
			status = $1,
			updated_at = NOW()
		WHERE id = $2
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		ApplianceStatusOffline,
		applianceID,
	)

	return err
}

func (r *applianceRepo) GetWithLatestReading(
	ctx context.Context,
	userID int64,
) ([]ApplianceWithReading, error) {

	query := `
		SELECT
			a.id,
			a.name,
			a.last_reading as last_seen,

			er.power

		FROM appliances a

		LEFT JOIN LATERAL (
			SELECT power
			FROM energy_readings
			WHERE appliance_id = a.id
			ORDER BY ts DESC
			LIMIT 1
		) er ON true

		WHERE a.user_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []ApplianceWithReading

	for rows.Next() {
		var a ApplianceWithReading
		if err := rows.Scan(
			&a.ID,
			&a.Name,
			&a.LastSeen,
			&a.Power,
		); err != nil {
			return nil, err
		}

		result = append(result, a)
	}

	return result, nil
}
