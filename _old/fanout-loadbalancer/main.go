package main

func printer(c <-chan int, name string) {
	for val, ok := <-c; ok; val, ok = <-c {
		println(name, ": ", val)
	}
}

func fanOut(in []int, out1 chan<- int, out2 chan<- int) {
	for i := range in {
		select {
		case out1 <- i:
		case out2 <- i:
		}
	}
}

func main() {
	vals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// warning: bufferred channels would need main goroutine to wait for workers to finish
	c1 := make(chan int)
	c2 := make(chan int)
	go printer(c1, "printer1")
	go printer(c2, "printer2")

	fanOut(vals, c1, c2)

	close(c1)
	close(c2)
}
