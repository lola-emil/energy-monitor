package alerts

import (
	"context"
	"energy-monitor-server/internal/model/alert"
	"energy-monitor-server/internal/model/energyreading"
	"energy-monitor-server/internal/model/setting"
	"log"
)

type Evaluator struct {
	alertRepo    alert.AlertRepository
	settingsRepo setting.SettingsRepository
}

func NewEvaluator(
	alertRepo alert.AlertRepository,
	settingsRepo setting.SettingsRepository,
) *Evaluator {
	return &Evaluator{
		alertRepo:    alertRepo,
		settingsRepo: settingsRepo,
	}
}

func (e *Evaluator) Evaluate(
	ctx context.Context,
	userID int64,
	applianceID int64,
	r *energyreading.EnergyReading,
) {
	settings, err := e.settingsRepo.GetByUserID(
		ctx,
		userID,
	)
	log.Println("Evaluating")
	log.Println(settings)
	if err != nil {
		log.Println("Failed loading settings:", err)
		return
	}

	log.Printf("Settings: %+v\n", settings)
	log.Printf("Reading: %+v\n", r)

	// Over Voltage
	if settings.EnableVoltageAlerts &&
		r.Voltage > settings.OverVoltageThreshold {

		e.createAlert(
			ctx,
			applianceID,
			alert.AlertTypeOverVoltage,
			alert.AlertSeverityHigh,
			"Over voltage detected",
		)
	}

	// Under Voltage
	if settings.EnableVoltageAlerts &&
		r.Voltage < settings.UnderVoltageThreshold {

		e.createAlert(
			ctx,
			applianceID,
			alert.AlertTypeUnderVoltage,
			alert.AlertSeverityMedium,
			"Under voltage detected",
		)

		log.Println("Under voltage detected")

	}

	// Over Current
	if settings.EnableCurrentAlerts &&
		r.Current > settings.OverCurrentThreshold {

		e.createAlert(
			ctx,
			applianceID,
			alert.AlertTypeOverCurrent,
			alert.AlertSeverityHigh,
			"Over current detected",
		)

		log.Println("Over current detected")
	}
}

func (e *Evaluator) createAlert(
	ctx context.Context,
	applianceID int64,
	alertType alert.AlertType,
	severity alert.AlertSeverity,
	message string,
) {
	err := e.alertRepo.Create(ctx, &alert.Alert{
		ApplianceID: &applianceID,
		Type:        alertType,
		Severity:    severity,
		Message:     message,
	})

	if err != nil {
		log.Println("Failed creating alert:", err)
	}
}
