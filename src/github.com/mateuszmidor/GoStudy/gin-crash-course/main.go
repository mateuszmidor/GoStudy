package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/controller"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/middleware"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var videoController controller.VideoController = controller.New(service.New())

func main() {
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middleware.Logger())
	server.Use(middleware.BasicAuth())
	server.Use(gindump.Dump())

	server.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"messsage": "OK"})
	})
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})
	setupLogOutput()
	server.Run(":8080")
}

func setupLogOutput() {
	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)
}
