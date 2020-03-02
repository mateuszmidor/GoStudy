package equal

import "testing"

func TestSlicesAreEqual(t *testing.T) {
	s1 := []string{"Tree", "Stone", "Grass"}
	s2 := []string{"Tree", "Stone", "Grass"}
	if !Equal(s1, s2) {
		t.Errorf("Two slices should be equal but Equal say they are not:\n%v\n%v", s1, s2)
	}
}

func TestSlicesAreNotEqual(t *testing.T) {
	s1 := []uint{1, 2, 3}
	s2 := []uint{1, 2, 4}
	if Equal(s1, s2) {
		t.Errorf("Two slices should be equal but Equal say they are not:\n%v\n%v", s1, s2)
	}
}

func TestEmptyMapsAreEqual(t *testing.T) {
	m1 := map[string]uint(nil) // nil map is equal empty map
	m2 := map[string]uint{}
	if Equal(m1, m2) {
		t.Errorf("Two maps should be equal but Equal say they are not:\n%v\n%v", m1, m2)
	}
}

func TestSameElementsMapsAreEqual(t *testing.T) {
	m1 := map[string]uint{"A": 33, "B": 34}
	m2 := map[string]uint{"A": 33, "B": 34}
	if !Equal(m1, m2) {
		t.Errorf("Two maps should be equal but Equal say they are not:\n%v\n%v", m1, m2)
	}
}

func TestDifferentNumberOfElementsMapsAreNotEqual(t *testing.T) {
	m1 := map[string]uint{"A": 33, "B": 34}
	m2 := map[string]uint{"A": 33}
	if Equal(m1, m2) {
		t.Errorf("Two maps should be equal but Equal say they are not:\n%v\n%v", m1, m2)
	}
}

func TestDifferentElementsMapsAreNotEqual(t *testing.T) {
	m1 := map[string]uint{"A": 33, "B": 34}
	m2 := map[string]uint{"A": 33, "B": 500}
	if Equal(m1, m2) {
		t.Errorf("Two maps should be equal but Equal say they are not:\n%v\n%v", m1, m2)
	}
}
