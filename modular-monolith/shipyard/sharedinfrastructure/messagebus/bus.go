package messagebus

type Subscriber func(msg Message)

type messageBus struct {
	messages    chan Message
	subscribers []Subscriber
}

func new(capacity int) *messageBus {
	messages := make(chan Message, capacity)
	return &messageBus{
		messages:    messages,
		subscribers: make([]Subscriber, 0),
	}
}

func (bus *messageBus) Subscribe(sub Subscriber) {
	bus.subscribers = append(bus.subscribers, sub)
}

func (bus *messageBus) Publish(msg Message) {
	bus.messages <- msg
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

var MessageBus *messageBus

func init() {
	MessageBus = new(100)
	MessageBus.Run()
}
