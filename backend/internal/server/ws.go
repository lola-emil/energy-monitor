package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WSHub struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
}

func NewHub() *WSHub {
	return &WSHub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
}

func (h *WSHub) Run() {
	log.Println("WS running..")
	for {
		msg := <-h.broadcast

		for client := range h.clients {
			err := client.WriteMessage(websocket.TextMessage, msg)

			if err != nil {
				log.Println(err)
				client.Close()
				delete(h.clients, client)
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *WSHub) HandleWSConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	h.clients[conn] = true
}
