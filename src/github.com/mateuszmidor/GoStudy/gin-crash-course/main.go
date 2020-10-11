package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/controller"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/service"
)

var videoController controller.VideoController = controller.New(service.New())

func main() {
	server := gin.Default()
	server.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"messsage": "OK"})
	})
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})
	server.Run(":8080")
}
