package internal

import (
	"log"
	"time"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"
)

// Rope as Domain Object.
type Rope struct{}

// Ropeworks as Domain Service.
type Ropeworks struct {
	ropes      chan *Rope
	messageBus messagebus.Bus
}

const numRopesPerSecond = 3

func NewRopeworks(bus messagebus.Bus) *Ropeworks {
	return &Ropeworks{
		ropes:      make(chan *Rope, 100),
		messageBus: bus,
	}
}

func (r *Ropeworks) Run() {
	go func() {
		for {
			for range numRopesPerSecond {
				// produce
				r.ropes <- &Rope{}

				// notify
				r.messageBus.Publish(&messagebus.ProductCreated{Name: "rope", Quantity: 1})

				// log
				log.Println("Ropeworks produced 1 rope")
			}
			time.Sleep(time.Second)
		}
	}()
}

func (r *Ropeworks) GetRopes(count int) []Rope {
	result := make([]Rope, 0, count)
	for i := 0; i < count; i++ {
		rope := <-r.ropes
		result = append(result, *rope)
	}
	return result
}
