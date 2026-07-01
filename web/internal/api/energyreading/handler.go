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
	ctx := r.Context()
	userID := httputil.GetUserID(r)

	rangeType := r.URL.Query().Get("range")
	if rangeType == "" {
		rangeType = "today"
	}

	var month *int

	monthStr := r.URL.Query().Get("month")
	if monthStr != "" {
		m, err := strconv.Atoi(monthStr)
		if err != nil {
			// handle invalid month
			return
		}
		month = &m
	}

	var year *int

	yearStr := r.URL.Query().Get("year")
	if yearStr != "" {
		y, err := strconv.Atoi(yearStr)
		if err != nil {
			// handle invalid year
			return
		}
		year = &y
	}

	var applianceID *int64
	if idStr := r.URL.Query().Get("appliance_id"); idStr != "" && idStr != "all" {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid appliance_id", http.StatusBadRequest)
			return
		}
		applianceID = &id
	}

	summary, err := h.service.GetSummary(
		ctx,
		userID,
		applianceID,
		rangeType,
		month, year,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

func (h *ReadingHandler) GetChart(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID := httputil.GetUserID(r)

	rangeType := r.URL.Query().Get("range")
	if rangeType == "" {
		rangeType = "today"
	}

	data, err := h.service.GetEnergyChart(
		r.Context(),
		userID,
		rangeType,
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

	var month *int

	monthStr := r.URL.Query().Get("month")
	if monthStr != "" {
		m, err := strconv.Atoi(monthStr)
		if err != nil {
			// handle invalid month
			return
		}
		month = &m
	}

	var year *int

	yearStr := r.URL.Query().Get("year")
	if yearStr != "" {
		y, err := strconv.Atoi(yearStr)
		if err != nil {
			// handle invalid year
			return
		}
		year = &y
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
		month,
		year,
	)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(data)
}
func (h *ReadingHandler) GetDetailedReadings(w http.ResponseWriter, r *http.Request) {
	userID := httputil.GetUserID(r)

	rangeType := r.URL.Query().Get("range")
	if rangeType == "" {
		rangeType = "today"
	}

	var month *int

	monthStr := r.URL.Query().Get("month")
	if monthStr != "" {
		m, err := strconv.Atoi(monthStr)
		if err != nil {
			// handle invalid month
			return
		}
		month = &m
	}

	var year *int

	yearStr := r.URL.Query().Get("year")
	if yearStr != "" {
		y, err := strconv.Atoi(yearStr)
		if err != nil {
			// handle invalid year
			return
		}
		year = &y
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	if pageSize <= 0 {
		pageSize = 10
	}

	var applianceID *int64
	if idStr := r.URL.Query().Get("appliance_id"); idStr != "" && idStr != "all" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		applianceID = &id
	}

	data, total, err := h.service.GetDetailedReadings(
		r.Context(),
		userID,
		applianceID,
		rangeType,
		page,
		pageSize,
		month, year,
	)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"data":  data,
		"total": total,
	})
}
