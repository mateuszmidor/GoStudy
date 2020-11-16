package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
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

	fmt.Println("Setting COLOR")
	c.Set("COLOR", "RED", time.Hour)
	fmt.Print("Done\n\n")

	fmt.Println("Getting COLOR")
	value, _ := c.Get("COLOR").Result()
	fmt.Println(value)
	fmt.Print("Done\n\n")
}
