package clients

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/sailworks"
)

// SailworksLocal implements the Sailworks interface and wraps a sailworks.Sailworks instance
type SailworksLocal struct {
	s *sailworks.Sailworks
}

func NewSailworksLocal() *SailworksLocal {
	log.Println("NewSailworksLocal client")
	return &SailworksLocal{s: sailworks.NewSailworks()}
}

func (sl *SailworksLocal) GetSails(count int) ([]Sail, error) {
	return make([]Sail, len(sl.s.GetSails(count))), nil
}

func (sl *SailworksLocal) Run() {
	sl.s.Run()
}
