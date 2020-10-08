package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	runWEB()
}

func runWEB() {
	log.SetPrefix("[APP] ")

	router := gin.Default()
	router.LoadHTMLGlob("data/*.html")
	router.GET("/", newIndexHandler())
	router.GET("/api/find", newFindRequestHandler())
	router.Run(":9000")
}
