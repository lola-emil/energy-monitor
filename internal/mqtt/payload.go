package mqtt

import "time"

type IncomingReading struct {
	DeviceCode string `json:"device_code"`

	Voltage     float64 `json:"voltage"`
	Current     float64 `json:"current"`
	Power       float64 `json:"power"`
	EnergyKWh   float64 `json:"energy_kwh"`
	FrequencyHz float64 `json:"frequency_hz"`

	Timestamp time.Time `json:"timestamp"`
}
