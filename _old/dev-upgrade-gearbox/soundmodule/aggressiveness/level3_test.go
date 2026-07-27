package aggressiveness_test

import (
	"shared/events"
	"soundmodule/aggressiveness"
	"shared/sounds"
	"testing"
)

func TestShouldPipeBlastOnGearDown(t *testing.T) {
	// given
	testCases := map[events.Event]sounds.Sound{
		events.DoubleGearDown: sounds.PipeBlast,
		events.GearDown:       sounds.PipeBlast,
		events.GearUp:         sounds.Silence,
	}
	soundGetter := aggressiveness.GetSoundsForEventsLevel3

	for event, expectedSound := range testCases {
		// when
		actualSounds := soundGetter(events.Events{event})

		// then
		if actualSounds.String() != string(expectedSound) {
			t.Errorf("Expected %q, got: %q", expectedSound, actualSounds)
		}
	}
}
