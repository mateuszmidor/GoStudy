package slices

import "unicode"

func squashSpaces(in string) (out string) {
	lastSpace := false
	for _, c := range in {
		if !lastSpace || !unicode.IsSpace(c) {
			out += string(c)
		}
		lastSpace = unicode.IsSpace(c)
	}
	return
}
