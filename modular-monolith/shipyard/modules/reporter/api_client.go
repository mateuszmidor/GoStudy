package reporter

import (
	"log"
	"sync"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"
)

type APIClient struct {
	mu              sync.RWMutex
	productCounters map[string]uint
}

func NewAPI() *APIClient {
	return &APIClient{
		productCounters: make(map[string]uint),
	}
}

func (r *APIClient) HandleMessage(msg messagebus.Message) {
	switch v := msg.(type) {
	case *messagebus.ProductCreated:
		r.mu.Lock()
		r.productCounters[v.Name] += v.Quantity
		r.mu.Unlock()
	default:
	}
}

func (r *APIClient) PrintReport() {
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

var _ API = &APIClient{}
