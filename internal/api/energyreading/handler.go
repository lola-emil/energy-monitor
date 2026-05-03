package energyreading

import (
	"encoding/json"
	"energy-monitor-server/internal/model/energyreading"
	httputil "energy-monitor-server/internal/utils/http"
	"net/http"
	"strconv"
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

func (h *ReadingHandler) GetAnalytics(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID := httputil.GetUserID(r)

	rangeType := r.URL.Query().Get("range")
	if rangeType == "" {
		rangeType = "today"
	}

	var applianceID *int64
	idStr := r.URL.Query().Get("appliance_id")

	if idStr != "" && idStr != "all" {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid appliance_id", http.StatusBadRequest)
			return
		}
		applianceID = &id
	}

	data, err := h.service.GetAnalytics(
		r.Context(),
		userID,
		applianceID,
		rangeType,
	)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(data)
}
