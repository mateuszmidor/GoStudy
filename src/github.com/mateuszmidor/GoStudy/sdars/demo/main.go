package main

import (
	"fmt"
	"time"
	"math/rand"
	"sdars"
	"adapters"
	"actors/hardware"
	"actors/cluster"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// hardware side of the communication; sends station list updates, receives station tune requests
	hardware:= hardware.NewHwActor()

	// cluster side of the communication; sends tune requests, receives station list updates
	cluster := cluster.NewClusterActor()

	// command processor
	tuner := sdars.NewTunerCommandProcessor()

	// adapts tuner hardware port for our specific HwActor
	hwAdapter := adapters.NewHardwareAdapter(&tuner.CommandQueue, &hardware)

	// adapts tuner cluster port for our specific ClusterActor
	clusterAdapter := adapters.NewClusterAdapter(&tuner.CommandQueue, &cluster)

	// tuner communicated with the outer world through its ports
	tuner.SetupPorts(hwAdapter, clusterAdapter)

	// run outer world actors in their own threads
	go hardware.Run()
	go cluster.Run()

	// this is shutdown condition for the tuner command processing
	// shutdownCondition := sdars.NewShutdownCondition()

	// start processing commands and block until shutdownCondition is fulfilled
	tuner.Run(sdars.NewShutdownCondition())

	// demo done
	fmt.Printf("TunerDemo done")
}