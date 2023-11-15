package main

import (
	"fmt"
	"time"
)

var counter int = 0

func main() {
	go incrementCounter()
	go incrementCounter()
	time.Sleep(time.Second * 10)
	fmt.Println(counter)
}

func incrementCounter() {
	for i := 0; i < 10000; i++ {
		counter++
	}
}
