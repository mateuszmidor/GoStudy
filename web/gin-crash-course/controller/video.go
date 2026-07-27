package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/entity"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/service"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/validators"
)

type VideoController interface {
	FindAll() ([]entity.Video, error)
	ShowAll(ctx *gin.Context)
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
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

func (c *controller) FindAll() ([]entity.Video, error) {
	return c.service.FindAll()
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos, err := c.FindAll()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	data := gin.H{
		"pageTitle": "Video Page",
		"videos":    videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
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
	return c.service.Save(v)
}

func (c *controller) Update(ctx *gin.Context) error {
	// get video ID from URL param (/videos/:id)
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	// get video data from request body
	var v entity.Video
	err = ctx.ShouldBindJSON(&v) // uses field tags to unmarshall json body into struct
	if err != nil {
		return err
	}
	v.ID = id

	// REQUIRED MANUAL STEP FOR CUSTOM VALIDATION
	err = validate.Struct(v)
	if err != nil {
		return err
	}

	// update the video in database
	return c.service.Update(v)
}

func (c *controller) Delete(ctx *gin.Context) error {
	// get video ID from URL param (/videos/:id)
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	// make up video
	var v entity.Video
	v.ID = id

	// delete the videlo from database
	return c.service.Delete(v)
}
