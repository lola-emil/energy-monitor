package alertapi

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *AlertHandler) {
	r.Get("/", h.List)
	r.Get("/{id}", h.Get)
	r.Put("/{id}/resolve", h.Resolve)
}
