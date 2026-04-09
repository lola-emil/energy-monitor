package deviceclaim

import "github.com/go-chi/chi/v5"

type DeviceClaimRoute struct {
	handler *DeviceClaimHandler
}

func NewDeviceClaimRoute(handler *DeviceClaimHandler) *DeviceClaimRoute {
	return &DeviceClaimRoute{
		handler: handler,
	}
}

func (r *DeviceClaimRoute) RegisterRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Post("/", r.handler.ClaimDevice)
	mux.Get("/{device-id}", r.handler.GetDeviceClaims)
	mux.Get("/{claim-id}", r.handler.GetDeviceClaim)

	return mux
}
