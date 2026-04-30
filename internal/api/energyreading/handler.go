package energyreading

import (
	"encoding/json"
	"energy-monitor-server/internal/model/energyreading"
	httputil "energy-monitor-server/internal/utils/http"
	"net/http"
)

type ReadingHandler struct {
	service *ReadingService
}

func NewReadingHandler(service *ReadingService) *ReadingHandler {
	return &ReadingHandler{service: service}
}

func (h *ReadingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req energyreading.EnergyReading

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", 400)
		return
	}

	userID := httputil.GetUserID(r)

	if err := h.service.Create(r.Context(), userID, &req); err != nil {
		http.Error(w, "failed", 500)
		return
	}

	json.NewEncoder(w).Encode(req)
}
func (h *ReadingHandler) GetSummary(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID := httputil.GetUserID(r)

	summary, err := h.service.GetSummary(
		r.Context(),
		userID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(summary)
}

func (h *ReadingHandler) GetChart(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID := httputil.GetUserID(r)

	data, err := h.service.GetEnergyChart(
		r.Context(),
		userID,
	)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(data)
}
