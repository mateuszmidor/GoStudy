package main

import (
	"log"
	"time"
)

const message = "Hello, world"

func main() {
	var counter int
	for {
		log.Println(counter, message)
		counter++
		time.Sleep(time.Second)
	}
}
