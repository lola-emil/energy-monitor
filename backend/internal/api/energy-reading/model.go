package energyreading

import "time"

type EnergyReading struct {
	ID       int64
	DeviceId int64
	Voltage  float32
	Current  float32
	PowerKwh float32

	timestamp time.Time
}
