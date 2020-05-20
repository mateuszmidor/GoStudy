package aggressiveness

import (
	"shared/events"
	"shared/sounds"
)

// GetSoundsForEventsLevel3 returns
func GetSoundsForEventsLevel3(eventsList events.Events) (soundList sounds.Sounds) {
	if eventsList.Contains(events.DoubleGearDown) {
		soundList = soundList.Append(sounds.PipeBlast)
	} else if eventsList.Contains(events.GearDown) {
		soundList = soundList.Append(sounds.PipeBlast)
	}
	return
}
