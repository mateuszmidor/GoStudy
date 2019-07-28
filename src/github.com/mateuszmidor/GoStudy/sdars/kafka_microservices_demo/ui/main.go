package main

import (
	"fmt"
	"hexagons/ui"
	"math/rand"
	"time"
	"utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// ui side of the communication; sends tune requests, receives subscription/station list updates
	ui := ui.NewUiRoot()

	// expose Ui kafka interface
	adapterUi := NewUIAdapter(&ui)
	ui.SetTunerPort(&adapterUi)

	// run ui
	go ui.Run()
	go adapterUi.RunKafkaConsumer()

	// wait for INT/TERM
	utils.NewShutdownCondition().Wait()

	// demo done
	fmt.Printf("Ui done\n")
}
