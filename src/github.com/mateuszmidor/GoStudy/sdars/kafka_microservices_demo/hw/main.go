package main

import (
	"fmt"
	"hexagons/hw"
	"math/rand"
	"time"
	"utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// hw side of the communication; sends subscription/station list updates, receives tune requests
	hw := hw.NewHwRoot()

	// expose HW kafka interface
	adapterHw := NewHwAdapter(&hw)
	hw.SetTunerPort(&adapterHw)

	// run hw
	go hw.Run()
	go adapterHw.RunKafkaConsumer()

	// wait for INT/TERM
	utils.NewShutdownCondition().Wait()

	// demo done
	fmt.Printf("Hw done\n")
}
