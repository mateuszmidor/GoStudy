package gearadjustment_test

import (
	"driver/drivingmode"
	"driver/gearadjustment"
	"shared/gas"
	"shared/gear"
	"testing"
)

func TestDoubleGearDownForCrazyDriveMode(t *testing.T) {
	// given
	dm := drivingmode.Stub{
		KickDown: gear.DoubleGearDown,
	}
	gas := gas.Full

	// when
	change := gearadjustment.AdjustForKickDown(dm, gas)

	// then
	if change != gear.DoubleGearDown {
		t.Errorf("Expected double gear down, got %v", change)
	}
}
