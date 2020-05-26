package main

import (
	gearboxdriver "driver"
	"driver/externalsystemsfacade"
	"fmt"
	"shared/gas"
	"soundmodule"
	"time"
)

type scenario = func() (gas.Value, gearboxdriver.DDrive, soundmodule.SoundModule)

func params(s externalsystemsfacade.Stub) {
	fmt.Printf("%4v RPM", s.CurrentRPM)
	if s.TrailorAttached {
		fmt.Print(", trailor")
	}
	if s.DrivingDownTheSlope {
		fmt.Print(", driving down")
	}
	fmt.Print(" - ")
}

func justDriving() (gas.Value, gearboxdriver.DDrive, soundmodule.SoundModule) {
	esf := externalsystemsfacade.Stub{
		CurrentRPM: 1500,
	}
	d := gearboxdriver.NewDDrive(&esf)
	sm := soundmodule.NewSoundModule()

	params(esf)
	println("just driving...")
	return gas.Half, d, sm
}

func speedingUp() (gas.Value, gearboxdriver.DDrive, soundmodule.SoundModule) {
	esf := externalsystemsfacade.Stub{
		CurrentRPM: 2500,
	}
	d := gearboxdriver.NewDDrive(&esf)
	sm := soundmodule.NewSoundModule()

	params(esf)
	println("speeding up...")
	return gas.Full, d, sm
}

func slowingDown() (gas.Value, gearboxdriver.DDrive, soundmodule.SoundModule) {
	esf := externalsystemsfacade.Stub{
		CurrentRPM: 500,
	}
	d := gearboxdriver.NewDDrive(&esf)
	sm := soundmodule.NewSoundModule()

	params(esf)
	println("slowing down..")
	return gas.Zero, d, sm
}

func breakingWithEngine() (gas.Value, gearboxdriver.DDrive, soundmodule.SoundModule) {
	esf := externalsystemsfacade.Stub{
		CurrentRPM:          1500,
		TrailorAttached:     true,
		DrivingDownTheSlope: true,
	}
	d := gearboxdriver.NewDDrive(&esf)
	sm := soundmodule.NewSoundModule()

	params(esf)
	println("breaking with engine..")
	return gas.Zero, d, sm
}

func kickDown1SportAggressiveness1() (gas.Value, gearboxdriver.DDrive, soundmodule.SoundModule) {
	esf := externalsystemsfacade.Stub{
		CurrentRPM: 1800,
	}

	d := gearboxdriver.NewDDrive(&esf)
	d.SetDrivingModeSport()
	sm := soundmodule.NewSoundModule()

	params(esf)
	println("kick down 1/sport/aggressiveness1...")
	return gas.New(0.5), d, sm
}

func kickDown2SportAggressiveness3() (gas.Value, gearboxdriver.DDrive, soundmodule.SoundModule) {
	esf := externalsystemsfacade.Stub{
		CurrentRPM: 2800,
	}

	d := gearboxdriver.NewDDrive(&esf)
	d.SetDrivingModeSport()
	d.SetAggressivenessLevel3()
	sm := soundmodule.NewSoundModule()
	sm.SetAggressivenessLevel3()

	params(esf)
	println("kick down 2/sport/aggressiveness3...")
	return gas.New(0.7), d, sm
}

func runPacer(pace time.Duration, scenarios []scenario) {
	for i := 0; i < 60; i++ {
		s := scenarios[i%len(scenarios)]
		gas, ddrive, soundModule := s()
		change, events := ddrive.HandleGas(gas)
		sounds := soundModule.HandleEvents(events)
		fmt.Printf("%v %v\n", change, sounds)
		println()
		time.Sleep(pace)
	}
}
func main() {
	sundayRide := []scenario{
		justDriving,
		slowingDown,
		breakingWithEngine,
		justDriving,
		speedingUp,
		kickDown1SportAggressiveness1,
		kickDown2SportAggressiveness3,
	}
	runPacer(3*time.Second, sundayRide)
}
