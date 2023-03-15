package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	// ToUpper function was added by us in the local copy of golang.org/x/example module
	fmt.Println(stringutil.ToUpper("Hello"))
}
