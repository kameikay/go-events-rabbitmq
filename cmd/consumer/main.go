package main

import (
	"fmt"

	"github.com/kameikay/events-golang/pkg/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp091.Delivery)
	go rabbitmq.Consume(ch, msgs)
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
