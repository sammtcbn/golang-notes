package main

import (
	"fmt"
    "bytes"
	"time"
    "encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)


var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    var out bytes.Buffer
    json.Indent(&out, msg.Payload(), "", "    ")
    fmt.Printf("%s\n%s\n\n", msg.Topic(), string(out.Bytes()))
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}

func main() {
	var broker = "tcp://test.mosquitto.org:1883"

	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID("go_mqtt_sub_example")
    //options.SetUsername("admin")
    //options.SetPassword("12345")
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "#"
	token = client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s\n", topic)

	for true {
		time.Sleep(time.Second)
	}

	fmt.Println("bye")
}
