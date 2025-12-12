package internal

import (
	"log"
	"time"

	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill"
)

// Mast as Domain Object.
type Mast struct{}

// Mastworks as Domain Service.
type Mastworks struct {
	masts      chan *Mast
	sawmillAPI sawmill.API
}

const numMastsPerSecond = 1
const planksPerMast = 3

func NewMastworks(sawmillAPI sawmill.API) *Mastworks {
	return &Mastworks{
		masts:      make(chan *Mast, 100),
		sawmillAPI: sawmillAPI,
	}
}

func (m *Mastworks) Run() {
	go func() {
		for {
			for i := 0; i < numMastsPerSecond; i++ {
				m.masts <- &Mast{}
			}
			time.Sleep(time.Second)
		}
	}()
}

func (m *Mastworks) GetMasts(count int) []Mast {
	result := make([]Mast, 0, count)
	for i := 0; i < count; i++ {
		// Request planks from sawmill for each mast
		planks, err := m.sawmillAPI.GetPlanks(planksPerMast)
		if err != nil {
			log.Printf("failed to get planks for mast: %v", err)
			continue
		}
		log.Printf("received %d planks for mast", len(planks))

		mast := <-m.masts
		result = append(result, *mast)
		log.Println("received 1 mast")
	}
	return result
}
