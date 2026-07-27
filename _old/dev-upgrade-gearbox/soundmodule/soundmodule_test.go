package soundmodule_test

import (
	"shared/events"
	"shared/sounds"
	"soundmodule"
	"testing"
)

func TestShouldPipeBlastOnGearDownInAggressivenessLevel3(t *testing.T) {
	// given
	testCases := map[events.Event]sounds.Sound{
		events.DoubleGearDown: sounds.PipeBlast,
		events.GearDown:       sounds.PipeBlast,
		events.GearUp:         sounds.Silence,
	}
	sm := soundmodule.NewSoundModule()
	sm.SetAggressivenessLevel3()

	for event, expectedSound := range testCases {
		// when
		actualSounds := sm.HandleEvents(events.Events{event})

		// then
		if actualSounds.String() != string(expectedSound) {
			t.Errorf("Expected %q, got: %q", expectedSound, actualSounds)
		}
	}
}

func TestShouldKeepSilenceOnGearDownInAggressivenessLevel1(t *testing.T) {
	// given
	testCases := map[events.Event]sounds.Sound{
		events.DoubleGearDown: sounds.Silence,
		events.GearDown:       sounds.Silence,
		events.GearUp:         sounds.Silence,
	}
	sm := soundmodule.NewSoundModule()
	sm.SetAggressivenessLevel1()

	for event, expectedSound := range testCases {
		// when
		actualSounds := sm.HandleEvents(events.Events{event})

		// then
		if actualSounds.String() != string(expectedSound) {
			t.Errorf("Expected %q, got: %q", expectedSound, actualSounds)
		}
	}
}
