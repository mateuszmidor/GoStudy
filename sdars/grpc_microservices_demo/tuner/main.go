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

	// tuner side of the communication;
	// sends subscription/station list updates(ui), tune requests(hw)
	// receives subscription/station list updates(hw), tune requests(ui)
	tuner := tuner.NewTunerRoot()

	// expose Tuner grpc interface
	adapterTuner := NewTunerAdapter(&tuner)
	tuner.SetHwPort(&adapterTuner)
	tuner.SetUiPort(&adapterTuner)

	// run tuner
	go tuner.Run()
	go adapterTuner.RunGrpcServer()

	// wait for INT/TERM
	utils.NewShutdownCondition().Wait()

	// demo done
	fmt.Printf("Tuner done\n")
}
