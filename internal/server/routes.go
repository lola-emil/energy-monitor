package server

import (
	"encoding/json"
	alertapi "energy-monitor-server/internal/api/alert"
	applianceapi "energy-monitor-server/internal/api/appliance"
	"energy-monitor-server/internal/api/settings"
	"energy-monitor-server/internal/auth"
	appmiddleware "energy-monitor-server/internal/middleware"
	"energy-monitor-server/internal/model/alert"
	"energy-monitor-server/internal/model/appliance"
	"energy-monitor-server/internal/model/energyreading"
	"energy-monitor-server/internal/model/setting"
	"energy-monitor-server/internal/model/user"
	"energy-monitor-server/internal/sse"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes(
	userRepo user.UserRepository,
	applianceRepo appliance.ApplianceRepository,
	alertRepo alert.AlertRepository,
	settingsRepo setting.SettingsRepository,
	readingRepo energyreading.ReadingRepository,
	broker *sse.Broker,

) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Mga service
	authService := auth.NewAuthService(userRepo)
	applianceService := applianceapi.NewApplianceService(applianceRepo)
	alertService := alertapi.NewAlertService(alertRepo)
	settingService := settings.NewSettingService(settingsRepo)

	// Mga handlers
	authHandler := auth.NewAuthHandler(authService)
	applianceHandler := applianceapi.NewApplianceHandler(applianceService)
	alertHandler := alertapi.NewAlertHandler(alertService)
	settingsHandler := settings.NewSettingHandler(settingService)

	r.Route("/api", func(r chi.Router) {

		// SSE
		r.Get("/stream", broker.ServeHTTP)

		r.Route("/auth", func(r chi.Router) {
			auth.RegisterRoutes(r, authHandler)
		})
		r.Group(func(r chi.Router) {
			r.Use(appmiddleware.AuthMiddleware)

			r.Route("/appliances", func(r chi.Router) {
				applianceapi.RegisterRoutes(r, applianceHandler)
			})

			r.Route("/alerts", func(r chi.Router) {
				alertapi.RegisterRoutes(r, alertHandler)
			})

			r.Route("/settings", func(r chi.Router) {
				settings.RegisterRoutes(r, settingsHandler)
			})
		})
	})

	spaHandler := SPAHandler("frontend/dist")
	r.Handle("/*", spaHandler)

	r.Get("/health", s.healthHandler)

	return r
}

func SPAHandler(staticPath string) http.Handler {
	// If VITE_DEV_SERVER_URL is set, proxy frontend requests to Vite
	if viteURL := os.Getenv("VITE_DEV_SERVER_URL"); viteURL != "" {
		target, err := url.Parse(viteURL)
		if err == nil {
			proxy := httputil.NewSingleHostReverseProxy(target)

			originalDirector := proxy.Director
			proxy.Director = func(req *http.Request) {
				originalDirector(req)

				// Preserve original host if you want Vite to see its own host instead:
				req.Host = target.Host

				// Optional forwarded headers
				req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
				req.Header.Set("X-Forwarded-Proto", "http")
			}

			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				proxy.ServeHTTP(w, r)
			})
		}
	}

	// Production static serving
	fs := http.FileServer(http.Dir(staticPath))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(staticPath, filepath.Clean(r.URL.Path))

		info, err := os.Stat(path)
		if os.IsNotExist(err) || (err == nil && info.IsDir()) {
			http.ServeFile(w, r, filepath.Join(staticPath, "index.html"))
			return
		}

		fs.ServeHTTP(w, r)
	})
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
