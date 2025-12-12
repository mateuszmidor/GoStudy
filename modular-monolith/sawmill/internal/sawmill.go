package internal

import (
	"log"
	"time"
)

// Plank as Domain Object.
type Plank struct{}

// Sawmill as Domain Service.
type Sawmill struct {
	planks chan *Plank
}

const numBeamsPerSecond = 2
const numPlanksPerSecond = 5

func NewSawmill() *Sawmill {
	return &Sawmill{
		planks: make(chan *Plank, 100),
	}
}

func (s *Sawmill) Run() {
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

func (s *Sawmill) GetPlanks(count int) []Plank {
	result := make([]Plank, 0, count)
	for i := 0; i < count; i++ {
		p := <-s.planks
		result = append(result, *p)
		log.Println("received 1 plank")
	}
	return result
}

