package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/entity"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/service"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate // custom video.Title validator

func New(s service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("StartsWithCapital", validators.ValidateStartsWithCapital)
	return &controller{service: s}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var v entity.Video
	err := ctx.ShouldBindJSON(&v) // uses field tags to unmarshall json body into struct
	if err != nil {
		return err
	}
	// REQUIRED MANUAL STEP FOR CUSTOM VALIDATION
	err = validate.Struct(v)
	if err != nil {
		return err
	}
	c.service.Save(v)
	return nil
}
