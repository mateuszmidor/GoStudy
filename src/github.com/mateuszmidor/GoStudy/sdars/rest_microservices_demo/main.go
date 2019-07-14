package main

import (
	"fmt"
	"hexagons/hw"
	"hexagons/tuner"
	"hexagons/ui"
	"math/rand"
	"time"
	"utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// hw side of the communication; sends subscription/station list updates, receives tune requests
	hw := hw.NewHwRoot()

	// ui side of the communication; sends tune requests, receives subscription/station list updates
	ui := ui.NewUiRoot()

	// tuner; the middle side. Does business logic and communicates with hw and ui
	tuner := tuner.NewTunerRoot()

	// communicate hexagons hw <-> tuner <-> ui over http rest
	adapterHw := NewHwAdapter(&hw)
	hw.SetTunerPort(&adapterHw)

	adapterTuner := NewTunerAdapter(&tuner)
	tuner.SetHwPort(&adapterTuner)
	tuner.SetUiPort(&adapterTuner)

	adapterUI := NewUIAdapter(&ui)
	ui.SetTunerPort(&adapterUI)

	// run all the parties
	go hw.Run()
	go tuner.Run()
	go ui.Run()
	go adapterHw.RunHTTPServer()
	go adapterTuner.RunHTTPServer()
	go adapterUI.RunHTTPServer()

	// wait for INT/TERM
	utils.NewShutdownCondition().Wait()

	// demo done
	fmt.Printf("TunerDemo done\n")
}
