package main

import (
	"fmt"
	"iter"
)

// type of this iterator is: iter.Seq[rune].
// rule: iterator calls "yield" function with every value it wants to return
func iterator1(yield func(rune) bool) {
	for c := 'a'; c <= 'z'; c++ {
		if !yield(c) {
			return
		}
	}
}

// type of this iterator is: iter.Seq2[int, int].
func iterator2(yield func(int, int) bool) {
	for i := 1; i <= 10; i++ {
		if !yield(i, i*i) {
			return
		}
	}
}

func demoPushIterators() {
	fmt.Println("PUSH ITERATORS")

	for val := range iterator1 {
		fmt.Println(string(val))
	}

	for num, numSquared := range iterator2 {
		fmt.Println(num, "->", numSquared)
	}
}

func demoPullIterators() {
	fmt.Println("PULL ITERATORS")

	next, stop := iter.Pull2(iterator2) // convert push iterator into pull iterator
	defer stop()
	for num, numSquared, ok := next(); ok; num, numSquared, ok = next() {
		fmt.Println(num, "->", numSquared)
	}
}

func main() {
	demoPushIterators()
	demoPullIterators()
}
