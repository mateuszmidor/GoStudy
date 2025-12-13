package sawmill

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/internal"
)

// APIImpl implements the sawmill module API.
type APIImpl struct {
	s *internal.Sawmill
}

func NewAPI() *APIImpl {
	log.Println("NewAPI for sawmill")
	return &APIImpl{s: internal.NewSawmill()}
}

func (sl *APIImpl) Run() {
	sl.s.Run()
}

func (sl *APIImpl) GetBeams(count int) ([]Beam, error) {
	return make([]Beam, len(sl.s.GetBeams(count))), nil
}
