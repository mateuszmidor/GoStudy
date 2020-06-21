package main

import (
	"fmt"
	"strings"
)

// index returns first match index
func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// include returns if vs contains t
func include(vs []string, t string) bool {
	return index(vs, t) >= 0
}

// any returns if any string in slice satisfies predicate
func any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// transform transforms input slice using f
func transform(vs []string, f func(string) string) []string {
	out := make([]string, len(vs))
	for i, v := range vs {
		out[i] = f(v)
	}
	return out
}

func main() {
	strs := []string{"peach", "apple", "pear", "plum"}

	fmt.Println(index(strs, "pear"))

	fmt.Println(include(strs, "grape"))

	fmt.Println(any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(transform(strs, strings.ToUpper))
}
