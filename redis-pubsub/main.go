package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func getClient() *redis.Client {
	options := redis.Options{
		Password:    "",               // default no password set
		Addr:        "localhost:6379", // default port
		DB:          0,                // default db
		DialTimeout: time.Second * 10,
	}
	c := redis.NewClient(&options)

	// ping to check the connection
	status := c.Ping(context.Background())
	_, err := status.Result()
	if err != nil {
		panic(status)
	}

	return c
}

// golang client method names directly copied from redis shell command names, so see: https://redis.io/commands/#pubsub
func main() {
	fmt.Println("connecting...")
	c := getClient()
	defer c.Close()
	fmt.Print("Done\n\n")

	ctx := context.Background()

	// run subscriber
	go func() {
		ps := c.Subscribe(ctx, "messages")
		for msg := range ps.Channel() {
			fmt.Println(msg.Payload)
		}
	}()

	// publish some messages
	time.Sleep(time.Second) // give subscriber time to get started
	c.Publish(ctx, "messages", "Litwo, ")
	time.Sleep(time.Second)
	c.Publish(ctx, "messages", "ojczyzno ")
	time.Sleep(time.Second)
	c.Publish(ctx, "messages", "moja! ")
	time.Sleep(time.Second) // give subscriber time to read last message

	fmt.Print("Done\n\n")
}
