package main

import (
	"fmt"
	"log"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	brokerURI = "tcp://localhost:1883"
	topic     = "test/topic"
	username  = "username"
	password  = "pass123"
	clientID  = "subscriber2"
	qos       = 2
)

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	log.Printf("Subscriber 2 received message: %s\n", message.Payload())
}

func main() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.OnConnect = func(client MQTT.Client) {
		if token := client.Subscribe(topic, qos, onMessageReceived); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	log.Println("Subscriber 2 connected to MQTT broker")

	// Keep the subscriber alive
	select {}
}
