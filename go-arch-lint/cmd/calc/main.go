// Command calc is the entrypoint of the go-arch-lint demo.
// It wires the vertical slices (add, multiply, power)
// into a standard-library HTTP server on port 9090.
package main

import (
	"log"
	"net/http"

	"go-arch-lint/slices/add"
	"go-arch-lint/slices/multiply"
	"go-arch-lint/slices/power"
)

const addr = ":9090"

func main() {
	mux := http.NewServeMux()

	add.Register(mux)
	multiply.Register(mux)
	power.Register(mux)

	log.Printf("calc listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
