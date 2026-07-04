package ropeworks

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/ropeworks/internal"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/eventbus"
)

type APIClient struct {
	r *internal.Ropeworks
}

func NewAPI(bus eventbus.Bus) *APIClient {
	log.Println("NewRopeworksLocal client")
	return &APIClient{r: internal.NewRopeworks(bus)}
}

func (rl *APIClient) GetRopes(count int) ([]Rope, error) {
	return make([]Rope, len(rl.r.GetRopes(count))), nil
}

func (rl *APIClient) Run() {
	rl.r.Run()
}
