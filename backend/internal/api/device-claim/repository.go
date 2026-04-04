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
