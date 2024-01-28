package gear_test

import (
	"shared/gear"
	"testing"
)

func TestShouldPanicOnSettingBelowRange(t *testing.T) {
	defer func() {
		p := recover()
		if p == nil {
			t.Errorf("Should panic on setting gear below minimum")
		}
	}()

	// given
	g := gear.New(1, 5)

	// when
	g.Set(0)

	// then
	// expect panic
}

func TestShouldPanicOnSettingAboveRange(t *testing.T) {
	defer func() {
		p := recover()
		if p == nil {
			t.Errorf("Should panic on setting gear past maximum")
		}
	}()

	// given
	g := gear.New(1, 5)

	// when
	g.Set(6)

	// then
	// expect panic
}

func TestShouldNotPassMinimum(t *testing.T) {
	// given
	g := gear.New(1, 5).Set(1)

	// when
	g = g.Down()

	// then
	if g != gear.New(1, 5).Set(1) {
		t.Errorf("Expected gear 1, got %v", g)
	}
}

func TestShouldNotPassMaximum(t *testing.T) {
	// given
	g := gear.New(1, 5).Set(5)

	// when
	g = g.Up()

	// then
	if g != gear.New(1, 5).Set(5) {
		t.Errorf("Expected gear 5, got %v", g)
	}
}

func TestShoulApplyChangeUpCorrectly(t *testing.T) {
	// given
	g := gear.New(1, 5).Set(3)
	c := gear.Change(1)

	// when
	g = g.ApplyChange(c)

	// then
	if g != gear.New(1, 5).Set(4) {
		t.Errorf("Expected gear 4, got %v", g)
	}
}

func TestShoulApplyChangeDownCorrectly(t *testing.T) {
	// given
	g := gear.New(1, 5).Set(3)
	c := gear.Change(-1)

	// when
	g = g.ApplyChange(c)

	// then
	if g != gear.New(1, 5).Set(2) {
		t.Errorf("Expected gear 2, got %v", g)
	}
}
