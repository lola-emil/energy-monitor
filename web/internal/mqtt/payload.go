package mqtt

import "time"

type IncomingReading struct {
	DeviceCode string `json:"id"`

	Voltage     float64 `json:"v"`
	Current     float64 `json:"A"`
	Power       float64 `json:"W"`
	EnergyKWh   float64 `json:"e_kWh"`
	FrequencyHz float64 `json:"hz"`

	Timestamp time.Time `json:"timestamp"`
}
