package dashboard

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterModule(db *sqlx.DB) *chi.Mux {
	dashboardRepo := NewDashboardRepo(db)
	dashboardHandler := NewDashboardHandler(dashboardRepo)
	dashboardRoute := NewDashboardRoute(dashboardHandler)

	return dashboardRoute.RegisterRoutes()
}
