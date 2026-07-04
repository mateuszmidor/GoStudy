package internal

import (
	"log"
	"time"

	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/eventbus"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedkernel"
)

// Mast as Domain Object.
type Mast struct{}

// Mastworks as Domain Service.
type Mastworks struct {
	masts      chan *Mast
	sawmillAPI sawmill.API
	eventBus eventbus.Bus
}

const numMastsPerSecond = 1
const beamsPerMast = 3

func NewMastworks(sawmillAPI sawmill.API, bus eventbus.Bus) *Mastworks {
	return &Mastworks{
		masts:      make(chan *Mast, 100),
		sawmillAPI: sawmillAPI,
		eventBus: bus,
	}
}

func (m *Mastworks) Run() {
	go func() {
		for {
			for range numMastsPerSecond {
				// produce
				beams, err := m.sawmillAPI.GetBeams(beamsPerMast) // beams are needed to produce masts
				if err != nil {
					log.Printf("Mastworks failed to get beams for mast: %v", err)
					continue
				}
				log.Printf("Mastworks received %d beams for making a mast", len(beams))
				m.masts <- &Mast{}

				// notify
				m.eventBus.Publish(&sharedkernel.ProductCreated{Name: "mast", Quantity: 1})

				// log
				log.Println("Mastworks produced 1 mast")
			}
			time.Sleep(time.Second)
		}
	}()
}

func (m *Mastworks) GetMasts(count int) []Mast {
	result := make([]Mast, 0, count)
	for i := 0; i < count; i++ {
		mast := <-m.masts
		result = append(result, *mast)
	}
	return result
}
