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

func (h *DashboardHandler) GetMonthlyAvgPower(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	monthlyEnergyConsumed, err := h.repo.GetMonthlyAvgPower(2026)

	if err != nil {
		log.Println("DB error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(monthlyEnergyConsumed); err != nil {
		log.Println("Response error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
