package clients

import "github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/ropeworks"

type Ropeworks interface {
	GetRopes(count int) []ropeworks.Rope
}
