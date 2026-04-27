package energyreading

import "time"

type EnergyReading struct {
	ID          int64     `json:"id" db:"id"`
	ApplianceID int64     `json:"appliance_id" db:"appliance_id"`
	Timestamp   time.Time `json:"timestamp" db:"ts"`

	Voltage     float64 `json:"voltage" db:"voltage"`       // volts
	Current     float64 `json:"current" db:"current"`       // amps
	Power       float64 `json:"power" db:"power"`           // watts
	EnergyKWh   float64 `json:"energy_kwh" db:"energy_kwh"` // incremental or cumulative, your choice
	FrequencyHz float64 `json:"frequency_hz" db:"frequency_hz"`
}

type ReadingSummary struct {
	TotalEnergyKWh float64 `json:"total_energy_kwh"`
	AvgVoltage     float64 `json:"avg_voltage"`
	AvgCurrent     float64 `json:"avg_current"`
	AvgPower       float64 `json:"avg_power"`
	EstimatedCost  float64 `json:"estimated_cost"`
}
