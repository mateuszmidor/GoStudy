package internal

import (
	"log"
	"time"
)

// Beam as Domain Object.
type Beam struct{}

// Sawmill as Domain Service.
type Sawmill struct {
	beams chan *Beam
}

const numBeamsPerSecond = 5

func NewSawmill() *Sawmill {
	return &Sawmill{
		beams: make(chan *Beam, 100),
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
}

func (s *Sawmill) GetBeams(count int) []Beam {
	result := make([]Beam, 0, count)
	for i := 0; i < count; i++ {
		b := <-s.beams
		result = append(result, *b)
		log.Println("Sawmill produced 1 beam")
	}
	return result
}
