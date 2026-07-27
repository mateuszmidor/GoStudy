package gear_test

import (
	"shared/gear"
	"testing"
)

func TestKeepGearAddGearDownReturnsGearDown(t *testing.T) {
	// given
	gc1 := gear.KeepCurrent
	gc2 := gear.GearDown

	// when
	sum := gc1.Add(gc2)

	//then
	if sum != gear.GearDown {
		t.Errorf("Expected gear down, got %v", sum)
	}
}

func TestGearDownAddGearDownReturnsDoubleGearDown(t *testing.T) {
	// given
	gc1 := gear.GearDown
	gc2 := gear.GearDown

	// when
	sum := gc1.Add(gc2)

	//then
	if sum != gear.DoubleGearDown {
		t.Errorf("Expected double gear down, got %v", sum)
	}
}
