package device

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	exists, err := h.deviceRepo.DeviceExists(body.DeviceSerial, body.ActivationCode)

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
