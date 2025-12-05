package clients

import "github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/sailworks"

type Sailworks interface {
	GetSails(count int) ([]sailworks.Sail, error)
}
