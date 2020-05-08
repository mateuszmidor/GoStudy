package main

import (
	"time"
)

func main() {
	println("Will allocate 1MB/s\n")

	for {
		// p = initialAllocation
		time.Sleep(1000 * time.Millisecond)

		// alloc 1 MB
		_ = make([]int8, 1024*1024)
	}
}
