package reporter

import "github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedinfrastructure/messagebus"

// API of the reporter module.
type API interface {
	Handle(msg messagebus.Message)
}
