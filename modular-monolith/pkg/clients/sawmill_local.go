package clients

import "github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/sawmill"

// SawmillLocal implements the Sawmill interface and wraps a sawmill.Sawmill instance
type SawmillLocal struct {
	s *sawmill.Sawmill
}

func NewSawmillLocal() *SawmillLocal {
	return &SawmillLocal{s: sawmill.NewSawmill()}
}

func (sl *SawmillLocal) Run() {
	sl.s.Run()
}

func (sl *SawmillLocal) GetBeams(count int) []sawmill.Beam {
	return sl.s.GetBeams(count)
}

func (sl *SawmillLocal) GetPlanks(count int) []sawmill.Plank {
	return sl.s.GetPlanks(count)
}
