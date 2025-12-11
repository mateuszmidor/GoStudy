package internal

import (
	"log"
	"time"
)

// Sail as Domain Object.
type Sail struct{}

// Sailworks as Domain Service.
type Sailworks struct {
	sails chan *Sail
}

const numSailsPerSecond = 1

func NewSailworks() *Sailworks {
	return &Sailworks{
		sails: make(chan *Sail, 100),
	}
}

func (s *Sailworks) Run() {
	go func() {
		for {
			for i := 0; i < numSailsPerSecond; i++ {
				s.sails <- &Sail{}
			}
			time.Sleep(time.Second)
		}
	}()
}

func (s *Sailworks) GetSails(count int) []Sail {
	result := make([]Sail, 0, count)
	for i := 0; i < count; i++ {
		sail := <-s.sails
		result = append(result, *sail)
		log.Println("received 1 sail")
	}
	return result
}
