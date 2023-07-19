package main

import (
	"context"
	"errors"
	"log"

	"golang.org/x/sync/errgroup"
)

const dividend = 100

func main() {
	// define work source
	dividers := make(chan int)

	// define worker
	worker := func() error {
		for divider := range dividers { // loop breaks when channel is closed
			if divider == 0 {
				return errors.New("division by zero. Finishing all workers")
			}
			log.Printf("%d/%d = %v", dividend, divider, float32(dividend)/float32(divider))
		}
		return nil
	}

	// run workers
	g, groupContext := errgroup.WithContext(context.Background()) // closes groupContext.Done() when any func executed with g.Go(...) returns error
	g.Go(worker)
	g.Go(worker)

	// define producer
	producer := func() error {
		defer close(dividers) // this will finish all workers
		for i := 10; i >= -10; i-- {
			select {
			case <-groupContext.Done(): // means: when some worker returned with error
				return groupContext.Err()
			case dividers <- i:
			}
		}
		return nil
	}

	// run producer
	g.Go(producer)

	// wait all workers and producer goroutines are done
	err := g.Wait()
	if err != nil {
		log.Println("error:", err)
	}
}
