package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/controller"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/entity"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/middleware"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	logFile, _                                 = os.Create("gin.log")
	loginController controller.LoginController = controller.NewLoginController()
	videoController controller.VideoController = setupVideoController()
)

func main() {
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)

	// CREATE PLAIN SERVER (NO DEFAULT MIDDLEWARE)
	server := gin.New()

	// RESOURCES
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("./templates/*.html")

	// MIDDLEWARE
	server.Use(gin.Recovery())
	server.Use(middleware.Logger()) // just a custom format logger
	server.Use(gindump.Dump())      // dump request and response headers and body in gin log

	// ROUTES
	// basic auth test
	server.GET("/batest", middleware.BasicAuth(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"messsage": "BasicAuth credentials OK"})
	})
	// JWT login
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{"jwt token": token}) // token should be stored by client and attached in subsequent requests headers
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	// API, requires JWT login first
	apiRoutes := server.Group("/api", middleware.AuthorizeJWT())
	apiRoutes.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})
	apiRoutes.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "video input is valid"})
		}
	})
	// HTML views, not JWT required
	viewRoutes := server.Group("/view")
	viewRoutes.GET("/videos", videoController.ShowAll)

	// START SERVER
	server.Run(":8080")
}

func setupVideoController() controller.VideoController {
	videoService := service.New()
	videoService.Save(entity.Video{
		Title:       "Morza Wszeteczne",
		Description: "Historie prosto z morza",
		URL:         "https://www.youtube.com/embed/Rt9Ne36LGUk",
		Author: entity.Person{
			FirstName: "Marcin",
			LastName:  "Mordka",
			Age:       43,
			Email:     "mordka@marcin.com",
		},
	})
	videoService.Save(entity.Video{
		Title:       "Conan: Droga do tronu",
		Description: "Zbiór opowiadań",
		URL:         "https://www.youtube.com/embed/bXho4G_VSoQ",
		Author: entity.Person{
			FirstName: "Robert",
			LastName:  "E. Howard",
			Age:       67,
			Email:     "howard@roberd.com",
		},
	})
	videoService.Save(entity.Video{
		Title:       "Achaja",
		Description: "Tom 1",
		URL:         "https://www.youtube.com/embed/3MtnbEj7gVg",
		Author: entity.Person{
			FirstName: "Andrzej",
			LastName:  "Ziemiański",
			Age:       44,
			Email:     "ziemianski@andrzej.com",
		},
	})
	return controller.New(videoService)
}
