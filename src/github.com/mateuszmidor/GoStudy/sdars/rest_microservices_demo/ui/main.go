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

	// expose UI http interface
	adapterUI := NewUIAdapter(&ui)
	ui.SetTunerPort(&adapterUI)

	// run ui
	go ui.Run()
	go adapterUI.RunHTTPServer()

	// wait for INT/TERM
	utils.NewShutdownCondition().Wait()

	// demo done
	fmt.Printf("TunerDemo done\n")
}
