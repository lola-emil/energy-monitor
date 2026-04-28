package user

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *UserHandler) {
	r.Get("/profile", h.GetProfile)
	r.Put("/profile", h.UpdateProfile)
}
