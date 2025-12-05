package main

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/ropeworks"
	"github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/sawmill"
	"github.com/mateuszmidor/GoStudy/modular-monolith/pkg/clients"
	"golang.org/x/sync/errgroup"
)

func buildShip(_sawmill clients.Sawmill, _ropeworks clients.Ropeworks) {
	planks := []sawmill.Plank{}
	beams := []sawmill.Beam{}
	ropes := []ropeworks.Rope{}

	g := errgroup.Group{}
	g.Go(func() error {
		planks = _sawmill.GetPlanks(15)
		return nil
	})
	g.Go(func() error {
		beams = _sawmill.GetBeams(3)
		return nil
	})
	g.Go(func() error {
		ropes = _ropeworks.GetRopes(9)
		return nil
	})
	g.Wait()
	log.Println("collected", len(planks), "planks,", len(beams), "beams,", len(ropes), "ropes")
	log.Println("ship built successfuly")
}
