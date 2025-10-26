// Project: ping-pong between 2 goroutines
// Usage: go run .
// ~7 mln pingpongs on 12 vcpu i7, 8gb ram
package main

import (
	"fmt"
	"time"
)

var counter uint64

type chantype int

var value chantype = 42
var a = make(chan chantype)
var b = make(chan chantype)

func main() {

	go ping()
	go pong()

	a <- value
	time.Sleep(1 * time.Second)
	fmt.Printf("in 1 second did %d ping-pongs\n", counter)
}

func ping() {
	for {
		v := <-a
		b <- v
		counter++
	}
}

func pong() {
	for {
		v := <-b
		a <- v
		counter++
	}
}
