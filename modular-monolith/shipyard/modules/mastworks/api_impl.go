package mastworks

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/mastworks/internal"
)

// APIImpl implements the mastworks module API.
type APIImpl struct {
	m *internal.Mastworks
}

func NewAPI(sawmillAPI sawmill.API) *APIImpl {
	log.Println("NewMastworksLocal client")
	return &APIImpl{m: internal.NewMastworks(sawmillAPI)}
}

func (ml *APIImpl) GetMasts(count int) ([]Mast, error) {
	return make([]Mast, len(ml.m.GetMasts(count))), nil
}

func (ml *APIImpl) Run() {
	ml.m.Run()
}
