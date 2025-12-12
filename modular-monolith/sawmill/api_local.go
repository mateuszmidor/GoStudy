package sawmill

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/internal"
)

// APILocal implements the sawmill module API.
type APILocal struct {
	s *internal.Sawmill
}

func NewSawmillLocal() *APILocal {
	log.Println("NewSawmillLocal client")
	return &APILocal{s: internal.NewSawmill()}
}

func (sl *APILocal) Run() {
	sl.s.Run()
}

func (sl *APILocal) GetPlanks(count int) ([]Plank, error) {
	return make([]Plank, len(sl.s.GetPlanks(count))), nil
}

