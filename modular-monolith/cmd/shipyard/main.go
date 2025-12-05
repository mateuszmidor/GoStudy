package main

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/configs"
	"github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/ropeworks"
	"github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/sailworks"
	"github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/sawmill"
	"github.com/mateuszmidor/GoStudy/modular-monolith/pkg/clients"
	"golang.org/x/sync/errgroup"
)

func main() {
	sawmill := clients.NewSawmillLocal()
	sawmill.Run()
	ropeworks := clients.NewRopeworksLocal()
	ropeworks.Run()
	// we have both, sailworks as built-in module and as a grpc service
	// sailworks := clients.NewSailworksLocal()
	sailworks := clients.NewSailworksGrpc(configs.SailworksAddr)
	sailworks.Run()
	buildShip(sawmill, ropeworks, sailworks)
}

func buildShip(_sawmill clients.Sawmill, _ropeworks clients.Ropeworks, _sailworks clients.Sailworks) {
	planks := []sawmill.Plank{}
	ropes := []ropeworks.Rope{}
	sails := []sailworks.Sail{}

	g := errgroup.Group{}
	g.Go(func() error {
		planks = _sawmill.GetPlanks(15)
		return nil
	})
	g.Go(func() error {
		ropes = _ropeworks.GetRopes(9)
		return nil
	})
	g.Go(func() error {
		var err error
		sails, err = _sailworks.GetSails(2)
		return err
	})
	if err := g.Wait(); err != nil {
		log.Fatal("buildShip failed:", err)
	}
	log.Println("collected", len(planks), "planks,", len(ropes), "ropes,", len(sails), "sails")
	log.Println("ship built successfuly")
}
