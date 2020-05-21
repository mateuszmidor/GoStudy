package main

import "time"

func worker(done chan bool) {
	println("working..")
	time.Sleep(time.Second)
	println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
