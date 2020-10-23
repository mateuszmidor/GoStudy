package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	msgs, err := channel.Consume("TestQueue", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	msg := <-msgs
	fmt.Printf("Received message: %s\n", msg.Body)
}
