package offline

import (
	"context"
	"log"
	"time"

	"energy-monitor-server/internal/model/alert"
	"energy-monitor-server/internal/model/appliance"
)

func StartOfflineChecker(
	applianceRepo appliance.ApplianceRepository,
	alertRepo alert.AlertRepository,
) {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			check(
				applianceRepo,
				alertRepo,
			)
		}
	}()
}

func check(
	applianceRepo appliance.ApplianceRepository,
	alertRepo alert.AlertRepository,
) {
	ctx := context.Background()

	appliances, err := applianceRepo.GetOfflineCandidates(
		ctx,
		2, // offline if no reading for 2 mins
	)

	if err != nil {
		log.Println("Offline check failed:", err)
		return
	}

	for _, a := range appliances {
		_ = applianceRepo.MarkOffline(ctx, a.ID)

		_ = alertRepo.Create(ctx, &alert.Alert{
			ApplianceID: &a.ID,
			Type:        alert.AlertTypeOffline,
			Name:        a.Name,
			Severity:    alert.AlertSeverityHigh,
			Message:     "Device is offline",
		})

		log.Printf(
			"Device offline detected: %s",
			a.DeviceCode,
		)
	}
}
