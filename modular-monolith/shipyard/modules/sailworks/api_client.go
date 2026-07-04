package sailworks

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/sailworks/internal"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/eventbus"
)

type APIClient struct {
	s *internal.Sailworks
}

func NewAPI(bus eventbus.Bus) *APIClient {
	log.Println("NewSailworksLocal client")
	return &APIClient{s: internal.NewSailworks(bus)}
}

func (sl *APIClient) GetSails(count int) ([]Sail, error) {
	return make([]Sail, len(sl.s.GetSails(count))), nil
}

func (sl *APIClient) Run() {
	sl.s.Run()
}
