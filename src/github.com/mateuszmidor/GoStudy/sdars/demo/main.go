package main

import (
	"fmt"
	"time"
	"math/rand"
	"actors/ui"
	"actors/hardware"
	"adapters"
	"hexagons/tuner"
) 

func main() {
	rand.Seed(time.Now().UnixNano())

	// hardware side of the communication; sends station list updates, receives station tune requests
	hardware:= hardware.NewHwActor()

	// ui side of the communication; sends tune requests, receives station list updates
	ui := ui.NewUiActor()

	// aggregate root
	tuner := tuner.NewTunerRoot()

	// hardware talks to tuner
	hwAdapter := adapters.NewHardwareAdapter(&tuner, &hardware)

	// ui talks to tuner
	uiAdapter := adapters.NewUiAdapter(&tuner, &ui)

	// tuner talks to hardware and ui
	tuner.SetupPorts(hwAdapter, uiAdapter)

	// run all the parties
	go ui.Run()
	go tuner.Run()
	go hardware.Run()

	// wait for INT/TERM
	NewShutdownCondition().Wait()

	// demo done
	fmt.Printf("TunerDemo done\n")
}