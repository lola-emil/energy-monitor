package energyreading

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterModule(db *sqlx.DB) *chi.Mux {
	energyReadingRepo := NewEnergyReadingRepo(db)
	energyReadingHandler := NewEnergyReadingHandler(energyReadingRepo)
	energyReadingroute := NewEnergyReadingRoute(energyReadingHandler)

	return energyReadingroute.RegisterRoutes()
}
