package stringgames

import "bytes"

// Reverse string
func Reverse(in string) string {
	var buf bytes.Buffer
	inAsRunes := []rune(in)
	buf.Grow(len(in))
	for i := len(inAsRunes) - 1; i >= 0; i-- {
		buf.WriteRune(inAsRunes[i])
	}
	return buf.String()
}
