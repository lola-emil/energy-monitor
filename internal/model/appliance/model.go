package appliance

import "time"

type ApplianceStatus string

const (
	ApplianceStatusOnline  ApplianceStatus = "online"
	ApplianceStatusOffline ApplianceStatus = "offline"
)

type Appliance struct {
	ID       int64           `json:"id" db:"id"`
	UserID   int64           `json:"user_id" db:"user_id"`
	Name     string          `json:"name" db:"name"`
	Location string          `json:"location" db:"location"`
	Status   ApplianceStatus `json:"status" db:"status"`

	DeviceCode string `json:"device_code" db:"device_code"`

	LastReading *time.Time `json:"last_reading,omitempty" db:"last_reading"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}
