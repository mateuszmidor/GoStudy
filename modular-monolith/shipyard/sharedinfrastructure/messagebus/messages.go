package messagebus

import "time"

type Message interface {
}

type LunchBreakStarted struct {
	Duration time.Duration
}
