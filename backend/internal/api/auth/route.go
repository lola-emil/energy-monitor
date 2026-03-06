package auth

import "github.com/go-chi/chi/v5"

type AuthRoute struct {
	handler *AuthHandler
}

func NewAuthRoute(handler *AuthHandler) *AuthRoute {
	return &AuthRoute{
		handler: handler,
	}
}

func (r *AuthRoute) RegisterRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Post("/login", r.handler.Login)
	mux.Post("/register", r.handler.Register)

	return mux
}
