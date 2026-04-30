package energyreading

import (
	"context"
	"energy-monitor-server/internal/model/energyreading"
	"energy-monitor-server/internal/model/setting"
	"energy-monitor-server/internal/services"
)

type ReadingService struct {
	repo         energyreading.ReadingRepository
	settingsRepo setting.SettingsRepository
	alertEngine  *services.AlertEngine
}

func NewReadingService(
	repo energyreading.ReadingRepository,
	settingsRepo setting.SettingsRepository,
	alertEngine *services.AlertEngine,
) *ReadingService {
	return &ReadingService{
		repo:         repo,
		settingsRepo: settingsRepo,
		alertEngine:  alertEngine,
	}
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

func (s *ReadingService) GetSummary(
	ctx context.Context,
	userID int64,
) (*energyreading.ReadingSummary, error) {
	summary, err := s.repo.GetSummary(
		ctx,
		userID,
		nil,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	settings, err := s.settingsRepo.GetByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	summary.EstimatedCost =
		(summary.TotalEnergyKWh * settings.RatePerKWh) +
			settings.FixedMonthlyCharge

	return summary, nil
}

func (s *ReadingService) GetEnergyChart(ctx context.Context, userID int64) ([]energyreading.ChartPoint, error) {
	return s.repo.GetEnergyChart(ctx, userID, "monthly")
}
