package messagebus

import "log"

type Subscriber func(msg Message)

// Bus interface for dependency injection
type Bus interface {
	Publish(msg Message)
	AddSubscriber(sub Subscriber)
	Run()
}

// messageBus implements Bus
type messageBus struct {
	messages    chan Message
	subscribers []Subscriber
}

func New(capacity int) Bus {
	messages := make(chan Message, capacity)
	return &messageBus{
		messages:    messages,
		subscribers: make([]Subscriber, 0),
	}
}

func (bus *messageBus) AddSubscriber(sub Subscriber) {
	bus.subscribers = append(bus.subscribers, sub)
}

func (bus *messageBus) Publish(msg Message) {
	select {
	case bus.messages <- msg:
	default:
		log.Println("WARNING: MessageBus queue full; discarding new message")
	}
}

func (bus *messageBus) Run() {
	go func() {
		for msg := range bus.messages {
			for _, sub := range bus.subscribers {
				sub(msg)
			}
		}
	}()
}
