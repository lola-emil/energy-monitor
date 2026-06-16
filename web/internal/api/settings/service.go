package settings

import (
	"context"
	"energy-monitor-server/internal/model/setting"
	"errors"
)

type SettingsService struct {
	repo setting.SettingsRepository
}

func NewSettingService(repo setting.SettingsRepository) *SettingsService {
	return &SettingsService{repo: repo}
}
func (s *SettingsService) Get(ctx context.Context, userID int64) (*setting.Settings, error) {
	return s.repo.GetByUserID(ctx, userID)
}
func (s *SettingsService) Update(
	ctx context.Context,
	settings *setting.Settings,
) error {
	// Basic validation

	if settings.RatePerKWh < 0 {
		return errors.New("rate_per_kwh must be greater than or equal to 0")
	}

	if settings.FixedMonthlyCharge < 0 {
		return errors.New("fixed_monthly_charge must be greater than or equal to 0")
	}

	if settings.RefreshIntervalSeconds <= 0 {
		return errors.New("refresh_interval_seconds must be greater than 0")
	}

	if settings.EnableVoltageAlerts {
		if settings.OverVoltageThreshold <= 0 {
			return errors.New("over_voltage_threshold must be greater than 0")
		}

		if settings.UnderVoltageThreshold <= 0 {
			return errors.New("under_voltage_threshold must be greater than 0")
		}

		if settings.OverVoltageThreshold <= settings.UnderVoltageThreshold {
			return errors.New("over voltage threshold must be greater than under voltage threshold")
		}
	}

	if settings.EnableCurrentAlerts {
		if settings.OverCurrentThreshold <= 0 {
			return errors.New("over_current_threshold must be greater than 0")
		}
	}

	return s.repo.Upsert(ctx, settings)
}
