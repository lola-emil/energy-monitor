package user

import "github.com/go-chi/chi/v5"

type UserRoute struct {
	handler *UserHandler
}

func NewUserRoute(handler *UserHandler) *UserRoute {
	return &UserRoute{
		handler: handler,
	}
}

func (r *UserRoute) RegisterRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/", r.handler.GetUser)

	return mux
}
