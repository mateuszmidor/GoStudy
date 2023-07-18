//go:build fast

package concat_test

import (
	"strings"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	builder := strings.Builder{}
	for i := 0; i < b.N; i++ {
		builder.WriteRune('a')
	}
	_ = builder.String()
}
