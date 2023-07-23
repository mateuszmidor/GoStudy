package main

import (
	"reflect"
	"testing"
)

func TestSimpleInject(t *testing.T) {
	// given
	type person struct {
		FirstName string `inject:"yes"`
		LastName  string `inject:"no"`
		Age       uint8  `inject:"yes"`
		email     string
	}
	var p person

	injector := NewInjector()
	injector.Set(reflect.TypeOf(string("")), "Jessica")
	injector.Set(reflect.TypeOf(uint8(0)), uint8(22))

	// when
	injector.Inject(&p)

	// then
	if p.FirstName != "Jessica" {
		t.Errorf("Expected FirstName: %q, got: %q", "Jessica", p.FirstName)
	}

	if p.LastName != "" {
		t.Errorf("Expected LastName: %q, got: %q", "", p.FirstName)
	}

	if p.Age != 22 {
		t.Errorf("Expected Age: %v, got: %v", 22, p.Age)
	}

	if p.email != "" {
		t.Errorf("Expected email: %q, got: %q", "", p.email)
	}
}
