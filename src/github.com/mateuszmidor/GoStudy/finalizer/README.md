# runtime.SetFinalizer

<https://youtu.be/us9hfJqncV8?t=1196>

## Code

```go
package main

import (
	"runtime"
	"time"
)

// s is the finalized object
func stringFinalizer(s *string) {
	println("Finalizer:", *s)
}

func main() {
	s := "carburetor"

	// when GC finds out "s" is to be cleaned up, it will run its finalizer first
	runtime.SetFinalizer(&s, stringFinalizer)

	// make sure the GC is run
	runtime.GC()

	// give some time for finalizer to do the work
	time.Sleep(time.Millisecond * 100)
}
```

Output:

```Finalizer: carburetor```
