package main

import "time"

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result1"
	}()

	select {
	case res := <-c1:
		println(res)
	case <-time.After(time.Second):
		println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		println(res)
	case <-time.After(3 * time.Second):
		println("timeout 2")
	}
}
