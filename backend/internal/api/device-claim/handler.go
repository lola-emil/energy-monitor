package deviceclaim

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type DeviceClaimHandler struct {
	repo *DeviceClaimRepo
}

func NewDeviceClaimHandler(repo *DeviceClaimRepo) *DeviceClaimHandler {
	return &DeviceClaimHandler{
		repo: repo,
	}
}

func (h *DeviceClaimHandler) GetDeviceClaims(w http.ResponseWriter, r *http.Request) {
	pathDeviceId := r.PathValue("device-id")
	id, err := strconv.ParseInt(pathDeviceId, 10, 64)

	claims, err := h.repo.GetDeviceClaims(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(claims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *DeviceClaimHandler) GetDeviceClaim(w http.ResponseWriter, r *http.Request) {
	pathClaimId := r.PathValue("claim-id")
	id, err := strconv.ParseInt(pathClaimId, 10, 64)

	claim, err := h.repo.GetDeviceClaimById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(claim); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
