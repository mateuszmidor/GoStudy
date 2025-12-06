package clients

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/ropeworks"
)

// RopeworksLocal implements the Ropeworks interface and wraps a ropeworks.Ropeworks instance
type RopeworksLocal struct {
	r *ropeworks.Ropeworks
}

func NewRopeworksLocal() *RopeworksLocal {
	log.Println("NewRopeworksLocal client")
	return &RopeworksLocal{r: ropeworks.NewRopeworks()}
}

func (rl *RopeworksLocal) GetRopes(count int) ([]Rope, error) {
	return make([]Rope, len(rl.r.GetRopes(count))), nil
}

func (rl *RopeworksLocal) Run() {
	rl.r.Run()
}
