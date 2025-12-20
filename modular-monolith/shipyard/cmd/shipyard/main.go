package main

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/configs"
	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/mastworks"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/reporter"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/ropeworks"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/sailworks"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"
	"golang.org/x/sync/errgroup"
)

func main() {
	// initialize external service client
	sawmillAPI := sawmill.NewGrpcClient(configs.SawmillAddr)

	// initialize local modules
	ropeworksAPI := ropeworks.NewAPI(messagebus.Instance)
	ropeworksAPI.Run()
	mastworksAPI := mastworks.NewAPI(sawmillAPI, messagebus.Instance)
	mastworksAPI.Run()
	sailworksAPI := sailworks.NewAPI(messagebus.Instance)
	sailworksAPI.Run()
	reporterAPI := reporter.NewAPI()
	messagebus.Instance.AddSubscriber(reporterAPI.HandleMessage)

	// execute the use case
	buildShip(ropeworksAPI, mastworksAPI, sailworksAPI)

	// print production statistics
	reporterAPI.PrintReport()
}

func buildShip(_ropeworks ropeworks.API, _mastworks mastworks.API, _sailworks sailworks.API) {
	ropes := []ropeworks.Rope{}
	masts := []mastworks.Mast{}
	sails := []sailworks.Sail{}

	// request resources
	g := errgroup.Group{}
	g.Go(func() error {
		var err error
		ropes, err = _ropeworks.GetRopes(9)
		return err
	})
	g.Go(func() error {
		var err error
		masts, err = _mastworks.GetMasts(2)
		return err
	})
	g.Go(func() error {
		var err error
		sails, err = _sailworks.GetSails(4)
		return err
	})

	// wait till are resources are collected
	if err := g.Wait(); err != nil {
		log.Fatalf("buildShip failed: %+v", err)
	}

	// success
	log.Printf("ship built successfully (%d ropes, %d masts, %d sails)", len(ropes), len(masts), len(sails))
}
