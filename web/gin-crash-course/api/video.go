package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/controller"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/dto"
)

type VideoAPI struct {
	loginController controller.LoginController
	videoController controller.VideoController
}

func NewVideoAPI(loginController controller.LoginController, videoController controller.VideoController) *VideoAPI {
	return &VideoAPI{
		loginController: loginController,
		videoController: videoController,
	}
}

// Paths Information

// TestBasicAuth godoc
// @Summary Checks BasicAuth functionality
// @Description Just for BasicAuth middleware testing purposes
// @ID TestBasicAuth
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param Authorization header string true "Use: Basic YWRtaW46cGFzcw=="
// @Success 200 {object} dto.JWT
// @Failure 401 {object} dto.Response
// @Router /auth/basicauth [get]
func (api *VideoAPI) TestBasicAuth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, dto.NewResponse("BasicAuth credentials OK"))
}

// Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username query string true "admin"
// @Param password query string true "pass"
// @Success 200 {object} dto.JWT
// @Failure 401 {object} dto.Response
// @Router /auth/token [post]
func (api *VideoAPI) Authenticate(ctx *gin.Context) {
	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, dto.NewJWT(token))
	} else {
		ctx.JSON(http.StatusUnauthorized, dto.NewResponse("Not Authorized"))
	}
}

// GetVideos godoc
// @Security bearerAuth
// @Summary List existing videos
// @Description Get all the existing videos
// @Tags videos,list
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Video
// @Failure 401 {object} dto.Response
// @Router /videos [get]
func (api *VideoAPI) GetVideos(ctx *gin.Context) {
	videos, err := api.videoController.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewResponse(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, videos)
	}
}

// CreateVideo godoc
// @Security bearerAuth
// @Summary Create new videos
// @Description Create a new video
// @Tags videos,create
// @Accept  json
// @Produce  json
// @Param video body entity.Video true "Create video"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos [post]
func (api *VideoAPI) CreateVideo(ctx *gin.Context) {
	err := api.videoController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewResponse(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, dto.NewResponse("Video created"))
	}
}

// UpdateVideo godoc
// @Security bearerAuth
// @Summary Update videos
// @Description Update a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Param video body entity.Video true "Update video"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos/{id} [put]
func (api *VideoAPI) UpdateVideo(ctx *gin.Context) {
	err := api.videoController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewResponse(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, dto.NewResponse("Video updated"))
	}
}

// DeleteVideo godoc
// @Security bearerAuth
// @Summary Remove videos
// @Description Delete a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos/{id} [delete]
func (api *VideoAPI) DeleteVideo(ctx *gin.Context) {
	err := api.videoController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewResponse(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, dto.NewResponse("Video deleted"))
	}
}
