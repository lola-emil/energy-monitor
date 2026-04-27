package auth

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *AuthHandler) {
	r.Post("/login", h.Login)
	r.Post("/register", h.Register)
}
