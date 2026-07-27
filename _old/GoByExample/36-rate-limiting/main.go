package main

import (
	"fmt"
	"time"
)

func main() {
	// example 1
	// dripping requestes
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// one token every 200ms
	limiter := time.Tick(200 * time.Millisecond)

	// work!
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// example 2
	// burst of 3 followed by dripping
	burstyLimiter := make(chan time.Time, 3)

	// load it up with 3 tokens
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// new token every 200ms
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// work!
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
