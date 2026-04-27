package energyreading

import (
	"context"
	"energy-monitor-server/internal/model/energyreading"
	"energy-monitor-server/internal/services"
)

type ReadingService struct {
	repo        energyreading.ReadingRepository
	alertEngine services.AlertEngine
}

func (s *ReadingService) Create(
	ctx context.Context,
	userID int64,
	reading *energyreading.EnergyReading,
) error {
	if err := s.repo.Create(ctx, reading); err != nil {
		return err
	}

	_ = s.repo.UpdateApplianceLastReading(ctx, reading.ApplianceID)

	return s.alertEngine.ProcessReading(ctx, userID, reading)
}
