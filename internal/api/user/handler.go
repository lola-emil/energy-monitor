package user

import (
	"encoding/json"
	httputil "energy-monitor-server/internal/utils/http"
	"log"
	"net/http"
)

type UserHandler struct {
	service *UserService
}

func (h *UserHandler) UpdateProfile(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID := httputil.GetUserID(r)

	var req UpdateProfileRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := h.service.UpdateProfile(
		r.Context(),
		userID,
		&req,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Profile updated successfully",
	})
}

func (h *UserHandler) GetProfile(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID := httputil.GetUserID(r)
	profile, err := h.service.GetProfile(r.Context(), userID)

	if err != nil {
		log.Println(err.Error())
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(profile)

}
