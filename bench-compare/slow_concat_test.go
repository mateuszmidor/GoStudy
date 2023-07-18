//go:build slow

package concat_test

import (
	"fmt"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	s := ""
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("%s%s", s, "a")
	}
}
