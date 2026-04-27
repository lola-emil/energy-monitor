package mqtt

import (
	"context"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartSubscriber(
	handler *Handler,
) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("energy-monitor-server")

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	topic := "energy/readings/+"

	token := client.Subscribe(topic, 1, func(
		client mqtt.Client,
		msg mqtt.Message,
	) {
		handler.HandleIncomingReading(
			context.Background(),
			msg.Payload(),
		)
	})

	token.Wait()

	log.Println("MQTT subscriber started:", topic)
}
