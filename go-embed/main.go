package main

import (
	_ "embed"
	"fmt"
)

//go:embed welcome.txt
var welcomeMessage string // this variable gets populated with contents of 'welcome.txt' when building the app

func main() {
	fmt.Println(welcomeMessage)
}
