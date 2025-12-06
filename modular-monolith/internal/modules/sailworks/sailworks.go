package sailworks

import (
	"log"
	"time"
)

type Sail struct{}
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
			for range numSailsPerSecond {
				s.sails <- &Sail{}
			}
			time.Sleep(time.Second)
		}
	}()
}

func (s *Sailworks) GetSails(count int) []Sail {
	result := make([]Sail, 0, count)
	for range count {
		sail := <-s.sails
		result = append(result, *sail)
		log.Println("received 1 sail")
	}
	return result
}
