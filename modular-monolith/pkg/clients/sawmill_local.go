package clients

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/sawmill"
)

// SawmillLocal implements the Sawmill interface and wraps a sawmill.Sawmill instance
type SawmillLocal struct {
	s *sawmill.Sawmill
}

func NewSawmillLocal() *SawmillLocal {
	log.Println("NewSawmillLocal client")
	return &SawmillLocal{s: sawmill.NewSawmill()}
}

func (sl *SawmillLocal) Run() {
	sl.s.Run()
}

func (sl *SawmillLocal) GetPlanks(count int) ([]Plank, error) {
	return make([]Plank, len(sl.s.GetPlanks(count))), nil
}
