package main

import (
	"fmt"
	"log"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	brokerURI = "tcp://localhost:1883"
	topic     = "test/topic"
	username  = "username"
	password  = "pass123"
	clientID  = "publisher"
	message   = "Hello, MQTT!"
	qos       = 2 // Quality of Service level
)

func main() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	log.Println("Publisher connected to MQTT broker")
	counter := 1

	// Publish messages
	for {
		msg := fmt.Sprintf("[%v] %v", counter, message)
		token := client.Publish(topic, qos, false, msg)
		token.Wait()
		log.Printf("Published message: %v\n", msg)
		counter++
		time.Sleep(1 * time.Second) // Sleep for 1 second between each publish
	}
}
