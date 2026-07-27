package main

import (
	"fmt"
	"time"
)

func startTimer(name string) func() {
	t := time.Now()
	return func() {
		d := time.Now().Sub(t)
		fmt.Printf("%s took %fs\n", name, d.Seconds())
	}
}

func main() {
	stop := startTimer("func main()")
	defer stop()

	time.Sleep(1 * time.Second)
}
