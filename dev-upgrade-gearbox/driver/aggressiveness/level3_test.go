package aggressiveness_test

import (
	"driver/aggressiveness"
	"driver/util"
	"testing"
)

func TestGetRPMMultiplier(t *testing.T) {
	// given
	al := aggressiveness.Level3{}

	// when
	rpmMultiplier := al.GetRPMMultiplier()

	// then
	if !util.IsEqual(rpmMultiplier, 1.3) {
		t.Errorf("Expected rpmMultiplier %f, got %f", 1.3, rpmMultiplier)
	}
}
