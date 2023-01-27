package main

import (
	"fmt"

	"github.com/mateuszmidor/GoStudy/viper-config/config"
)

func main() {
	cfg, err := config.LoadFromFile("config.yaml", config.FormatYAML)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", cfg)
}
