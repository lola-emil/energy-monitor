package dashboard

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type DashboardHandler struct {
	repo *DashboardRepo
}

func NewDashboardHandler(repo *DashboardRepo) *DashboardHandler {
	return &DashboardHandler{
		repo: repo,
	}
}

func (h *DashboardHandler) GetOverview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dateStr := r.URL.Query().Get("month")

	layout := "2006-01-02"

	t, err := time.Parse(layout, dateStr)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	overview, err := h.repo.GetOverview(t)

	if err != nil {
		log.Println("DB error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(overview); err != nil {
		log.Println("Response error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *DashboardHandler) GetMonthlyConsumption(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	monthlyConsumption, err := h.repo.GetMonthyEnergyConsumption(r.Context())

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(monthlyConsumption); err != nil {
		log.Println("Response error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
