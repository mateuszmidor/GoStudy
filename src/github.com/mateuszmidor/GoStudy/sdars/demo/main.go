package main

import (
	"fmt"
	"time"
	"math/rand"
	"adapters"
	"hexagons/hw"
	"hexagons/tuner"
	"hexagons/ui"
) 

func main() {
	rand.Seed(time.Now().UnixNano())

	// hw side of the communication; sends subscription/station list updates, receives tune requests
	hw := hw.NewHwRoot()

	// ui side of the communication; sends tune requests, receives subscription/station list updates
	ui := ui.NewUiRoot()

	// tuner; the middle side. Does business logic and communicates with hw and ui
	tuner := tuner.NewTunerRoot()

	// communicate hexagons tuner <-> ui, tuner <-> hw
	adapterTunerHw := adapters.NewHwAdapter(&tuner, &hw)
	adapterTunerUi := adapters.NewUiAdapter(&tuner, &ui)

	hw.SetTunerOutPort(&adapterTunerHw)
	tuner.SetupUiPortOut(&adapterTunerUi)
	tuner.SetupHwPortOut(&adapterTunerHw)
	ui.SetupTunerPortOut(&adapterTunerUi)

	// run all the parties
	go hw.Run()
	go tuner.Run()
	go ui.Run()

	// wait for INT/TERM
	NewShutdownCondition().Wait()

	// demo done
	fmt.Printf("TunerDemo done\n")
}