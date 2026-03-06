package mymqtt

import (
	"backend/internal/ws"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type TopicHandler struct {
	wsHub *ws.WSHub
}

func NewTopicHandler(wsHub *ws.WSHub) *TopicHandler {
	return &TopicHandler{
		wsHub: wsHub,
	}
}

func (th *TopicHandler) SubEnergyReadinTopic(c mqtt.Client, m mqtt.Message) {
	payload := string(m.Payload())

	log.Println("MQTT:", payload)

	th.wsHub.Broadcast <- []byte(payload)
}
