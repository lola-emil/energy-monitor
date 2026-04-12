package mymqtt

import (
	energyreading "backend/internal/api/energy-reading"
	"backend/internal/event"
	jwtutil "backend/internal/pkg/jwt-util"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang-jwt/jwt/v5"
)

type TopicHandler struct {
	sseBroker *event.Hub
	repo      *Repository
}

func NewTopicHandler(sseBroker *event.Hub, repo *Repository) *TopicHandler {
	if sseBroker == nil {
		panic("sseBroker cannot be nil")
	}
	if repo == nil {
		panic("repo cannot be nil")
	}

	return &TopicHandler{
		repo:      repo,
		sseBroker: sseBroker,
	}
}

func (th *TopicHandler) RegisterDevice(c mqtt.Client, m mqtt.Message) {
	log.Println("Registering Device")
	payload := string(m.Payload())

	var data DeviceRegister
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		log.Println("error unmarshalling JSON:", err)
		return
	}

	var secretCode = "REG456"

	if data.RegistrationCode != secretCode {
		log.Println("Unknown Device")
		return
	}

	device, err := th.repo.DeviceExists(data.DeviceCode)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Exists check", err.Error())
		return
	}

	var deviceID int64

	if device != nil {
		log.Println("Device already exists")
		deviceID = device.ID
	} else {
		id, err := th.repo.SaveDevice(data)
		if err != nil {
			log.Println(err.Error())
			log.Println("Error Saving Device")
			return
		}

		deviceID = id
	}

	token := c.Publish(fmt.Sprintf("device/%s/register/success", data.DeviceCode), 0, false, strconv.Itoa(int(deviceID)))

	token.Wait()

	if token.Error() != nil {
		log.Println("publish error:", token.Error())
		return
	}

	log.Printf("Device %s successfully registered with id %d", data.DeviceCode, deviceID)
}

func (th *TopicHandler) SubEnergyReadinTopic(c mqtt.Client, m mqtt.Message) {
	topic := m.Topic()
	parts := strings.Split(topic, "/")

	if len(parts) < 3 {
		log.Println("invalid topic:", topic)
		return
	}

	deviceIDStr := parts[1]
	deviceID, err := strconv.ParseInt(deviceIDStr, 10, 64)

	if err != nil {
		log.Println("invalid device id in topic:", deviceIDStr)
		return
	}

	payload := string(m.Payload())

	log.Println(payload)

	var sensorData SensorData
	if err := json.Unmarshal([]byte(payload), &sensorData); err != nil {
		log.Println("error unmarshalling JSON:", err)

		resp := map[string]string{
			"status":  "error",
			"message": "invalid_payload",
		}

		data, _ := json.Marshal(resp)
		c.Publish(fmt.Sprintf("device/%d/sensor/response", deviceID), 0, false, data).Wait()
		return
	}

	deviceClaim, err := th.repo.GetDeviceClaimByDeviceId(deviceID)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Saving sensor data...", deviceID)
	body := energyreading.EnergyReadingBody{
		DeviceId: deviceID,
		Voltage:  sensorData.Voltage,
		Current:  sensorData.Current,
		PowerKwh: sensorData.PowerDraw,
	}
	reading, err := th.repo.SaveDeviceReadings(body)
	log.Println("Reading: ", reading)

	if th.sseBroker == nil {
		log.Println("sseBroker is NIL")
	}
	if th.sseBroker != nil && th.sseBroker.Broadcast == nil {
		log.Println("Broadcast channel is NIL")
	}
	if deviceClaim == nil {
		log.Println("deviceClaim is NIL")
	}

	if err != nil {
		log.Println(err.Error())
		return
	}

	th.sseBroker.Broadcast <- event.Event{
		UserID:   deviceClaim.UserId,
		DeviceID: deviceID,
		Value:    body,
	}
}

func (th *TopicHandler) AuthenticateDevice(c mqtt.Client, m mqtt.Message) {
	payload := string(m.Payload())

	fmt.Println(payload)

	topic := m.Topic()
	parts := strings.Split(topic, "/")
	if len(parts) < 3 {
		log.Println("invalid topic:", topic)
		return
	}

	deviceIDStr := parts[1]
	deviceID, err := strconv.ParseInt(deviceIDStr, 10, 64)
	if err != nil {
		log.Println("invalid device id in topic:", deviceIDStr)
		return
	}

	var sensorData DeviceAuth
	if err := json.Unmarshal([]byte(payload), &sensorData); err != nil {
		log.Println("error unmarshalling JSON:", err)

		resp := map[string]string{
			"status":  "error",
			"message": "invalid_payload",
		}

		data, _ := json.Marshal(resp)
		c.Publish(fmt.Sprintf("device/%d/auth/response", deviceID), 0, false, data).Wait()
		return
	}

	device, err := th.repo.GetDeviceByIdAndSerial(deviceID, sensorData.DeviceSerial)
	if err != nil {
		log.Println("device lookup failed:", err)

		resp := map[string]string{
			"status":  "error",
			"message": "device_not_found",
		}

		data, _ := json.Marshal(resp)
		c.Publish(fmt.Sprintf("device/%d/auth/response", deviceID), 0, false, data).Wait()
		return
	}

	// jwt
	jwtToken, err := jwtutil.CreateToken(jwt.MapClaims{
		"device_id": device.ID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
		"iss":       "what-the-fack",
	})

	resp := map[string]any{
		"status":    "success",
		"device_id": device.ID,
		"message":   "device_authenticated",
		"token":     jwtToken,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Println("error marshalling response:", err)
		return
	}

	token := c.Publish(fmt.Sprintf("device/%d/auth/response", device.ID), 0, false, data)
	token.Wait()

	if token.Error() != nil {
		log.Println("publish error:", token.Error())
	}
}
