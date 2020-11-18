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

func add(key, val string) string {
	c := getClient()
	defer c.Close()
	ctx := context.Background()

	addArgs := redis.XAddArgs{
		Values: map[string]interface{}{key: val}, // message stores key-value paris, can be many pairs in one message
		ID:     "*",                              // auto-generate ID; IDs allow XRead to specify first message to be read
		Stream: "messages",
	}
	result := c.XAdd(ctx, &addArgs)

	fmt.Println(result)
	return result.Val() // return ID of added message
}

func read(startID string, count int64) {
	c := getClient()
	defer c.Close()
	ctx := context.Background()

	readArgs := redis.XReadArgs{
		Streams: []string{"messages", startID}, // read next id higher than startID
		Count:   count,
	}
	result := c.XRead(ctx, &readArgs)

	fmt.Println(result)
}

func xrange(startID, endID string) {
	c := getClient()
	defer c.Close()
	ctx := context.Background()

	result := c.XRange(ctx, "messages", startID, endID)
	fmt.Println(result)
}

// golang client method names directly copied from redis shell command names, so see: https://redis.io/commands#stream
func main() {
	fmt.Println("connecting...")
	c := getClient()
	c.Close()
	fmt.Print("Done\n\n")

	fmt.Println("Inserting messages into stream")
	idRed := add("color", "RED")
	idGreen := add("color", "GREEN")
	idBlue := add("color", "BLUE")
	fmt.Print("Done\n\n")

	fmt.Println("Reading messages one by one")
	read("0", 1)     // read first message in stream
	read(idRed, 1)   // read message inserted after idRed
	read(idGreen, 1) // read message inserted after idGreen
	fmt.Print("Done\n\n")

	fmt.Println("Reading range of messages")
	xrange(idRed, idBlue)
	fmt.Print("Done\n\n")
}
