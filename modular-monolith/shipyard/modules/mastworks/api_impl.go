package mastworks

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/mastworks/internal"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"
)

// APIImpl implements the mastworks module API.
type APIImpl struct {
	m *internal.Mastworks
}

func NewAPI(sawmillAPI sawmill.API, bus messagebus.Bus) *APIImpl {
	log.Println("NewMastworksLocal client")
	return &APIImpl{m: internal.NewMastworks(sawmillAPI, bus)}
}

func (ml *APIImpl) GetMasts(count int) ([]Mast, error) {
	return make([]Mast, len(ml.m.GetMasts(count))), nil
}

func (ml *APIImpl) Run() {
	ml.m.Run()
}
