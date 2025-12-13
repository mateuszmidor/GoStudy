package ropeworks

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/ropeworks/internal"
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
