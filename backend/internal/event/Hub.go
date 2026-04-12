package event

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Hub struct {
	Clients    map[int64]map[chan Event]bool
	Register   chan subscription
	Unregister chan subscription
	Broadcast  chan Event
}

type subscription struct {
	UserID int64
	Ch     chan Event
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[int64]map[chan Event]bool),
		Register:   make(chan subscription),
		Unregister: make(chan subscription),
		Broadcast:  make(chan Event, 10),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case sub := <-h.Register:
			if h.Clients[sub.UserID] == nil {
				h.Clients[sub.UserID] = make(map[chan Event]bool)
			}
			h.Clients[sub.UserID][sub.Ch] = true

		case sub := <-h.Unregister:
			if conns, ok := h.Clients[sub.UserID]; ok {
				if _, exists := conns[sub.Ch]; exists {
					delete(conns, sub.Ch)
					close(sub.Ch)
				}
				if len(conns) == 0 {
					delete(h.Clients, sub.UserID)
				}
			}

		case event := <-h.Broadcast:
			if conns, ok := h.Clients[event.UserID]; ok {
				for ch := range conns {
					select {
					case ch <- event:
					default:
						// drop slow client
						close(ch)
						delete(conns, ch)
					}
				}
			}
		}

	}
}

func (h *Hub) SseHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("userid")
		id, err := strconv.ParseInt(userId, 10, 64)

		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		flusher, ok := w.(http.Flusher)

		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		ch := make(chan Event, 10)

		sub := subscription{
			UserID: id,
			Ch:     ch,
		}

		h.Register <- sub

		defer func() {
			h.Unregister <- sub
		}()

		for {
			select {
			case evt := <-ch:
				jsonData, _ := json.Marshal(evt)
				fmt.Fprintf(w, "event: read-%d\ndata: %s\n\n", evt.DeviceID, jsonData)

				flusher.Flush()

			case <-r.Context().Done():
				return
			}
		}

	}
}
