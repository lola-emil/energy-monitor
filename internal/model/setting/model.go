package setting

import "time"

type Settings struct {
	Id     int64 `json:"id" db:"id"`
	UserID int64 `json:"user_id" db:"user_id"`

	// Billing
	Currency           string  `json:"currency" db:"currency"`
	RatePerKWh         float64 `json:"rate_per_kwh" db:"rate_per_kwh"`
	FixedMonthlyCharge float64 `json:"fixed_monthly_charge" db:"fixed_monthly_charge"`

	// Display
	DefaultAnalyticsRange  string `json:"default_analytics_range" db:"default_analytics_range"` // "today", "7d", "month"
	RefreshIntervalSeconds int    `json:"refresh_interval_seconds" db:"refresh_interval_seconds"`
	TimeFormat             string `json:"time_format" db:"time_format"` // "24h" or "12h"

	// Alerts
	EnableVoltageAlerts   bool    `json:"enable_voltage_alerts" db:"enable_voltage_alerts"`
	OverVoltageThreshold  float64 `json:"over_voltage_threshold" db:"over_voltage_threshold"`
	UnderVoltageThreshold float64 `json:"under_voltage_threshold" db:"under_voltage_threshold"`

	EnableCurrentAlerts  bool    `json:"enable_current_alerts" db:"enable_current_alerts"`
	OverCurrentThreshold float64 `json:"over_current_threshold" db:"over_current_threshold"`

	EnableOfflineAlerts bool `json:"enable_offline_alerts" db:"enable_offline_alerts"`

	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
