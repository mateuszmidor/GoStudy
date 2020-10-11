package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/entity"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(s service.VideoService) VideoController {
	return &controller{service: s}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var v entity.Video
	ctx.BindJSON(&v) // uses field tags to unmarshall json body into struct
	c.service.Save(v)
	return v
}
