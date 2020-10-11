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
	logFile, _ := os.Create("gin.log")
	defer logFile.Close()
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)

	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middleware.Logger())    // just a custom format logger
	server.Use(middleware.BasicAuth()) // authorization in user:pass form
	server.Use(gindump.Dump())         // dump request and response headers and body in gin log

	server.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"messsage": "OK"})
	})
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "video input is valid"})
		}
	})
	server.Run(":8080")
}
