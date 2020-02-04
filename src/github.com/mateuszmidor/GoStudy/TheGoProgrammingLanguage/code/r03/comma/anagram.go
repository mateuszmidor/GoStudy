package stringgames

// Anagrams checks if s1 and s2 are anagrams (s1 is made from letters of s2)
func Anagrams(s1, s2 string) bool {
	letters := make(map[int]int)

	for c := range []rune(s1) {
		letters[c]++
	}

	for c := range []rune(s2) {
		letters[c]--
	}

	for _, v := range letters {
		if v != 0 {
			return false
		}
	}
	return true
}
