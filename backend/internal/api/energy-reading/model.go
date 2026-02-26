package energyreading

import (
	"time"

	"github.com/shopspring/decimal"
)

type EnergyReading struct {
	ID int64

	DeviceId int64

	Voltage  decimal.Decimal
	Current  decimal.Decimal
	PowerKwh decimal.Decimal

	timestamp time.Time
}
