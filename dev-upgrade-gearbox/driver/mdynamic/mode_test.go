package mdynamic

import (
	"driver/uslizg"
	"testing"
)

func TestGearChangeShouldBeAllowedWhenNoUślizg(t *testing.T) {
	// given
	var mdynamic Mode
	mdynamic.Enable()

	// when
	isGearChangeAllowed := mdynamic.IsGearChangeAllowed(uslizg.AngularSpeedForNoUślizg)

	// then
	if !isGearChangeAllowed {
		t.Errorf("For no uslizg a gear change should be allowed")
	}
}

func TestGearChangeShouldBeDisallowedWhenUślizg(t *testing.T) {
	// given
	var mdynamic Mode
	mdynamic.Enable()

	// when
	isGearChangeAllowed := mdynamic.IsGearChangeAllowed(uslizg.AngularSpeedForUślizg)

	// then
	if isGearChangeAllowed {
		t.Errorf("For uslizg a gear change should be disallowed")
	}
}
