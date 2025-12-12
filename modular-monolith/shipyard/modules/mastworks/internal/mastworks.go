package internal

import (
	"log"
	"time"
)

// Mast as Domain Object.
type Mast struct{}

// Mastworks as Domain Service.
type Mastworks struct {
	masts chan *Mast
}

const numMastsPerSecond = 1

func NewMastworks() *Mastworks {
	return &Mastworks{
		masts: make(chan *Mast, 100),
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
		mast := <-m.masts
		result = append(result, *mast)
		log.Println("received 1 mast")
	}
	return result
}
