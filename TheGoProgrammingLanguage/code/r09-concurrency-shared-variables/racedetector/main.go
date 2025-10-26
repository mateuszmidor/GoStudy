// Project: race detector demo
// Usage: go run -race .
package main

import (
	"fmt"
	"sync"
)

var counter uint64

func incCounterTimes(numTimes int) {
	for i := 0; i < numTimes; i++ {
		// atomic.AddUint64(&counter, 1) // this will not do data race
		counter++
	}
}

func main() {

	const numParallel = 10
	const numIncrements = 100000
	var wg sync.WaitGroup

	for i := 0; i < numParallel; i++ {
		wg.Add(1)
		go func() {
			incCounterTimes(numIncrements)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Num parallel: %d, expected count %d, actual %d\n", numParallel, numParallel*numIncrements, counter)

}
