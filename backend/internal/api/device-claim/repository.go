package deviceclaim

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type DeviceClaimRepo struct {
	db *sqlx.DB
}

func NewDeviceClaimRepo(db *sqlx.DB) *DeviceClaimRepo {
	return &DeviceClaimRepo{
		db: db,
	}
}

func (r *DeviceClaimRepo) GetDeviceClaimById(id int64) (*DeviceClaim, error) {
	query := "SELECT * FROM device_claims WHERE id = $1"

	var device DeviceClaim
	if err := r.db.Get(&device, query, id); err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *DeviceClaimRepo) GetDeviceClaims(ctx context.Context, deviceId int64) ([]DeviceClaim, error) {
	query := "SELECT * FROM device_claims "

	var devices []DeviceClaim

	if err := r.db.SelectContext(ctx, &devices, query); err != nil {
		return []DeviceClaim{}, err
	}

	return devices, nil
}

func (r *DeviceClaimRepo) DeviceAlreadyTaken(ctx context.Context, deviceId int64, userId int64) (bool, error) {
	query := "SELECT EXISTS(SELECT * FROM device_claims WHERE device_id = $1 AND user_id = $2)"

	var exists bool
	err := r.db.Get(&exists, query, deviceId, userId)

	return exists, err
}

func (r *DeviceClaimRepo) ClaimDevice(ctx context.Context, claim DeviceClaim) (int64, error) {
	query := `
		INSERT INTO device_claims (
			device_id,
			user_id,
			device_name
		) VALUES (
			$1, 
			$2, 
			$3
		 )

		 RETURNING id
	`

	var id int64
	err := r.db.QueryRowContext(ctx, query,
		claim.DeviceId,
		claim.UserId,
		claim.DeviceName).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
