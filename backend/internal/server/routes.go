package server

import (
	"backend/internal/api/auth"
	"backend/internal/api/device"
	energyreading "backend/internal/api/energy-reading"
	"backend/internal/api/user"
	"backend/internal/ws"
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// SPAHandler serves a single page application.
func SPAHandler(staticPath string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Join internally call path.Clean to prevent directory traversal
		path := filepath.Join(staticPath, r.URL.Path)

		// check whether a file exists or is a directory at the given path
		fi, err := os.Stat(path)
		if os.IsNotExist(err) || fi.IsDir() {

			// set cache control header to prevent caching
			// this is to prevent the browser from caching the index.html
			// and serving old build of SPA App
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

			// file does not exist or path is a directory, serve index.html
			http.ServeFile(w, r, filepath.Join(staticPath, "index.html"))
			return
		}

		if err != nil {
			// if we got an error (that wasn't that the file doesn't exist) stating the
			// file, return a 500 internal server error and stop
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set cache control header to serve file for a year
		// static files in this case need to be cache busted
		// (usualy by appending a hash to the filename)
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")

		// otherwise, use http.FileServer to serve the static file
		http.FileServer(http.Dir(staticPath)).ServeHTTP(w, r)
	})
}

func (s *Server) RegisterRoutes() http.Handler {
	viteURL, _ := url.Parse("http://localhost:5173")
	viteProxy := httputil.NewSingleHostReverseProxy(viteURL)

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

	// Serve SPA
	if os.Getenv("ENV") == "dev" {
		r.Handle("/*", viteProxy)
	} else {
		r.Handle("/*", SPAHandler("./frontend/dist"))
	}

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
