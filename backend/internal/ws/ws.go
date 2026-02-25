package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWSConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatalf("Websocket error: %s", err)
		return
	}

	defer ws.Close()

	for {
		// Read message from browser
		_, msg, err := ws.ReadMessage()

		if err != nil {
			fmt.Println("read error: ", err)
			break
		}

		fmt.Printf("Received: %s\n", msg)

		// Write message back to browser
		if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Write error: ", err)
			break
		}
	}
}
