package dashboard

import "github.com/go-chi/chi/v5"

type DashboardRoute struct {
	handler *DashboardHandler
}

func NewDashboardRoute(handler *DashboardHandler) *DashboardRoute {
	return &DashboardRoute{
		handler: handler,
	}
}

func (r *DashboardRoute) RegisterRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/overview", r.handler.GetOverview)
	mux.Get("/montly-consumed", r.handler.GetMonthlyAvgPower)

	return mux
}
