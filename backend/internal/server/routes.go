package server

import (
	"backend/internal/api/auth"
	"backend/internal/api/device"
	energyreading "backend/internal/api/energy-reading"
	"backend/internal/api/user"
	"backend/internal/ws"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Mount("/auth", auth.RegisterModule(s.db.GetInstance()))

	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", user.RegisterModule(s.db.GetInstance()))
		r.Mount("/devices", device.RegisterModule(s.db.GetInstance()))
		r.Mount("/energy-readings", energyreading.RegisterModule(s.db.GetInstance()))
	})

	r.Get("/health", s.healthHandler)

	// INITILIZE WEBSOCKET
	r.HandleFunc("/ws", ws.HandleWSConnections)

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
