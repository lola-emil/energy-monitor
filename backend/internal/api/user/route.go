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

	mux.Get("/", r.handler.GetUsers)
	mux.Get("/{id}", r.handler.GetUser)
	mux.Patch("/{id}", r.handler.UpdateUser)

	return mux
}
