package main

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/configs"
)

func main() {
	log.Println("running SawmillGrpcSvc at", configs.SawmillAddr)
	log.Fatal(RunGrpcService(configs.SawmillAddr))
}
