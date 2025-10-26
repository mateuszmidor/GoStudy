package main

import "log"

func calcDiffBits(a, b byte) int {
	count := 0
	for i := uint(0); i < 8; i++ {
		if (a>>i)&1 != (b>>i)&1 {
			count++
		}
	}
	return count
}

func calcDiffBitsInSlices(a []byte, b []byte) (result int) {
	if len(a) != len(b) {
		log.Fatal("slices must be of same length")
	}

	for i := 0; i < len(a); i++ {
		result += calcDiffBits(a[i], b[i])
	}

	return
}
