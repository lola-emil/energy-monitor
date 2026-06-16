package mqtt

import (
	"context"
	"encoding/json"
	"energy-monitor-server/internal/alerts"
	"energy-monitor-server/internal/model/appliance"
	"energy-monitor-server/internal/model/energyreading"
	"energy-monitor-server/internal/sse"
	"log"
)

type Handler struct {
	applianceRepo  appliance.ApplianceRepository
	readingRepo    energyreading.ReadingRepository
	broker         *sse.Broker
	alertEvaluator *alerts.Evaluator
}

func NewHandler(
	applianceRepo appliance.ApplianceRepository,
	readingRepo energyreading.ReadingRepository,
	broker *sse.Broker,
	alertEvaluator *alerts.Evaluator,
) *Handler {
	return &Handler{
		applianceRepo:  applianceRepo,
		readingRepo:    readingRepo,
		broker:         broker,
		alertEvaluator: alertEvaluator,
	}
}

func (h *Handler) HandleIncomingReading(
	ctx context.Context,
	payload []byte,
) {
	var msg IncomingReading

	if err := json.Unmarshal(payload, &msg); err != nil {
		log.Println("MQTT invalid payload:", err)
		return
	}

	if msg.DeviceCode == "" {
		log.Println("MQTT missing device_code")
		return
	}

	// Find appliance using protected device code
	applianceData, err := h.applianceRepo.GetByDeviceCode(
		ctx,
		msg.DeviceCode,
	)
	if err != nil {
		log.Println("Appliance not found:", err)
		return
	}

	readingData := &energyreading.EnergyReading{
		ApplianceID: applianceData.ID,
		Timestamp:   msg.Timestamp,

		Voltage:     msg.Voltage,
		Current:     msg.Current,
		Power:       msg.Power,
		EnergyKWh:   msg.EnergyKWh,
		FrequencyHz: msg.FrequencyHz,
	}

	if err := h.readingRepo.Create(
		ctx,
		readingData,
	); err != nil {
		log.Println("Failed saving reading:", err)
		return
	}

	h.alertEvaluator.Evaluate(
		ctx,
		applianceData.UserID,
		applianceData.ID,
		readingData,
	)

	h.broker.Broadcast(map[string]any{
		"type":        "reading_update",
		"device_code": msg.DeviceCode,
		"power":       msg.Power,
		"voltage":     msg.Voltage,
		"current":     msg.Current,
		"energy_kwh":  msg.EnergyKWh,
		"timestamp":   msg.Timestamp,
	})

	// Update appliance heartbeat / online status
	if err := h.applianceRepo.UpdateLastReading(
		ctx,
		applianceData.ID,
		msg.Timestamp,
	); err != nil {
		log.Println("Failed updating appliance status:", err)
	}

	log.Printf(
		"Reading saved: device=%s appliance_id=%d power=%.2fW",
		msg.DeviceCode,
		applianceData.ID,
		msg.Power,
	)
}
