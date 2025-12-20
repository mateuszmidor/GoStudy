package sailworks

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/sailworks/internal"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"
)

// APIImpl implements the sailworks module API.
type APIImpl struct {
	s *internal.Sailworks
}

func NewAPI(bus messagebus.Bus) *APIImpl {
	log.Println("NewSailworksLocal client")
	return &APIImpl{s: internal.NewSailworks(bus)}
}

func (sl *APIImpl) GetSails(count int) ([]Sail, error) {
	return make([]Sail, len(sl.s.GetSails(count))), nil
}

func (sl *APIImpl) Run() {
	sl.s.Run()
}
