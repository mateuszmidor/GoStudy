// Project: build a long pipeline of goroutines connected with channels and send msg through it
// Usage: go run .
// 1mln pipes takes 2.6GB and sending struct{} 1-3us, sending int 350ms
// 5mln takes all of 8GB ram, lots of swap, takes forever
// 10mln pipes killed by kernel (out of memory)
package main

import (
	"fmt"
	"time"
)

type chantype int

func main() {
	const numPipes = 1000000
	pipelineStart := make(chan chantype)
	var pipelineEnd chan chantype

	begin := pipelineStart
	for i := 0; i < numPipes; i++ {
		pipelineEnd = make(chan chantype)
		go passOver(begin, pipelineEnd)
		begin = pipelineEnd // end becomes new begin
	}

	timeStart := time.Now()
	pipelineStart <- 42
	result := <-pipelineEnd
	delta := time.Now().Sub(timeStart)
	fmt.Printf("Time %v, num pipes %d, value received %d\n", delta, numPipes, result)
}

func passOver(in <-chan chantype, out chan<- chantype) {
	val := <-in
	out <- val
}
