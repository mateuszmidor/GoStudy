package sailworks

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/sailworks/internal"
)

// APILocal implements the sailworks module API.
type APILocal struct {
	s *internal.Sailworks
}

func NewSailworksLocal() *APILocal {
	log.Println("NewSailworksLocal client")
	return &APILocal{s: internal.NewSailworks()}
}

func (sl *APILocal) GetSails(count int) ([]Sail, error) {
	return make([]Sail, len(sl.s.GetSails(count))), nil
}

func (sl *APILocal) Run() {
	sl.s.Run()
}
