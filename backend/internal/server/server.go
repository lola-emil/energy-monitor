package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"backend/internal/database"
	"backend/internal/event"
	mymqtt "backend/internal/my-mqtt"
)

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db: database.New(),
	}

	sseBroker := event.NewHub()
	go sseBroker.Run()

	mymqtt.StartMQTT(sseBroker, NewServer.db.GetInstance())
	// defer mqttClient.Disconnect(250)

	routes := NewServer.RegisterRoutes(sseBroker)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      routes,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
