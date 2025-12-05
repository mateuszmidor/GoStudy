package ropeworks

import (
	"log"
	"time"
)

type Rope struct{}
type Ropeworks struct {
	ropes chan *Rope
}

const numRopesPerSecond = 3

func NewRopeworks() *Ropeworks {
	return &Ropeworks{
		ropes: make(chan *Rope, 100),
	}
}

func (r *Ropeworks) Run() {
	go func() {
		for {
			for i := 0; i < numRopesPerSecond; i++ {
				r.ropes <- &Rope{}
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
		log.Println("received 1 rope")
	}
	return result
}
