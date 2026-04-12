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

type Device struct {
	ID         int64  `db:"id" json:"id"`
	DeviceCode string `db:"device_code" json:"device_code"`
	// UserId       int64  `db:"user_id" json:"user_id"`
	IsActive   bool       `db:"is_active" json:"is_active"`
	LastActive *time.Time `db:"last_active" json:"last_active"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type SensorData struct {
	Voltage   decimal.Decimal `json:"v"`
	Current   decimal.Decimal `json:"A"`
	PowerKwh  decimal.Decimal `json:"e"`
	PowerDraw decimal.Decimal `json:"w"`
}

type DeviceRegister struct {
	DeviceCode       string `json:"s"`
	RegistrationCode string `json:"c"`
}

type DeviceAuth struct {
	DeviceId       int64  `json:"id"`
	DeviceSerial   string `json:"serial"`
	ActivationCode string `json:"act_code"`
}

type DeviceClaim struct {
	ID         int64  `db:"id" json:"id"`
	DeviceId   int64  `db:"device_id" json:"device_id"`
	UserId     int64  `db:"user_id" json:"user_id"`
	DeviceName string `db:"device_name" json:"device_name"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}
