package main

import (
	"fmt"
	"hexagons/tuner"
	"math/rand"
	"time"
	"utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// tuner; the middle side. Does business logic and communicates with hw and ui
	tuner := tuner.NewTunerRoot()

	// expose Tuner http interface
	adapterTuner := NewTunerAdapter(&tuner)
	tuner.SetHwPort(&adapterTuner)
	tuner.SetUiPort(&adapterTuner)

	// run tuner
	go tuner.Run()
	go adapterTuner.RunHTTPServer()

	// wait for INT/TERM
	utils.NewShutdownCondition().Wait()

	// demo done
	fmt.Printf("Tuner done\n")
}
