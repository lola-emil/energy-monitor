package services

import (
	"context"
	"energy-monitor-server/internal/model/alert"
	"energy-monitor-server/internal/model/appliance"
	"energy-monitor-server/internal/model/energyreading"
	"energy-monitor-server/internal/model/setting"
	"fmt"
)

type AlertEngine struct {
	settingsRepo  setting.SettingsRepository
	alertRepo     alert.AlertRepository
	applianceRepo appliance.ApplianceRepository
}

func NewAlertEngine(
	settingsRepo setting.SettingsRepository,
	alertRepo alert.AlertRepository,
	applianceRepo appliance.ApplianceRepository,
) *AlertEngine {
	return &AlertEngine{
		settingsRepo:  settingsRepo,
		alertRepo:     alertRepo,
		applianceRepo: applianceRepo,
	}
}

func (s *AlertEngine) ProcessReading(
	ctx context.Context,
	userID int64,
	reading *energyreading.EnergyReading,
) error {
	settings, err := s.settingsRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	// OVER VOLTAGE
	if settings.EnableVoltageAlerts &&
		reading.Voltage > settings.OverVoltageThreshold {

		s.createIfNotExists(
			ctx,
			reading.ApplianceID,
			alert.AlertTypeOverVoltage,
			alert.AlertSeverityHigh,
			"Voltage exceeded threshold",
		)
	}

	// UNDER VOLTAGE
	if settings.EnableVoltageAlerts &&
		reading.Voltage < settings.UnderVoltageThreshold {

		s.createIfNotExists(
			ctx,
			reading.ApplianceID,
			alert.AlertTypeUnderVoltage,
			alert.AlertSeverityMedium,
			"Voltage dropped below threshold",
		)
	}

	// OVER CURRENT
	if settings.EnableCurrentAlerts &&
		reading.Current > settings.OverCurrentThreshold {

		s.createIfNotExists(
			ctx,
			reading.ApplianceID,
			alert.AlertTypeOverCurrent,
			alert.AlertSeverityHigh,
			"Current exceeded threshold",
		)
	}

	return nil
}

func (s *AlertEngine) createIfNotExists(
	ctx context.Context,
	applianceID int64,
	alertType alert.AlertType,
	severity alert.AlertSeverity,
	message string,
) {

	device, err := s.applianceRepo.GetByID(ctx, applianceID)

	if err != nil {
		fmt.Println(err)
		return
	}

	exists, err := s.alertRepo.HasActiveAlert(
		ctx,
		applianceID,
		alertType,
	)
	if err != nil || exists {
		fmt.Println(err)
		return
	}

	_ = s.alertRepo.Create(ctx, &alert.Alert{
		ApplianceID: &applianceID,
		Name:        &device.Name,
		Type:        alertType,
		Severity:    severity,
		Message:     message,
	})
}
