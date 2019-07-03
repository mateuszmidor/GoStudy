package main

import (
	"fmt"
	"time"
	"math/rand"
	"actors/hardware"
	"adapters"
	"hexagons/ui"
	"hexagons/tuner"
) 

func main() {
	rand.Seed(time.Now().UnixNano())

	// hardware side of the communication; sends station list updates, receives station tune requests
	hardware := hardware.NewHwActor()

	// ui side of the communication; sends tune requests, receives station list updates
	ui := ui.NewUiRoot()

	// aggregate root
	tuner := tuner.NewTunerRoot()

	// hardware talks to tuner
	hwAdapter := adapters.NewHardwareAdapter(&tuner, &hardware)

	// communicate hexagons tuner <-> ui
	adapterTunerUi := adapters.NewUiAdapter(&tuner, &ui)
	ui.SetupTunerPortOut(&adapterTunerUi)
	tuner.SetupUiPortOut(&adapterTunerUi)

	// tuner talks to hardware and ui
	tuner.SetupHwPortOut(&hwAdapter)

	// ui.SetupPorts(uiAdapter)
	// hardware.SetupPorts(hwAdapter)
	// run all the parties

	go ui.Run()
	go tuner.Run()
	go hardware.Run()

	// wait for INT/TERM
	NewShutdownCondition().Wait()

	// demo done
	fmt.Printf("TunerDemo done\n")
}