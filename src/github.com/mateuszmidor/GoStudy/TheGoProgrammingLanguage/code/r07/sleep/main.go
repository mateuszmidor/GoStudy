// Project: flag utility usage
// Usage: go run . -period 50ms
//                 -period 2m30s
//                 -period 1.5h
package main

import (
	"flag"
	"fmt"
	"time"
)

//						   -period   1s              help text
var period *time.Duration = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse() // here the "period" variable is populated with value
	fmt.Printf("Sleeping for %v...\n", *period)
	time.Sleep(*period)
	fmt.Println("Done.")
}
