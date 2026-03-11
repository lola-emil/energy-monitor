package device

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type DeviceRepo struct {
	db *sqlx.DB
}

func NewDeviceRepo(db *sqlx.DB) *DeviceRepo {
	return &DeviceRepo{
		db: db,
	}
}

func (r *DeviceRepo) DeviceExists(deviceSerial string, activationCode string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM devices WHERE device_serial=$1 OR activation_code=$2)`

	var exists bool
	err := r.db.Get(&exists, query, deviceSerial, activationCode)
	return exists, err
}

func (r *DeviceRepo) GetDeviceById(id int64) (*Device, error) {
	query := "SELECT * FROM devices WHERE id = $1"

	var device Device
	if err := r.db.Get(&device, query, id); err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *DeviceRepo) GetDevices(ctx context.Context) ([]Device, error) {
	query := "SELECT * FROM devices "

	var devices []Device

	if err := r.db.SelectContext(ctx, &devices, query); err != nil {
		return []Device{}, err
	}

	return devices, nil
}

func (r *DeviceRepo) SaveDevice(ctx context.Context, record DeviceRequest) (int64, error) {
	query := `
		INSERT INTO devices (
			device_name,
			device_serial,
			activation_code
		) VALUES (
			$1,
			$2,
			$3 
		)
		RETURNING id
	`
	var id int64
	err := r.db.QueryRowContext(ctx, query,
		record.DeviceName,
		record.DeviceSerial,
		record.ActivationCode,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *DeviceRepo) UpdateDeviceById(ctx context.Context, id int64, data Device) (int64, error) {

	query := `
		UPDATE devices
		SET
			device_name = $1,
			device_serial = $2,
			activation_code = $3
		WHERE
			id = $4
		RETURNING id
	`

	result, err := r.db.ExecContext(ctx,
		query,
		data.DeviceName,
		data.DeviceSerial,
		data.ActivationCode,
	)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *DeviceRepo) DeleteDeviceById(ctx context.Context, id int64) (int64, error) {
	query := "DELETE FROM devices WHERE id = $1"

	result, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return 0, nil
	}

	return result.RowsAffected()
}

func (r *DeviceRepo) GetDeviceClaimById(id int64) (*DeviceClaim, error) {
	query := "SELECT * FROM device_claims WHERE id = $1"

	var device DeviceClaim
	if err := r.db.Get(&device, query, id); err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *DeviceRepo) GetDeviceClaims(ctx context.Context, deviceId int64) ([]DeviceClaim, error) {
	query := "SELECT * FROM device_claims "

	var devices []DeviceClaim

	if err := r.db.SelectContext(ctx, &devices, query); err != nil {
		return []DeviceClaim{}, err
	}

	return devices, nil
}
