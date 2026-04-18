package dashboard

import "time"

type Overview struct {
	TotalEnergyConsumed float64 `json:"total_consumed"`
	AvgVoltage          float64 `json:"avg_volt"`
	AvgPowerDraw        float64 `json:"avg_power"`
	AvgCurrent          float64 `json:"avg_freq"`
}

type MonthlyPower struct {
	Month        time.Time `json:"month"`
	AvgPowerDraw float64   `json:"avg_power_draw"`
}

type MonthlyConsumption struct {
	Month       time.Time `db:"month" json:"month"`
	Consumption float64   `db:"energy_kwh" json:"energy_kwh"`
}
