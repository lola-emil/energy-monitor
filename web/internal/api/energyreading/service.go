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
	applianceID *int64,
	rangeType string,
) (*energyreading.ReadingSummary, error) {

	summary, err := s.repo.GetSummary(
		ctx,
		userID,
		applianceID,
		rangeType,
	)
	if err != nil {
		return nil, err
	}

	settings, err := s.settingsRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if rangeType == "today" {
		summary.EstimatedCost =
			(summary.TotalEnergyKWh * settings.RatePerKWh) +
				(settings.FixedMonthlyCharge / 30.0)
	} else {
		summary.EstimatedCost =
			(summary.TotalEnergyKWh * settings.RatePerKWh) +
				settings.FixedMonthlyCharge
	}

	summary.BillingRate = settings.RatePerKWh

	return summary, nil
}
func (s *ReadingService) GetEnergyChart(
	ctx context.Context,
	userID int64,
	rangeType string,
) ([]energyreading.ChartPoint, error) {
	return s.repo.GetEnergyChart(ctx, userID, rangeType)
}

func (s *ReadingService) GetAnalytics(
	ctx context.Context,
	userID int64,
	applianceID *int64,
	rangeType string,
) (*energyreading.AnalyticsResponse, error) {

	// Summary
	summary, err := s.repo.GetAnalyticsSummary(
		ctx,
		userID,
		applianceID,
		rangeType,
	)
	if err != nil {
		return nil, err
	}

	// 2. Energy (timestamp-based integration)
	energy, err := s.repo.GetAnalyticsEnergyChart(
		ctx,
		userID,
		applianceID,
		rangeType,
	)
	if err != nil {
		return nil, err
	}

	// OPTIONAL: compute total energy for summary
	var totalEnergy float64
	for _, e := range energy {
		totalEnergy += e.Value
	}
	summary.TotalEnergyKWh = totalEnergy

	// 3. Voltage & Current
	vc, err := s.repo.GetVoltageCurrentChart(
		ctx,
		userID,
		applianceID,
		rangeType,
	)
	if err != nil {
		return nil, err
	}

	return &energyreading.AnalyticsResponse{
		Summary:        *summary,
		Energy:         energy,
		VoltageCurrent: vc,
	}, nil
}

func (s *ReadingService) GetDetailedReadings(
	ctx context.Context,
	userID int64,
	applianceID *int64,
	rangeType string,
	page int,
	pageSize int,
) ([]energyreading.EnergyReading, int, error) {

	offset := (page - 1) * pageSize

	return s.repo.GetDetailedReadings(
		ctx,
		userID,
		applianceID,
		rangeType,
		pageSize,
		offset,
	)
}
