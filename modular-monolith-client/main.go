package main

import (
	"fmt"

	"github.com/mateuszmidor/GoStudy/modular-monolith/pkg/clients"
)

func main() {
	sailworks := clients.NewSailworksGrpc(":9001")
	sails, err := sailworks.GetSails(3)
	fmt.Println(sails, err)
}
