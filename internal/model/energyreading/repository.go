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
	GetSummary(ctx context.Context, userID int64, applianceID *int64, from, to *time.Time) (*ReadingSummary, error)
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
	from, to *time.Time,
) (*ReadingSummary, error) {
	var s ReadingSummary

	query := `
		SELECT
			COALESCE(SUM(er.energy_kwh), 0) as total_energy_kwh,
			COALESCE(AVG(er.voltage), 0) as avg_voltage,
			COALESCE(AVG(er.current), 0) as avg_current,
			COALESCE(AVG(er.power), 0) as avg_power
		FROM energy_readings er
		JOIN appliances a ON er.appliance_id = a.id
		WHERE a.user_id = $1
	`

	err := r.db.GetContext(ctx, &s, query, userID)
	if err != nil {
		return nil, err
	}

	return &s, nil
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
