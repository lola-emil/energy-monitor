package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"backend/internal/database"
	mymqtt "backend/internal/my-mqtt"
	"backend/internal/sse"
	"backend/internal/ws"
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

	wsHub := ws.NewHub()
	sseBroker := sse.NewSSEBroker()

	go sseBroker.Run()
	go wsHub.Run()

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
