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

func (r *DeviceRepo) DeviceExists(deviceSerial string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM devices WHERE device_serial=$1)`

	var exists bool
	err := r.db.Get(&exists, query, deviceSerial)
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

func (r *DeviceRepo) GetDeviceByCode(deviceCode string) (*Device, error) {
	query := "SELECT * FROM devices WHERE device_code = $1"

	var device Device
	if err := r.db.Get(&device, query, deviceCode); err != nil {
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

func (r *DeviceRepo) GetUserDevices(ctx context.Context, userID int64) ([]DeviceClaimResponse, error) {
	query := `
	SELECT dc.*, d.is_active, d.last_active, d.device_code FROM devices AS d
	JOIN device_claims AS dc ON dc.device_id = d.id
	WHERE dc.user_id = $1;
	`

	var devices []DeviceClaimResponse

	if err := r.db.SelectContext(ctx, &devices, query, userID); err != nil {
		return []DeviceClaimResponse{}, err
	}

	return devices, nil
}

func (r *DeviceRepo) GetUserDevice(ctx context.Context, userID int64, deviceID int64) (*DeviceClaimResponse, error) {
	query := `
	SELECT dc.*, d.is_active, d.last_active, d.device_code FROM devices AS d
	JOIN device_claims AS dc ON dc.device_id = d.id
	WHERE dc.user_id = $1
	and d.id = $2
	`

	var device DeviceClaimResponse

	if err := r.db.Get(&device, query, userID, deviceID); err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *DeviceRepo) SaveDevice(ctx context.Context, record DeviceRequest) (int64, error) {
	query := `
		INSERT INTO devices (
			device_code,
		) VALUES (
			$1,
		)
		RETURNING id
	`
	var id int64
	err := r.db.QueryRowContext(ctx, query,
		record.DeviceCode,
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
			device_serial = $1,
		WHERE
			id = $2
		RETURNING id
	`

	result, err := r.db.ExecContext(ctx,
		query,
		data.DeviceCode,
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
