package mastworks

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/mastworks/internal"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"
)

type APIClient struct {
	m *internal.Mastworks
}

func NewAPI(sawmillAPI sawmill.API, bus messagebus.Bus) *APIClient {
	log.Println("NewMastworksLocal client")
	return &APIClient{m: internal.NewMastworks(sawmillAPI, bus)}
}

func (ml *APIClient) GetMasts(count int) ([]Mast, error) {
	return make([]Mast, len(ml.m.GetMasts(count))), nil
}

func (ml *APIClient) Run() {
	ml.m.Run()
}
