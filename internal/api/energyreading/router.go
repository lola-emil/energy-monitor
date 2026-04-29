package energyreading

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *ReadingHandler) {
	r.Get("/summary", h.GetSummary)
}

// r.Get("/summary", h.GetSummary)
