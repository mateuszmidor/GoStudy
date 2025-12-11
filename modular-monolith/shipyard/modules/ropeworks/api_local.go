package ropeworks

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/ropeworks/internal"
)

// APILocal implements the ropeworks module API.
type APILocal struct {
	r *internal.Ropeworks
}

func NewLocalAPI() *APILocal {
	log.Println("NewRopeworksLocal client")
	return &APILocal{r: internal.NewRopeworks()}
}

func (rl *APILocal) GetRopes(count int) ([]Rope, error) {
	return make([]Rope, len(rl.r.GetRopes(count))), nil
}

func (rl *APILocal) Run() {
	rl.r.Run()
}
