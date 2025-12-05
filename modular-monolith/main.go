package main

import (
	"github.com/mateuszmidor/GoStudy/modular-monolith/pkg/clients"
)

func main() {
	sawmillIndustry := clients.NewSawmillLocal()
	sawmillIndustry.Run()
	ropeworksIndustry := clients.NewRopeworksLocal()
	ropeworksIndustry.Run()
	buildShip(sawmillIndustry, ropeworksIndustry)
}
