package aggressiveness

import (
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/events"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/sounds"
)

// Level represents aggressiveness level
type Level func(eventsList events.Events) (soundList sounds.Sounds)
