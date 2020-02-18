// Project: read only/write only channels building a pipeline
// Usage: go ru n.
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}

func counter(c chan<- int) {
	for i := 0; i <= 100; i++ {
		c <- i
	}
	close(c)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(c <-chan int) {
	for x := range c {
		fmt.Println(x)
	}
}
