package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "8080", "-port=80")
	data := flag.String("data", "../../../data/", "-data=./data")
	flag.Parse()

	runWEB(*port, *data)
}

func runWEB(port, data_dir string) {
	log.SetPrefix("[APP] ")

	router := gin.Default()
	router.LoadHTMLGlob("data/*.html")
	router.GET("/", newIndexHandler())
	router.GET("/api/find", newFindRequestHandler(data_dir))
	router.Run(":" + port)
}
