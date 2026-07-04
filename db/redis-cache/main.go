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

// golang client method names directly copied from redis shell command names, so see: https://redis.io/commands/#string
func main() {
	fmt.Println("connecting...")
	c := getClient()
	defer c.Close()
	fmt.Print("Done\n\n")

	ctx := context.Background()

	// =================== STRINGS
	fmt.Println("Setting COLOR to RED")
	c.Set(ctx, "COLOR", "RED", time.Hour)
	// c.MSet(ctx, "key1", "value1", "key2", "value2") // set multiple key-value pairs in one call
	fmt.Print("Done\n\n")

	fmt.Println("Getting COLOR")
	value, _ := c.Get(ctx, "COLOR").Result()
	fmt.Println(value)
	fmt.Print("Done\n\n")

	fmt.Println("Setting PRICE to 19")
	c.Set(ctx, "PRICE", 19, time.Hour)
	fmt.Print("Done\n\n")

	fmt.Println("Incrementing PRICE")
	c.Incr(ctx, "PRICE")
	fmt.Print("Done\n\n")

	fmt.Println("Getting PRICE")
	value, _ = c.Get(ctx, "PRICE").Result()
	fmt.Println(value)
	fmt.Print("Done\n\n")

	// c.Append(ctx, "helloGreeting", " World") // append string to already existing string

	// =================== LISTS - allow left/right adding/deleting
	fmt.Println("Inserting bowl, pot, pan into list utensils")
	c.LPush(ctx, "utensils", "bowl", "pot", "pan") // lpush adds to the beginning of the list, so the order will be reversed. See also: rpush
	fmt.Print("Done\n\n")

	fmt.Println("Getting utensils list items")
	utensils := c.LRange(ctx, "utensils", 0, -1) // -1 = last item in the list
	fmt.Println(utensils.Val())
	fmt.Print("Done\n\n")

	fmt.Println("Getting utensils list length")
	len := c.LLen(ctx, "utensils")
	fmt.Println(len.Val())
	fmt.Print("Done\n\n")

	// c.LPop(ctx, "utensils") // remove left
	// c.RPop(ctx, "utensils") // remove right

	// =================== SETS - allow checking if element is in set, evaluating set union, intersection, moving element between sets
	fmt.Println("Inserting ford, bmw, honda into set cars")
	c.SAdd(ctx, "cars", "ford", "bmw", "honda")
	fmt.Print("Done\n\n")

	fmt.Println("Getting cars set members")
	cars := c.SMembers(ctx, "cars")
	fmt.Println(cars.Val())
	fmt.Print("Done\n\n")

	fmt.Println("Checking honda is in cars")
	hasHonda := c.SIsMember(ctx, "cars", "honda").Val()
	fmt.Println(hasHonda)
	fmt.Print("Done\n\n")

	// =================== SORTED SETS - member has associated score and sorted by that score; good for scoreboards eg in massive multiplayer games
	// c.ZAdd(...), c.ZRange(...)

	// =================== HSETS - can assign multiple key-values to key. Like structs: Andrzej (age: 33, email: dexlab@o2.pl)
	// c.HSet(...), c.HGet(...), c.HGetAll()

	// save changes
	fmt.Println("Saving changes to disk")
	c.Save(ctx)
	fmt.Print("Done\n\n")

	// clear database contents
	fmt.Println("Flushing database")
	c.FlushAll(ctx)
	fmt.Print("Done\n\n")
}
