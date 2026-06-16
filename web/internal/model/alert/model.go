package alert

import "time"

type AlertSeverity string

const (
	AlertSeverityInfo   AlertSeverity = "info"
	AlertSeverityMedium AlertSeverity = "medium"
	AlertSeverityHigh   AlertSeverity = "high"
)

type AlertType string

const (
	AlertTypeOverVoltage  AlertType = "over_voltage"
	AlertTypeUnderVoltage AlertType = "under_voltage"
	AlertTypeOverCurrent  AlertType = "over_current"
	AlertTypeOffline      AlertType = "offline"
)

type Alert struct {
	ID          int64         `json:"id" db:"id"`
	ApplianceID *int64        `json:"appliance_id,omitempty" db:"appliance_id"` // nil for system-wide
	Type        AlertType     `json:"type" db:"type"`
	Severity    AlertSeverity `json:"severity" db:"severity"`
	Message     string        `json:"message" db:"message"`

	TriggeredAt time.Time  `json:"triggered_at" db:"triggered_at"`
	ResolvedAt  *time.Time `json:"resolved_at,omitempty" db:"resolved_at"`
}
