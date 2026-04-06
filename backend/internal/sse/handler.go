package sse

import (
	"fmt"
	"net/http"
)

func (b *SSEBroker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	client := make(SSEClient)

	// Register client
	b.register <- client

	// Cleanup on disconnect
	defer func() {
		b.unregister <- client
	}()

	ctx := r.Context()

	for {
		select {
		case msg := <-client:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()

		case <-ctx.Done():
			return
		}
	}
}
