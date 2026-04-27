package alertapi

import (
	"encoding/json"
	"energy-monitor-server/internal/model/alert"
	httputil "energy-monitor-server/internal/utils/http"
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
