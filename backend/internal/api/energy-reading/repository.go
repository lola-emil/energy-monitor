package energyreading

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type EnergyReadingRepo struct {
	db *sqlx.DB
}

func NewEnergyReadingRepo(db *sqlx.DB) *EnergyReadingRepo {
	return &EnergyReadingRepo{
		db: db,
	}
}

func (r *EnergyReadingRepo) GetEnergyReadingById(id int64) (*EnergyReading, error) {
	query := "SELECT * FROM energy_readings WHERE id = $1"

	var reading EnergyReading
	if err := r.db.Get(&reading, query, id); err != nil {
		return nil, err
	}

	return &reading, nil
}

func (r *EnergyReadingRepo) GetEnergyReadings(ctx context.Context) ([]EnergyReading, error) {
	query := "SELECT * FROM energy_readings WHERE 1"

	var readings []EnergyReading

	if err := r.db.SelectContext(ctx, &readings, query); err != nil {
		return []EnergyReading{}, err
	}

	return readings, nil
}

func (r *EnergyReadingRepo) SaveEnergyReading(ctx context.Context, record EnergyReading) (int64, error) {
	query := `
		INSERT INTO energy_readings (
			device_id,
			voltage,
			current,
			power_kwh
		) VALUES (
			$1,
			$2,
			$3,
			$4
		)

		RETURNING id
	`
	var id int64
	err := r.db.QueryRowContext(ctx, query, map[string]any{
		"device_id": record.DeviceId,
		"voltage":   record.Voltage,
		"current":   record.Current,
		"power_kwh": record.PowerKwh,
	}).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *EnergyReadingRepo) UpdateEnergyReadingById(ctx context.Context, id int64, data EnergyReading) (int64, error) {

	query := `
		UPDATE energy_readings
		SET
			device_id = $1,
			voltage =   $2
			current =   $3
			power_kwh = $4
		WHERE
			id = $5
	`

	result, err := r.db.ExecContext(ctx, query,
		data.DeviceId,
		data.Voltage,
		data.Current,
		data.PowerKwh,
	)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *EnergyReadingRepo) DeleteEnergyReadingById(ctx context.Context, id int64) (int64, error) {
	query := "DELETE FROM energy_readings WHERE id = $1"

	result, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return 0, nil
	}

	return result.RowsAffected()
}
