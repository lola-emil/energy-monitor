package settings

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *SettingsHandler) {
	r.Get("/", h.Get)
	r.Put("/", h.Update)
}
