package applianceapi

import (
	"encoding/json"
	"energy-monitor-server/internal/model/appliance"
	httputil "energy-monitor-server/internal/utils/http"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ApplianceHandler struct {
	service *ApplianceService
}

func NewApplianceHandler(service *ApplianceService) *ApplianceHandler {
	return &ApplianceHandler{service: service}
}

func (h *ApplianceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name       string `json:"name"`
		Location   string `json:"location"`
		DeviceCode string `json:"device_code"`
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	a := &appliance.Appliance{
		UserID:     httputil.GetUserID(r),
		Name:       req.Name,
		Location:   req.Location,
		DeviceCode: req.DeviceCode,
	}

	if err := h.service.Create(r.Context(), a); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(a)
}

func (h *ApplianceHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.service.List(r.Context(), httputil.GetUserID(r))
	if err != nil {
		http.Error(w, "failed", 500)
		return
	}

	json.NewEncoder(w).Encode(list)
}

func (h *ApplianceHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	a, err := h.service.Get(r.Context(), id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(a)
}

func (h *ApplianceHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	var req struct {
		Name     string `json:"name"`
		Location string `json:"location"`
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	a := &appliance.Appliance{
		ID:       id,
		UserID:   httputil.GetUserID(r),
		Name:     req.Name,
		Location: req.Location,
	}

	if err := h.service.Update(r.Context(), a); err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ApplianceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err := h.service.Delete(r.Context(), httputil.GetUserID(r), id); err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ApplianceHandler) GetStatus(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID := httputil.GetUserID(r)

	data, err := h.service.GetStatus(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(data)
}
