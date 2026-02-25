package device

import "time"

type Device struct {
	ID           int64
	DeviceName   string `db:"device_name" json:"device_name"`
	DeviceSerial string `db:"device_serial" json:"device_serial"`
	UserId       int64  `db:"user_id" json:"user_id"`
	IsActive     bool   `db:"is_active" json:"is_active"`
	ApiKeyHash   string `db:"apikey_hash" json:"-"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type DeviceActivationTokens struct {
	ID             int64  `db:"id" json:"id"`
	DeviceId       int64  `db:"device_id" json:"device_id"`
	ActivationCode string `db:"activation_code" json:"-"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
