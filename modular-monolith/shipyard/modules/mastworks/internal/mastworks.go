package internal

import (
	"log"
	"time"

	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"
)

// Mast as Domain Object.
type Mast struct{}

// Mastworks as Domain Service.
type Mastworks struct {
	masts      chan *Mast
	sawmillAPI sawmill.API
	messageBus messagebus.Bus
}

const numMastsPerSecond = 1
const beamsPerMast = 3

func NewMastworks(sawmillAPI sawmill.API, bus messagebus.Bus) *Mastworks {
	return &Mastworks{
		masts:      make(chan *Mast, 100),
		sawmillAPI: sawmillAPI,
		messageBus: bus,
	}
}

func (m *Mastworks) Run() {
	go func() {
		for {
			for range numMastsPerSecond {
				m.masts <- &Mast{}
				m.messageBus.Publish(&messagebus.ProductCreated{
					Name:     "mast",
					Quantity: 1,
				})
			}
			time.Sleep(time.Second)
		}
	}()
}

func (m *Mastworks) GetMasts(count int) []Mast {
	result := make([]Mast, 0, count)
	for i := 0; i < count; i++ {
		// Request beams from sawmill for each mast
		beams, err := m.sawmillAPI.GetBeams(beamsPerMast)
		if err != nil {
			log.Printf("Mastworks failed to get beams for mast: %v", err)
			continue
		}
		log.Printf("Mastworks received %d beams for making a mast", len(beams))

		mast := <-m.masts
		result = append(result, *mast)
	}
	return result
}
