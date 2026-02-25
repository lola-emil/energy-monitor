package device

import (
	"encoding/json"
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

func (h *DeviceHandler) TestFunction(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "hello from device module",
	}

	json.NewEncoder(w).Encode(response)
}
