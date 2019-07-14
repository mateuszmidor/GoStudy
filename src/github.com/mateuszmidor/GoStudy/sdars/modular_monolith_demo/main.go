package main

import (
	"fmt"
	"hexagons/hw"
	"hexagons/tuner"
	"hexagons/ui"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// hw side of the communication; sends subscription/station list updates, receives tune requests
	hw := hw.NewHwRoot()

	// ui side of the communication; sends tune requests, receives subscription/station list updates
	ui := ui.NewUiRoot()

	// tuner; the middle side. Does business logic and communicates with hw and ui
	tuner := tuner.NewTunerRoot()

	// communicate hexagons hw <-> tuner <-> ui
	adapterTunerHW := NewTunerHwAdapter(&tuner, &hw)
	hw.SetTunerPort(&adapterTunerHW)
	tuner.SetHwPort(&adapterTunerHW)

	adapterTunerUI := NewTunerUIAdapter(&tuner, &ui)
	tuner.SetUiPort(&adapterTunerUI)
	ui.SetTunerPort(&adapterTunerUI)

	// run all the parties
	go hw.Run()
	go tuner.Run()
	go ui.Run()

	// wait for INT/TERM
	NewShutdownCondition().Wait()

	// demo done
	fmt.Printf("TunerDemo done\n")
}
