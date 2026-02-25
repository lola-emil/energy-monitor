package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterModule(db *sqlx.DB) *chi.Mux {
	authRepo := NewAuthRepo(db)
	authHandler := NewAuthHandler(authRepo)
	authRoute := NewAuthRoute(authHandler)

	return authRoute.RegisterRoutes()
}
