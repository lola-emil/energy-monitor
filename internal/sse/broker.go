package sse

import (
	"encoding/json"
	"log"
)

type Broker struct {
	clients map[chan []byte]bool
}

func NewBroker() *Broker {
	return &Broker{
		clients: make(map[chan []byte]bool),
	}
}

func (b *Broker) AddClient(ch chan []byte) {
	b.clients[ch] = true
}

func (b *Broker) RemoveClient(ch chan []byte) {
	delete(b.clients, ch)
	close(ch)
}

func (b *Broker) Broadcast(v any) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Println("SSE marshal error:", err)
		return
	}

	for client := range b.clients {
		select {
		case client <- data:
		default:
			// skip slow clients
		}
	}
}
