package string

// Capital R means this function is exported and visible from other packages
func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b) / 2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}