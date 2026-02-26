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

func (r *DeviceRepo) GetDeviceById(id int64) (*Device, error) {
	query := "SELECT * FROM devices WHERE id = $1"

	var device Device
	if err := r.db.Get(&device, query, id); err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *DeviceRepo) GetDevices(ctx context.Context) ([]Device, error) {
	query := "SELECT * FROM devices WHERE 1"

	var devices []Device

	if err := r.db.SelectContext(ctx, &devices, query); err != nil {
		return []Device{}, err
	}

	return devices, nil
}

func (r *DeviceRepo) SaveDevice(ctx context.Context, record Device) (int64, error) {
	query := `
		INSERT INTO devices (
			device_name,
			device_serial,
			activation_code
		) VALUES (
			:device_name
			:device_serial,
			:activation_code 
		)
	`

	result, err := r.db.ExecContext(ctx, query, map[string]any{
		"device_name":     record.DeviceName,
		"device_serial":   record.DeviceSerial,
		"activation_code": record.ActivationCode,
	})

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *DeviceRepo) UpdateDeviceById(ctx context.Context, id int64, data Device) (int64, error) {

	query := `
		UPDATE devices
		SET
			device_name = :device_name,
			device_serial = :device_serial
			activation_code = :activation_code
		WHERE
			id = :id
	`

	result, err := r.db.ExecContext(ctx, query, map[string]any{
		"device_name":     data.DeviceName,
		"device_serial":   data.DeviceSerial,
		"activation_code": data.ActivationCode,
	})

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
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
	query := "SELECT * FROM device_claims WHERE 1"

	var devices []DeviceClaim

	if err := r.db.SelectContext(ctx, &devices, query); err != nil {
		return []DeviceClaim{}, err
	}

	return devices, nil
}
