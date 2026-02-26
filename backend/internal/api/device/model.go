package device

import "time"

type Device struct {
	ID           int64  `db:"id" json:"id"`
	DeviceName   string `db:"device_name" json:"device_name"`
	DeviceSerial string `db:"device_serial" json:"device_serial"`
	// UserId       int64  `db:"user_id" json:"user_id"`
	IsActive   bool      `db:"is_active" json:"is_active"`
	LastActive time.Time `db:"last_active" json:"last_active"`
	ApiKeyHash string    `db:"apikey_hash" json:"-"`

	ActivationCode string `db:"activation_code" json:"-"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type DeviceClaim struct {
	ID       int64 `db:"id" json:"id"`
	DeviceId int64 `db:"device_id" json:"device_id"`
	UserId   int64 `db:"user_id" json:"user_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
