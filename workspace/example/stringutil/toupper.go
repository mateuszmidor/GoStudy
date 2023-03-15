package stringutil

import "unicode"

// ToUpper uppercases all the runes in its argument string.
// Function added by us
func ToUpper(s string) string {
	r := []rune(s)
	for i := range r {
		r[i] = unicode.ToUpper(r[i])
	}
	return string(r)
}
