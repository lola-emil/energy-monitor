package deviceclaim

import (
	"backend/internal/api/device"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterModule(db *sqlx.DB) *chi.Mux {
	claimsRepo := NewDeviceClaimRepo(db)
	deviceRepo := device.NewDeviceRepo(db)
	claimsHandler := NewDeviceClaimHandler(claimsRepo, deviceRepo)
	claimsRoute := NewDeviceClaimRoute(claimsHandler)

	return claimsRoute.RegisterRoutes()
}
