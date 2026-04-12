package deviceclaim

import (
	"backend/internal/api/device"
	jwtutil "backend/internal/pkg/jwt-util"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type DeviceClaimHandler struct {
	repo       *DeviceClaimRepo
	deviceRepo *device.DeviceRepo
}

func NewDeviceClaimHandler(repo *DeviceClaimRepo, deviceRepo *device.DeviceRepo) *DeviceClaimHandler {
	return &DeviceClaimHandler{
		repo:       repo,
		deviceRepo: deviceRepo,
	}
}

func (h *DeviceClaimHandler) GetDeviceClaims(w http.ResponseWriter, r *http.Request) {
	pathDeviceId := r.PathValue("device-id")
	id, err := strconv.ParseInt(pathDeviceId, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No claims yet", http.StatusBadRequest)
		} else {
			http.Error(w, fmt.Sprintf("SQL: %s", err.Error()), http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(claim); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *DeviceClaimHandler) ClaimDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	var body DeviceClaimRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusInternalServerError)
		return
	}

	device, err := h.deviceRepo.GetDeviceByCode(body.DeviceCode)

	if err != nil {
		http.Error(w, "Invalid device", http.StatusBadRequest)
		return
	}

	val := r.Context().Value("claims")
	if val == nil {
		http.Error(w, "missing claims in context", http.StatusUnauthorized)
		return
	}

	claims, ok := val.(*jwtutil.AccessTokenClaims)
	if !ok {
		log.Printf("claims value: %#v\n", r.Context().Value("claims"))

		http.Error(w, "invalid claims type", http.StatusUnauthorized)
		return
	}

	if !ok {
		http.Error(w, "Error ang claims", http.StatusInternalServerError)
		return
	}

	userId := claims.UserID

	deviceTaken, err := h.repo.DeviceAlreadyTaken(r.Context(), device.ID, userId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if deviceTaken {
		http.Error(w, "Device Taken", http.StatusConflict)
		return
	}

	newClaim := DeviceClaim{
		DeviceId:   device.ID,
		UserId:     userId,
		DeviceName: body.DeviceName,
	}

	claimId, err := h.repo.ClaimDevice(r.Context(), newClaim)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error claiming device", http.StatusInternalServerError)
	}

	response := map[string]any{
		"status":   "success",
		"message":  "Claim successful",
		"claim_id": claimId,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
