package aggressiveness

import (
	"driver/util"
	"testing"
)

func TestGetRPMMultiplier(t *testing.T) {
	// given
	al := Level1{}

	// when
	rpmMultiplier := al.GetRPMMultiplier()

	// then
	if !util.IsEqual(rpmMultiplier, 1.0) {
		t.Errorf("Expected rpmMultiplier %f, got %f", 1.0, rpmMultiplier)
	}
}
