package energyreading

import (
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/shopspring/decimal"
)

type EnergyReading struct {
	ID int64

	DeviceId int64           `db:"device_id"`
	Voltage  decimal.Decimal `db:"voltage"`
	Current  decimal.Decimal `db:"current"`
	PowerKwh decimal.Decimal `db:"power_kwh"`

	timestamp time.Time
}

// --- DTOs
type EnergyReadingRequest struct {
	DeviceId int64           `json:"device_id" validate:"required,gt=0"`
	Voltage  decimal.Decimal `json:"voltage" validate:"required,numeric12_2"`
	Current  decimal.Decimal `json:"current" validate:"required,numeric12_2"`
	PowerKwh decimal.Decimal `json:"power_kwh" validate:"required,numeric12_2"`
}

func numeric12_2(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(decimal.Decimal)
	if !ok {
		return false
	}

	// Must be >= 0 (optional but recommended for energy readings)
	if val.IsNegative() {
		return false
	}

	// Check scale (decimal places)
	if val.Exponent() < -2 {
		return false
	}

	// Check total digits
	str := val.Abs().StringFixed(2)
	digits := 0
	for _, r := range str {
		if r >= '0' && r <= '9' {
			digits++
		}
	}

	return digits <= 12
}
