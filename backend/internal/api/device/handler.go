package device

import (
	jwtutil "backend/internal/pkg/jwt-util"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type DeviceHandler struct {
	deviceRepo *DeviceRepo
}

func NewDeviceHandler(deviceRepo *DeviceRepo) *DeviceHandler {
	return &DeviceHandler{
		deviceRepo: deviceRepo,
	}
}

func (h *DeviceHandler) AddDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	var body DeviceRequest

	fmt.Printf("Default: %v\n", body)

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid json body", http.StatusInternalServerError)
		return
	}

	exists, err := h.deviceRepo.DeviceExists(body.DeviceCode)

	if err != nil {
		log.Println("Exists check", err.Error())
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "Device already exists", http.StatusBadRequest)
		return
	}

	// hashedActivationCode, err := password.HashPassword(body.ActivationCode, password.DefaultParams)

	// if err != nil {
	// 	log.Println("Activation code", err.Error())
	// 	http.Error(w, "Server error", http.StatusInternalServerError)
	// 	return
	// }

	// body.ActivationCode = hashedActivationCode

	affectedRows, err := h.deviceRepo.SaveDevice(r.Context(), body)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"message":  "Registration successful",
		"affected": affectedRows,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *DeviceHandler) GetDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	devices, err := h.deviceRepo.GetDevices(r.Context())

	if err != nil {
		log.Println("DB error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(devices); err != nil {
		log.Println("Response error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *DeviceHandler) GetUserDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value("claims").(*jwtutil.AccessTokenClaims)

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId := claims.UserID

	devices, err := h.deviceRepo.GetUserDevices(r.Context(), userId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No device added yet", http.StatusNotFound)
		} else {
			log.Println(err.Error())
			http.Error(w, fmt.Sprintf("SQL: %s", err.Error()), http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(devices); err != nil {
		log.Println("Response error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *DeviceHandler) GetUserDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pathDeviceId := r.PathValue("id")
	id, err := strconv.ParseInt(pathDeviceId, 10, 64)

	fmt.Println("Bwesit")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	claims, ok := r.Context().Value("claims").(*jwtutil.AccessTokenClaims)

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	devices, err := h.deviceRepo.GetUserDevice(r.Context(), claims.UserID, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No device added yet", http.StatusNotFound)
		} else {
			log.Println(err.Error())
			http.Error(w, fmt.Sprintf("SQL: %s", err.Error()), http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(devices); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
