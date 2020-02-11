package main

import (
	"bufio"
	"io"
)

// Wordfreq calculates number of each word in input reader
func Wordfreq(in io.Reader) (result map[string]int) {
	result = make(map[string]int)
	input := bufio.NewScanner(in)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		result[word]++
	}
	return
}
