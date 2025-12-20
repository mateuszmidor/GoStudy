package internal

import (
	"time"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"
)

// Sail as Domain Object.
type Sail struct{}

// Sailworks as Domain Service.
type Sailworks struct {
	sails      chan *Sail
	messageBus messagebus.Bus
}

const numSailsPerSecond = 2

func NewSailworks(bus messagebus.Bus) *Sailworks {
	return &Sailworks{
		sails:      make(chan *Sail, 100),
		messageBus: bus,
	}
}

func (s *Sailworks) Run() {
	go func() {
		for {
			for i := 0; i < numSailsPerSecond; i++ {
				s.sails <- &Sail{}
				s.messageBus.Publish(&messagebus.ProductCreated{
					Name:     "sail",
					Quantity: 1,
				})
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
	}
	return result
}
