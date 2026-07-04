package eventbus

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedkernel"
)

type Subscriber func(evt sharedkernel.Event)

// Bus interface for dependency injection
type Bus interface {
	Publish(evt sharedkernel.Event)
	AddSubscriber(sub Subscriber)
	Run()
}

// eventBus implements Bus
type eventBus struct {
	events      chan sharedkernel.Event
	subscribers []Subscriber
}

func New(capacity int) Bus {
	events := make(chan sharedkernel.Event, capacity)
	return &eventBus{
		events:      events,
		subscribers: make([]Subscriber, 0),
	}
}

func (bus *eventBus) AddSubscriber(sub Subscriber) {
	bus.subscribers = append(bus.subscribers, sub)
}

func (bus *eventBus) Publish(evt sharedkernel.Event) {
	select {
	case bus.events <- evt:
	default:
		log.Println("WARNING: EventBus queue full; discarding new event")
	}
}

func (bus *eventBus) Run() {
	go func() {
		for evt := range bus.events {
			for _, sub := range bus.subscribers {
				sub(evt)
			}
		}
	}()
}
