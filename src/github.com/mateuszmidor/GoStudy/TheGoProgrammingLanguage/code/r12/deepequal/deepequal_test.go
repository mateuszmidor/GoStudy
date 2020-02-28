package deepequal_test

import (
	"reflect"
	"testing"
)

func TestDeepEqualTrue(t *testing.T) {
	s1 := []string{"Tree", "Stone", "Grass"}
	s2 := []string{"Tree", "Stone", "Grass"}
	if !reflect.DeepEqual(s1, s2) {
		t.Errorf("Two slices should be equal but DeepEqual say they are not:\n%v\n%v", s1, s2)
	}
}

func TestDeepEqualFalse(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	if reflect.DeepEqual(s1, s2) {
		t.Errorf("Two slices should be equal but DeepEqual say they are not:\n%v\n%v", s1, s2)
	}
}
