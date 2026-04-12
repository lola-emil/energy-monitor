package event

import (
	energyreading "backend/internal/api/energy-reading"
)

type Event struct {
	UserID   int64                           `json:"user_id"`
	DeviceID int64                           `json:"device_id"`
	Value    energyreading.EnergyReadingBody `json:"value"`
}
