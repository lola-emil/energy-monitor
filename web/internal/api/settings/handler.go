package settings

import (
	"encoding/json"
	"energy-monitor-server/internal/model/setting"
	httputil "energy-monitor-server/internal/utils/http"
	"log"
	"net/http"
)

type SettingsHandler struct {
	service *SettingsService
}

func NewSettingHandler(service *SettingsService) *SettingsHandler {
	return &SettingsHandler{service: service}
}

func (h *SettingsHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID := httputil.GetUserID(r)

	settings, err := h.service.Get(r.Context(), userID)
	if err != nil {
		log.Println(err.Error())
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(settings)
}

func (h *SettingsHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := httputil.GetUserID(r)

	var req setting.Settings
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	req.UserID = userID

	if err := h.service.Update(r.Context(), &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]any{
		"message": "shit",
	}

	json.NewEncoder(w).Encode(response)

}
