package mymqtt

import (
	"time"

	"github.com/shopspring/decimal"
)

type EnergyReading struct {
	DeviceId int64           `db:"device_id"`
	Bucket   *time.Time      `db:"bucket"`
	Voltage  decimal.Decimal `db:"voltage"`
	Current  decimal.Decimal `db:"current"`
	PowerKwh decimal.Decimal `db:"power_kwh"`
}

type EnergyReadingBody struct {
	DeviceId int64

	Voltage  decimal.Decimal
	Current  decimal.Decimal
	PowerKwh decimal.Decimal
}

type Device struct {
	ID         int64  `db:"id" json:"id"`
	DeviceName string `db:"device_name" json:"device_name"`
	// UserId       int64  `db:"user_id" json:"user_id"`
	IsActive   bool       `db:"is_active" json:"is_active"`
	LastActive *time.Time `db:"last_active" json:"last_active"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type SensorData struct {
	Token    string          `json:"token"`
	Voltage  decimal.Decimal `json:"voltage"`
	Current  decimal.Decimal `json:"current"`
	PowerKwh decimal.Decimal `json:"power"`
}

type DeviceRegister struct {
	DeviceCode       string `json:"device_code"`
	RegistrationCode string `json:"register_code"`
}

type DeviceAuth struct {
	DeviceId       int64  `json:"id"`
	DeviceSerial   string `json:"serial"`
	ActivationCode string `json:"act_code"`
}
