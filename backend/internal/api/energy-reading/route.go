package energyreading

import "github.com/go-chi/chi/v5"

type EnergyReadingRoute struct {
	handler *EnergyReadingHandler
}

func NewEnergyReadingRoute(handler *EnergyReadingHandler) *EnergyReadingRoute {
	return &EnergyReadingRoute{
		handler: handler,
	}
}

func (r *EnergyReadingRoute) RegisterRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/", r.handler.TestFunction)

	return mux
}
