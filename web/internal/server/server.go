package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"energy-monitor-server/internal/alerts"
	"energy-monitor-server/internal/database"
	"energy-monitor-server/internal/model/alert"
	"energy-monitor-server/internal/model/appliance"
	"energy-monitor-server/internal/model/energyreading"
	"energy-monitor-server/internal/model/setting"
	"energy-monitor-server/internal/model/user"
	"energy-monitor-server/internal/mqtt"
	"energy-monitor-server/internal/offline"
	"energy-monitor-server/internal/sse"
)

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.New()

	NewServer := &Server{
		port: port,

		db: db,
	}

	dbInstance := db.GetInstance()

	userRepo := user.NewUserRepository(dbInstance)
	applianceRepo := appliance.NewApplianceRepo(dbInstance)
	alertRepo := alert.NewAlertRepository(dbInstance)
	settingsRepo := setting.NewSettingsRepository(dbInstance)
	readingRepo := energyreading.NewReadingRepository(dbInstance)

	// MQTT
	broker := sse.NewBroker()
	// ALERT EVAL
	alertEvaluator := alerts.NewEvaluator(alertRepo, settingsRepo)

	mqttHandler := mqtt.NewHandler(
		applianceRepo,
		readingRepo,
		broker,
		alertEvaluator,
	)

	go mqtt.StartSubscriber(mqttHandler)

	// Offline Checker
	offline.StartOfflineChecker(applianceRepo)

	// Declare Server config
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", NewServer.port),
		Handler: NewServer.RegisterRoutes(
			userRepo,
			applianceRepo,
			alertRepo,
			settingsRepo,
			readingRepo,
			broker,
		),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
