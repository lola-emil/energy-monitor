package device

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterModule(db *sqlx.DB) *chi.Mux {
	deviceRepo := NewDeviceRepo(db)
	deviceHandler := NewDeviceHandler(deviceRepo)
	deviceRoute := NewDeviceRoute(deviceHandler)

	return deviceRoute.RegisterRoutes()
}
