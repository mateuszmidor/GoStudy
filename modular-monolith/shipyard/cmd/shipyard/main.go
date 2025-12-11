package main

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/configs"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/ropeworks"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/sailworks"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/sawmill"
	"golang.org/x/sync/errgroup"
)

func main() {
	sawmillAPI := sawmill.NewLocalAPI()
	sawmillAPI.Run()
	ropeworksAPI := ropeworks.NewLocalAPI()
	ropeworksAPI.Run()
	sailworksAPI := sailworks.NewSailworksGRPC(configs.SailworksAddr)
	sailworksAPI.Run()
	buildShip(sawmillAPI, ropeworksAPI, sailworksAPI)
}

func buildShip(_sawmill sawmill.API, _ropeworks ropeworks.API, _sailworks sailworks.API) {
	planks := []sawmill.Plank{}
	ropes := []ropeworks.Rope{}
	sails := []sailworks.Sail{}

	g := errgroup.Group{}
	g.Go(func() error {
		var err error
		planks, err = _sawmill.GetPlanks(15)
		return err
	})
	g.Go(func() error {
		var err error
		ropes, err = _ropeworks.GetRopes(9)
		return err
	})
	g.Go(func() error {
		var err error
		sails, err = _sailworks.GetSails(2)
		return err
	})
	if err := g.Wait(); err != nil {
		log.Fatalf("buildShip failed: %+v", err)
	}
	log.Println("collected", len(planks), "planks,", len(ropes), "ropes,", len(sails), "sails")
	log.Println("### ship built successfuly ###")
}
