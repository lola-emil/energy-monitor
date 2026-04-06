package energyreading

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type EnergyReadingHandler struct {
	energyReadingRepo *EnergyReadingRepo
}

func NewEnergyReadingHandler(energyReadingRepo *EnergyReadingRepo) *EnergyReadingHandler {
	return &EnergyReadingHandler{
		energyReadingRepo: energyReadingRepo,
	}
}

func (h *EnergyReadingHandler) TestFunction(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "hello from energy-reading module",
	}

	json.NewEncoder(w).Encode(response)
}

func (h *EnergyReadingHandler) GetEnergyReadings(w http.ResponseWriter, r *http.Request) {
	readings, err := h.energyReadingRepo.GetEnergyReadings(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(readings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *EnergyReadingHandler) GetEnergyReading(w http.ResponseWriter, r *http.Request) {
	pathReadingId := r.PathValue("id")
	id, err := strconv.ParseInt(pathReadingId, 10, 64)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	reading, err := h.energyReadingRepo.GetEnergyReadingById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(reading); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
