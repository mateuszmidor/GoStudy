package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func getClient() *redis.Client {
	options := redis.Options{
		Password:    "mypass",         // password specified in conf/redis.conf; by default there is no password set
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

func main() {
	fmt.Println("connecting...")
	c := getClient()
	defer c.Close()
	fmt.Print("Done\n\n")

	ctx := context.Background()

	fmt.Println("Setting COLOR")
	c.Set(ctx, "COLOR", "RED", time.Hour)
	fmt.Print("Done\n\n")

	fmt.Println("Getting COLOR")
	value, _ := c.Get(ctx, "COLOR").Result()
	fmt.Println(value)
	fmt.Print("Done\n\n")
}
