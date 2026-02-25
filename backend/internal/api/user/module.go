package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterModule(db *sqlx.DB) *chi.Mux {
	userRepo := NewUserRepo(db)
	userHandler := NewUserHandler(userRepo)
	userRoute := NewUserRoute(userHandler)

	return userRoute.RegisterRoutes()
}
