package reporter

import "github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/sharedkernel"

// API of the reporter module.
type API interface {
	HandleEvent(msg sharedkernel.Event)
	PrintReport()
}
