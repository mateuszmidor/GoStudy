package reporter

import (
	"log"
	"sync"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"
)

// APIImpl implements the reporter module API.
type APIImpl struct {
	mu              sync.RWMutex
	productCounters map[string]uint
}

func NewAPI() *APIImpl {
	return &APIImpl{
		productCounters: make(map[string]uint),
	}
}

// HandleMessage func is a subscriber of the global MessageBus
func (r *APIImpl) HandleMessage(msg messagebus.Message) {
	switch v := msg.(type) {
	case *messagebus.ProductCreated:
		r.mu.Lock()
		r.productCounters[v.Name] += v.Quantity
		r.mu.Unlock()
	default:
		// Ignore other message types
	}
}

// PrintReport logs the total quantity of each product that has been created.
func (r *APIImpl) PrintReport() {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.productCounters) == 0 {
		log.Println("No products created yet")
		return
	}

	log.Println("=== Production Report ===")
	for productName, totalQuantity := range r.productCounters {
		log.Printf("total %ss: %d", productName, totalQuantity)
	}
	log.Println("=====================")
}
