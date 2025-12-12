package mastworks

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/mastworks/internal"
)

// APILocal implements the mastworks module API.
type APILocal struct {
	m *internal.Mastworks
}

func NewLocalAPI(sawmillAPI sawmill.API) *APILocal {
	log.Println("NewMastworksLocal client")
	return &APILocal{m: internal.NewMastworks(sawmillAPI)}
}

func (ml *APILocal) GetMasts(count int) []Mast {
	return make([]Mast, len(ml.m.GetMasts(count)))
}

func (ml *APILocal) Run() {
	ml.m.Run()
}
