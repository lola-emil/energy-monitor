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
			:device_id,
			:voltage,
			:current,
			:power_kwh
		)
	`

	result, err := r.db.ExecContext(ctx, query, map[string]any{
		"device_id": record.DeviceId,
		"voltage":   record.Voltage,
		"current":   record.Current,
		"power_kwh": record.PowerKwh,
	})

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *EnergyReadingRepo) UpdateEnergyReadingById(ctx context.Context, id int64, data EnergyReading) (int64, error) {

	query := `
		UPDATE energy_readings
		SET
			device_id = :device_id,
			voltage =   :voltage
			current =   :current
			power_kwh = :power_kwh
		WHERE
			id = :id
	`

	result, err := r.db.ExecContext(ctx, query, map[string]any{
		"device_id": data.DeviceId,
		"voltage":   data.Voltage,
		"current":   data.Current,
		"power_kwh": data.PowerKwh,
	})

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *EnergyReadingRepo) DeleteEnergyReadingById(ctx context.Context, id int64) (int64, error) {
	query := "DELETE FROM energy_readings WHERE id = $1"

	result, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return 0, nil
	}

	return result.RowsAffected()
}
