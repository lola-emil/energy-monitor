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

	mux.Get("/", r.handler.TestFunction)

	return mux
}
