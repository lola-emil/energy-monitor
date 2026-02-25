package energyreading

import (
	"encoding/json"
	"net/http"
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
