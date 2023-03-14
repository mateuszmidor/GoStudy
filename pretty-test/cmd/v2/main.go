package main

import (
	"os"

	"github.com/mateuszmidor/pretty/v2"
)

func main() {
	pretty.Print(os.Stdout, 4)
}
