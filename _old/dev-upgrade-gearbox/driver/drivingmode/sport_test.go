package drivingmode_test

import (
	"driver/aggressiveness"
	"driver/drivingmode"
	"shared/gear"
	"driver/types"
	"driver/util"
	"shared/gas"
	"testing"
)

var gasKickDown1 = gas.New(0.4)
var gasKickDown2 = gas.New(0.8)

func TestSportGetOptimalRange(t *testing.T) {
	// given
	var aggressiveness aggressiveness.Level = aggressiveness.Stub{Multiplier: 1.5}
	var sport drivingmode.Mode = drivingmode.NewSport(1000.0, 2000.0, gasKickDown1, gasKickDown2)

	// when
	var minRPM, maxRPM types.RPM = sport.GetOptimalRPM(aggressiveness)

	// then
	if !util.IsEqual(1500.0, minRPM) {
		t.Errorf("Expected minRPM %f, got %f", 1500.0, minRPM)
	}

	if !util.IsEqual(3000.0, maxRPM) {
		t.Errorf("Expected maxRPM %f, got %f", 3000.0, maxRPM)
	}
}

func TestSportNoKickDownReturnsKeepGear(t *testing.T) {
	// given
	var sport drivingmode.Mode = drivingmode.NewSport(1000.0, 2000.0, gasKickDown1, gasKickDown2)

	// when
	gc := sport.GetKickDownGearChange(gas.Zero)

	// then
	if gc != gear.KeepCurrent {
		t.Errorf("Expected keep gear, got %v", gc)
	}
}

func TestSportKickDown1ReturnsGearDown(t *testing.T) {
	// given
	var sport drivingmode.Mode = drivingmode.NewSport(1000.0, 2000.0, gasKickDown1, gasKickDown2)

	// when
	gc := sport.GetKickDownGearChange(gasKickDown1)

	// then
	if gc != gear.GearDown {
		t.Errorf("Expected gear down, got %v", gc)
	}
}

func TestSportKickDown2ReturnsDoubleGearDown(t *testing.T) {
	// given
	var sport drivingmode.Mode = drivingmode.NewSport(1000.0, 2000.0, gasKickDown1, gasKickDown2)

	// when
	gc := sport.GetKickDownGearChange(gasKickDown2)

	// then
	if gc != gear.DoubleGearDown {
		t.Errorf("Expected double gear down, got %v", gc)
	}
}
