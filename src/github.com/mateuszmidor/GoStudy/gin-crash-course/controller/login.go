package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/dto"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

func NewLoginController() LoginController {
	return &loginController{
		loginService: service.NewLoginService(),
		jwtService:   service.NewJWTService(),
	}
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func (c *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := c.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return c.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
