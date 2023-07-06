package main

import (
	"context"
	"fmt"
	"time"
)

// context.WithValue key should be user-defined type not just string
var val = struct{ key string }{"Age"}

// wait until ctx.Done() is signalled
func waitDone(ctx context.Context, s string) {
	<-ctx.Done()
	fmt.Printf("%s - Age: %v\n", s, ctx.Value(val))
}

func main() {
	// a hierarchy : ctxRoot <- ctx1 <- ctx2 <- ctx3
	// when parent context gets cancelled/timeout, all children contexts also get cancelled in a waterfall manner
	ctxRoot := context.WithValue(context.Background(), val, 32) // value will propagate to children contexts ctx2 and ctx3
	ctx1, cancelCtx1 := context.WithCancel(ctxRoot)
	ctx2, cancalCtx2 := context.WithTimeout(ctx1, 4*time.Second)
	ctx3 := context.WithValue(ctx2, val, 87) // override value from ctxRoot

	defer cancalCtx2() // good practice is to cancel context upon finishing processing so context doesnt leak
	go waitDone(ctx3, "ctx3")
	go waitDone(ctx2, "ctx2")
	go waitDone(ctx1, "ctx1")

	println("Wait 5 sec and cancel ctx1, but ctx2 will timeout sooner and cancel ctx3") // so the order: ctx3, ctx2, ctx1
	time.Sleep(time.Second * 5)
	cancelCtx1() // will also cancel child context ctx3

	// let the stdout flush
	time.Sleep(time.Second)
}
