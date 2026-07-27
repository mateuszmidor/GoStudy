// Project: show how to call C function from go. In this case: bzip2 compress function
// Usage: go run .
package main

import (
	"bzip"
	"io"
	"log"
	"os"
)

func main() {
	// fmt.Fprintf(os.Stderr, "Warning: for some reason this app ends with %q\n", "panic: runtime error: cgo argument has Go pointer to Go pointer")

	// io.Copy(os.Stdout, os.Stdin) // for testing

	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}

	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: closing: %v\n", err)
	}
}
