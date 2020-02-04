package stringgames

import (
	"log"
	"math"
)

func getNumTriplets(s string) int {
	return int(math.Ceil(float64(len(s)) / 3))
}

func getTriplet(s string, nTriplet int) (string, bool) {
	if nTriplet < 0 {
		log.Printf("Cant get negative triplet %d from %q\n", nTriplet, s)
		return "", false
	}

	numTriplets := getNumTriplets(s)
	if nTriplet >= numTriplets {
		log.Printf("Cant get triplet %d, there is only %d triplets in %q\n", nTriplet, numTriplets, s)
		return "", false
	}

	strlen := len(s)
	nTripletFromStringEnd := numTriplets - nTriplet - 1
	stop := strlen - nTripletFromStringEnd*3
	start := strlen - (nTripletFromStringEnd+1)*3
	if start < 0 {
		start = 0
	}

	return s[start:stop], true
}

// Comma turns 123456 -> 123,456
func Comma(in string) string {
	numTriples := getNumTriplets(in)
	out, _ := getTriplet(in, numTriples-1)
	for i := numTriples - 2; i >= 0; i-- {
		triplet, _ := getTriplet(in, i)
		if triplet != "-" {
			out = triplet + "," + out
		} else {
			out = triplet + out
		}
	}

	return out
}
