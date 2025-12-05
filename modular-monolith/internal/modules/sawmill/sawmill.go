package sawmill

import (
	"log"
	"time"
)

type Beam struct{}
type Plank struct{}
type Sawmill struct {
	beams  chan *Beam
	planks chan *Plank
}

const numBeamsPerSecond = 2
const numPlanksPerSecond = 5

func NewSawmill() *Sawmill {
	return &Sawmill{
		beams:  make(chan *Beam, 100),
		planks: make(chan *Plank, 100),
	}
}

func (s *Sawmill) Run() {
	// Beam producer
	go func() {
		for {
			for i := 0; i < numBeamsPerSecond; i++ {
				s.beams <- &Beam{}
			}
			time.Sleep(time.Second)
		}
	}()

	// Plank producer
	go func() {
		for {
			for i := 0; i < numPlanksPerSecond; i++ {
				s.planks <- &Plank{}
			}
			time.Sleep(time.Second)
		}
	}()
}

func (s *Sawmill) GetBeams(count int) []Beam {
	result := make([]Beam, 0, count)
	for i := 0; i < count; i++ {
		b := <-s.beams
		result = append(result, *b)
		log.Println("received 1 beam")
	}
	return result

}

func (s *Sawmill) GetPlanks(count int) []Plank {
	result := make([]Plank, 0, count)
	for i := 0; i < count; i++ {
		p := <-s.planks
		result = append(result, *p)
		log.Println("received 1 plank")
	}
	return result
}
