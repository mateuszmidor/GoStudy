package main

import (
	"log"

	"github.com/mateuszmidor/GoStudy/modular-monolith/configs"
	"github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/sailworks"
)

func main() {
	log.Println("running sailworks svc at", configs.SailworksAddr)
	log.Fatal(sailworks.RunSailworksSvc(configs.SailworksAddr))
}
