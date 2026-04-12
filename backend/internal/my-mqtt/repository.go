package mymqtt

import (
	energyreading "backend/internal/api/energy-reading"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetDeviceByIdAndSerial(id int64, serial string) (*Device, error) {
	query := "SELECT * FROM devices WHERE id = $1 AND device_serial = $2"

	var device Device
	if err := r.db.Get(&device, query, id, serial); err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *Repository) GetDeviceClaimByDeviceId(deviceID int64) (*DeviceClaim, error) {
	query := "SELECT * FROM device_claims WHERE device_id = $1"
	var device DeviceClaim

	if err := r.db.Get(&device, query, deviceID); err != nil {
		return nil, err
	}

	return &device, nil

}

func (r *Repository) SaveDeviceReadings(data energyreading.EnergyReadingBody) (int64, error) {
	query := `
	INSERT INTO readings_raw 
	(device_id, voltage, current, power_kwh)
	VALUES
	($1, $2, $3, $4)
	RETURNING device_id
	`

	var id int64
	err := r.db.QueryRow(query,
		data.DeviceId,
		data.Voltage,
		data.Current, data.PowerKwh,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) SaveDevice(data DeviceRegister) (int64, error) {
	query := `
		INSERT INTO devices (
			device_code
		) VALUES (
			$1
		)
		RETURNING id
	`
	var id int64
	err := r.db.QueryRow(query,
		data.DeviceCode,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) DeviceExists(deviceCode string) (*Device, error) {
	query := "SELECT * FROM devices WHERE device_code = $1"

	var device Device
	if err := r.db.Get(&device, query, deviceCode); err != nil {
		return nil, err
	}

	return &device, nil
}
