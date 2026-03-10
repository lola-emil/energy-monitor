package mymqtt

import (
	"time"

	"github.com/shopspring/decimal"
)

type EnergyReading struct {
	ID int64

	DeviceId int64           `db:"device_id"`
	Voltage  decimal.Decimal `db:"voltage"`
	Current  decimal.Decimal `db:"current"`
	PowerKwh decimal.Decimal `db:"power_kwh"`

	timestamp *time.Time
}

type Device struct {
	ID           int64  `db:"id" json:"id"`
	DeviceName   string `db:"device_name" json:"device_name"`
	DeviceSerial string `db:"device_serial" json:"device_serial"`
	// UserId       int64  `db:"user_id" json:"user_id"`
	IsActive   bool       `db:"is_active" json:"is_active"`
	LastActive *time.Time `db:"last_active" json:"last_active"`
	ApiKeyHash string     `db:"apikey_hash" json:"-"`

	ActivationCode string `db:"activation_code" json:"-"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type SensorData struct {
	Token    string
	DeviceId int64           `json:"id"`
	Voltage  decimal.Decimal `json:"voltage"`
	Current  decimal.Decimal `json:"current"`
	PowerKwh decimal.Decimal `json:"power"`
}

type DeviceAuth struct {
	DeviceId       int64  `json:"id"`
	DeviceSerial   string `json:"serial"`
	ActivationCode string `json:"act_code"`
}
