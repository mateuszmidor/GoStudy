package main

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/configs"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/sailworks"
)

func main() {
	log.Println("running SailworksGrpcSvc at", configs.SailworksAddr)
	log.Fatal(sailworks.RunSailworksGrpcSvc(configs.SailworksAddr))
}
