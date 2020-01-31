package main

// Project: flag & pointers example
// Usage: ./echo4 -s ! red green blue

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", " ", "separator") // sep is a pointer to string
func main() {
	flag.Parse() // fill "sep" with actual flag value
	fmt.Println(strings.Join(flag.Args(), *sep))
}

