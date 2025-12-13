package main

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/configs"
	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill"
)

func main() {
	log.Println("running SawmillGrpcSvc at", configs.SawmillAddr)
	log.Fatal(sawmill.RunGrpcService(configs.SawmillAddr))
}
