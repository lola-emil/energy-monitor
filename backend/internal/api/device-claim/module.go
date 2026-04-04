package deviceclaim

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterModule(db *sqlx.DB) *chi.Mux {
	claimsRepo := NewDeviceClaimRepo(db)
	claimsHandler := NewDeviceClaimHandler(claimsRepo)
	claimsRoute := NewDeviceClaimRoute(claimsHandler)

	return claimsRoute.RegisterRoutes()
}
