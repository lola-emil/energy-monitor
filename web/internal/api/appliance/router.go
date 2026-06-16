package applianceapi

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *ApplianceHandler) {

	r.Get("/", h.List)
	r.Get("/status", h.GetStatus)
	r.Post("/", h.Create)
	r.Get("/{id}", h.Get)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
}
