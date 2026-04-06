package dashboard

import "time"

type Overview struct {
	TotalEnergyConsumed float64
	AvgVoltage          float64
	AvgPowerDraw        float64
	AvgCurrent          float64
}

type MonthlyPower struct {
	Month        time.Time `json:"month"`
	AvgPowerDraw float64   `json:"avg_power_draw"`
}
