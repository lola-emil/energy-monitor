package alertapi

import (
	"encoding/json"
	"energy-monitor-server/internal/model/alert"
	httputil "energy-monitor-server/internal/utils/http"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AlertHandler struct {
	service *AlertService
}

func NewAlertHandler(service *AlertService) *AlertHandler {
	return &AlertHandler{service: service}
}

func (h *AlertHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := httputil.GetUserID(r)

	filter := alert.AlertFilter{
		Status:      r.URL.Query().Get("status"),
		Severity:    r.URL.Query().Get("severity"),
		ApplianceID: httputil.ParseIntPtr(r.URL.Query().Get("appliance_id")),
	}

	alerts, err := h.service.List(r.Context(), userID, filter)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed", 500)
		return
	}

	json.NewEncoder(w).Encode(alerts)
}

func (h *AlertHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID := httputil.GetUserID(r)
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	a, err := h.service.Get(r.Context(), userID, id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(a)
}

func (h *AlertHandler) Resolve(w http.ResponseWriter, r *http.Request) {
	userID := httputil.GetUserID(r)
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err := h.service.Resolve(r.Context(), userID, id); err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *AlertHandler) GetAnalyticsAlerts(w http.ResponseWriter, r *http.Request) {
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
		id, _ := strconv.ParseInt(idStr, 10, 64)
		applianceID = &id
	}

	data, err := h.service.GetAnalyticsAlerts(
		r.Context(),
		userID,
		applianceID,
		rangeType,
		month, year,
	)

	log.Println(len(data))

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func (h *AlertHandler) GetByAppliance(w http.ResponseWriter, r *http.Request) {
	userID := httputil.GetUserID(r)

	idStr := chi.URLParam(r, "id")
	applianceID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid appliance id", 400)
		return
	}

	data, err := h.service.GetRecentByAppliance(
		r.Context(),
		userID,
		applianceID,
	)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(data)
}
