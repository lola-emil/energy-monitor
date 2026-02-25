package device

import "github.com/jmoiron/sqlx"

type DeviceRepo struct {
	db *sqlx.DB
}

func NewDeviceRepo(db *sqlx.DB) *DeviceRepo {
	return &DeviceRepo{
		db: db,
	}
}
