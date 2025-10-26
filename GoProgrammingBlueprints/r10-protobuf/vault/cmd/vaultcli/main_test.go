package main

import "testing"

func TestPop(t *testing.T) {
	args := []string{"one", "two", "three"}
	var s string
	s, args = pop(args)
	if s != "one" {
		t.Errorf("Unexpected value: %s", s)
	}
	s, args = pop(args)
	if s != "two" {
		t.Errorf("Unexpected value: %s", s)
	}
	s, args = pop(args)
	if s != "three" {
		t.Errorf("Unexpected value: %s", s)
	}
	s, args = pop(args)
	if s != "" {
		t.Errorf("Unexpected value: %s", s)
	}
}
