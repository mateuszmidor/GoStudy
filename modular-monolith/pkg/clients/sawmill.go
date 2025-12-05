package clients

import "github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/sawmill"

type Sawmill interface {
	Run()
	GetBeams(count int) []sawmill.Beam
	GetPlanks(count int) []sawmill.Plank
}
