package main

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				println("received job", j)
			} else {
				println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		println("sent job", j)
	}

	close(jobs)
	println("sent all jobs")

	<-done
}
