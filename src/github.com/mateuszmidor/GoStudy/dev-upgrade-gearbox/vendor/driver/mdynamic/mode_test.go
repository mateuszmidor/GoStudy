package mdynamic

import (
	"driver/uślizg"
	"testing"
)

func TestGearChangeShouldBeAllowedWhenNoUślizg(t *testing.T) {
	// given
	var mdynamic Mode
	mdynamic.Enable()

	// when
	isGearChangeAllowed := mdynamic.IsGearChangeAllowed(uślizg.AngularSpeedForNoUślizg)

	// then
	if !isGearChangeAllowed {
		t.Errorf("For no uślizg a gear change should be allowed")
	}
}

func TestGearChangeShouldBeDisallowedWhenUślizg(t *testing.T) {
	// given
	var mdynamic Mode
	mdynamic.Enable()

	// when
	isGearChangeAllowed := mdynamic.IsGearChangeAllowed(uślizg.AngularSpeedForUślizg)

	// then
	if isGearChangeAllowed {
		t.Errorf("For uślizg a gear change should be disallowed")
	}
}
