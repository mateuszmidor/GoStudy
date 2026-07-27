package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 fired")
	}()
	stop := timer2.Stop() // after stopping, timer2.C will never get signalled
	if stop {
		fmt.Println("Timer2 stopped")
	}

	time.Sleep(3 * time.Second)
}
