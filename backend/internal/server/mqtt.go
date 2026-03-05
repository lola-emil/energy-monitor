package server

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartMQTT(wsHub *WSHub) mqtt.Client {

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

	token = client.Subscribe("#", 0, func(c mqtt.Client, m mqtt.Message) {
		payload := string(m.Payload())

		log.Println("MQTT:", payload)

		wsHub.broadcast <- []byte(payload)
	})

	token.Wait()
	if token.Error() != nil {
		log.Println("Subscribe error:", token.Error())
		return client
	}

	log.Println("Subscribed to test/topic")

	return client
}
