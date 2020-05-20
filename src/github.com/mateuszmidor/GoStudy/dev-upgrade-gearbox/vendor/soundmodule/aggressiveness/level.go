package aggressiveness

import (
	"shared/events"
	"shared/sounds"
)

// Level represents aggressiveness level
type Level func(eventsList events.Events) (soundList sounds.Sounds)
