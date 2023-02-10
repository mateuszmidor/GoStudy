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

	_, err = channel.QueueDeclare("TestQueue", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.Publish("", "TestQueue", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello world!"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Published message to the queue")
}
