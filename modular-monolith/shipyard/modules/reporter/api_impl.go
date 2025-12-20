package reporter

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"
)

// APIImpl implements the reporter module API.
type APIImpl struct{}

func NewAPI() *APIImpl {
	return &APIImpl{}
}

// HandleMessage messages coming from MessageBus
func (r *APIImpl) HandleMessage(msg messagebus.Message) {
	switch v := msg.(type) {
	case *messagebus.ProductCreated:
		log.Printf("%s created (%d)", v.Name, v.Quantity)
	default:
		// Ignore other message types
	}
}
