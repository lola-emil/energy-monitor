package mymqtt

import (
	"backend/internal/ws"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jmoiron/sqlx"
)

func StartMQTT(wsHub *ws.WSHub, db *sqlx.DB) mqtt.Client {

	opts := mqtt.NewClientOptions().
		AddBroker("tcp://127.0.0.1:1883").
		SetClientID("go_service").
		SetAutoReconnect(true)

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Println("MQTT connection lost:", err)
	}

	client := mqtt.NewClient(opts)

	token := client.Connect()
	token.Wait()

	if token.Error() != nil {
		log.Fatal(token.Error())
	}

	log.Println("MQTT connected")

	// DEFINE ANG MGA HANDLER
	repo := NewRepository(db)
	topicHandler := NewTopicHandler(wsHub, repo)

	token = client.Subscribe("device/+/auth", 0, topicHandler.AuthenticateDevice)
	token = client.Subscribe("device/+/sensor", 0, topicHandler.SubEnergyReadinTopic)

	token.Wait()
	if token.Error() != nil {
		log.Println("Subscribe error:", token.Error())
		return client
	}

	log.Println("Subscribed to test/topic")

	return client
}
