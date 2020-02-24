package examples

import (
	"testing"
)

func TestFibb(t *testing.T) {
	expected := uint(1134903170)
	actual := fibb(45)
	if actual != expected {
		t.Errorf("actual != expected: %d != %d\n", actual, expected)
	}
}

func BenchmarkFibb(b *testing.B) {
	fibb(45)
}

func BenchmarkAfter(b *testing.B) {
	after()
}

func BenchmarkAlloc(b *testing.B) {
	_ = allocLots()
}
