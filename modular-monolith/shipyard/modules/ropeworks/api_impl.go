package ropeworks

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/ropeworks/internal"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"
)

// APIImpl implements the ropeworks module API.
type APIImpl struct {
	r *internal.Ropeworks
}

func NewAPI() *APIImpl {
	log.Println("NewRopeworksLocal client")
	return &APIImpl{r: internal.NewRopeworks()}
}

func (rl *APIImpl) GetRopes(count int) ([]Rope, error) {
	return make([]Rope, len(rl.r.GetRopes(count))), nil
}

func (rl *APIImpl) Run() {
	rl.r.Run()
}

// Handle messages coming from other modules/main program
func (a *APIImpl) Handle(msg messagebus.Message) {
	switch v := msg.(type) {
	case *messagebus.LunchBreakStarted:
		log.Println("ropeworks handles LunchBreakStarted event:", v.Duration)
	default:
		log.Println("ropeworks handles unknown message")
	}
}
