package device

import "github.com/go-chi/chi/v5"

type DeviceRoute struct {
	handler *DeviceHandler
}

func NewDeviceRoute(handler *DeviceHandler) *DeviceRoute {
	return &DeviceRoute{
		handler: handler,
	}
}

func (r *DeviceRoute) RegisterRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Post("/", r.handler.AddDevice)
	mux.Get("/", r.handler.GetDevices)

	return mux
}
