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
			COALESCE(SUM(er.energy_kwh), 0) AS total_energy_kwh,
			COALESCE(MAX(er.power), 0) AS peak_power,
			(
				SELECT COUNT(*)
				FROM appliances a
				WHERE
					a.user_id = $1
					AND a.status = 'online'
			) AS active_devices,
			(
				SELECT COUNT(*)
				FROM alerts al
				JOIN appliances a ON a.id = al.appliance_id
				WHERE
					a.user_id = $1
					AND al.resolved_at IS NULL
			) AS active_alerts
		FROM energy_readings er
		JOIN appliances a ON a.id = er.appliance_id
		WHERE
			a.user_id = $1
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
