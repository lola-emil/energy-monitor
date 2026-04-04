package mymqtt

import (
	jwtutil "backend/internal/pkg/jwt-util"
	"backend/internal/ws"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang-jwt/jwt/v5"
)

type TopicHandler struct {
	wsHub *ws.WSHub
	repo  *Repository
}

func NewTopicHandler(wsHub *ws.WSHub, repo *Repository) *TopicHandler {
	return &TopicHandler{
		wsHub: wsHub,
		repo:  repo,
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

	exists, err := th.repo.DeviceExists(data.DeviceCode)

	if err != nil {
		log.Println("Exists check", err.Error())
		return
	}

	if exists {
		log.Println("Device already exists")
		return
	}

	id, err := th.repo.SaveDevice(data)

	if err != nil {
		log.Println(err.Error())
		log.Println("Error Saving Device")
		return
	}

	responsePayload := map[string]int64{
		"device_id": id,
	}

	resPayload, _ := json.Marshal(responsePayload)

	token := c.Publish(fmt.Sprintf("device/%s/register/response", data.DeviceCode), 0, false, resPayload)

	token.Wait()

	if token.Error() != nil {
		log.Println("publish error:", token.Error())
		return
	}
	log.Printf("Device %s successfully registered with id %d", data.DeviceCode, id)
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

	log.Println(sensorData)

	// VERIFY ANG TOKEN
	// claims, err := jwtutil.VerifyToken(sensorData.Token)

	// if err != nil {
	// 	resp := map[string]string{
	// 		"message": "Unauthorized",
	// 	}

	// 	data, _ := json.Marshal(resp)
	// 	c.Publish(fmt.Sprintf("device/%d/sensor/response", deviceID), 0, false, data).Wait()
	// 	return
	// }

	// // SAVE ANG DATA SA DB
	// if v, ok := claims["device_id"].(float64); ok {
	log.Println("Saving sensor data...", deviceID)
	body := EnergyReadingBody{
		DeviceId: deviceID,
		Voltage:  sensorData.Voltage,
		Current:  sensorData.Current,
		PowerKwh: sensorData.PowerKwh,
	}
	reading, err := th.repo.SaveDeviceReadings(body)
	log.Println("Reading: ", reading)

	if err != nil {
		resp := map[string]string{
			"message": "SHIT",
		}

		log.Println(err.Error())

		data, _ := json.Marshal(resp)
		c.Publish(fmt.Sprintf("device/%d/sensor/response", deviceID), 0, false, data).Wait()
		return
	}
	// }

	// log.Println("MQTT:", payload)

	th.wsHub.Broadcast <- []byte(payload)
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
