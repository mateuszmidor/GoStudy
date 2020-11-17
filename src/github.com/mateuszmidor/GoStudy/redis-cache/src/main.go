package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func getClient() *redis.Client {
	options := redis.Options{
		Addr:        "localhost:6379",
		Password:    "",
		DB:          0, // default
		DialTimeout: time.Second * 10,
	}
	return redis.NewClient(&options)
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
