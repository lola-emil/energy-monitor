package offline

import (
	"context"
	"log"
	"time"

	"energy-monitor-server/internal/model/appliance"
)

func StartOfflineChecker(
	applianceRepo appliance.ApplianceRepository,
) {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			check(
				applianceRepo,
			)
		}
	}()
}

func check(
	applianceRepo appliance.ApplianceRepository,
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

		log.Printf(
			"Device offline detected: %s",
			a.DeviceCode,
		)
	}
}
