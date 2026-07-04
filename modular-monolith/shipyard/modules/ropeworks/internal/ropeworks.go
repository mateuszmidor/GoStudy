package internal

import (
	"log"
	"time"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/eventbus"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedkernel"
)

// Rope as Domain Object.
type Rope struct{}

// Ropeworks as Domain Service.
type Ropeworks struct {
	ropes      chan *Rope
	eventBus eventbus.Bus
}

const numRopesPerSecond = 3

func NewRopeworks(bus eventbus.Bus) *Ropeworks {
	return &Ropeworks{
		ropes:      make(chan *Rope, 100),
		eventBus: bus,
	}
}

func (r *Ropeworks) Run() {
	go func() {
		for {
			for range numRopesPerSecond {
				// produce
				r.ropes <- &Rope{}

				// notify
				r.eventBus.Publish(&sharedkernel.ProductCreated{Name: "rope", Quantity: 1})

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
