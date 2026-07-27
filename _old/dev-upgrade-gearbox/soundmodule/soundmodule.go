package soundmodule

import (
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/events"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/sounds"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/soundmodule/aggressiveness"
)

// SoundModule turns gear box events into sounds
type SoundModule struct {
	getSoundsForAggressivenessLevelEvents aggressiveness.Level
}

// NewSoundModule is constructor
func NewSoundModule() (sm SoundModule) {
	sm.SetAggressivenessLevel1()
	return
}

// HandleEvents turns events into proper sounds
func (sm *SoundModule) HandleEvents(list events.Events) (result sounds.Sounds) {
	result = append(result, sm.getSoundsForAggressivenessLevelEvents(list)...)
	return
}

// SetAggressivenessLevel1 is setter
func (sm *SoundModule) SetAggressivenessLevel1() {
	sm.getSoundsForAggressivenessLevelEvents = aggressiveness.GetSoundsForEventsLevel1
}

// SetAggressivenessLevel2 is setter
func (sm *SoundModule) SetAggressivenessLevel2() {
	sm.getSoundsForAggressivenessLevelEvents = aggressiveness.GetSoundsForEventsLevel2
}

// SetAggressivenessLevel3 is setter
func (sm *SoundModule) SetAggressivenessLevel3() {
	sm.getSoundsForAggressivenessLevelEvents = aggressiveness.GetSoundsForEventsLevel3
}
