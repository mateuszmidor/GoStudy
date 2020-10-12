package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/api"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/controller"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/docs"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/entity"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/middleware"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/repo"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/service"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	logFile, _                                 = os.Create("gin.log")
	loginController controller.LoginController = controller.NewLoginController()
	videoController controller.VideoController = setupVideoController()
	videoAPI        *api.VideoAPI              = api.NewVideoAPI(loginController, videoController)
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	// Swagger 2.0 meta info
	docs.SwaggerInfo.Title = "Video API"
	docs.SwaggerInfo.Description = "YouTube Video API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	// Swagger END

	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)

	// CREATE PLAIN SERVER (NO DEFAULT MIDDLEWARE)
	server := gin.New()

	// RESOURCES
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("./templates/*.html")

	// MIDDLEWARE
	server.Use(gin.Recovery())
	server.Use(middleware.Logger()) // just a custom format logger
	// server.Use(gindump.Dump())      // dump request and response headers and body in gin log

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.GET("/basicauth", middleware.BasicAuth(), videoAPI.TestBasicAuth)
			login.POST("/token", videoAPI.Authenticate)
		}

		videos := apiRoutes.Group("/videos", middleware.AuthorizeJWT())
		{
			videos.GET("", videoAPI.GetVideos)
			videos.POST("", videoAPI.CreateVideo)
			videos.PUT(":id", videoAPI.UpdateVideo)
			videos.DELETE(":id", videoAPI.DeleteVideo)
		}
	}
	// HTML views, not JWT required
	viewRoutes := server.Group("/view")
	viewRoutes.GET("/videos", videoController.ShowAll)

	// Swagger endpoint
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// START SERVER
	server.Run(":8080")
}

func setupVideoController() controller.VideoController {
	videoRepo := repo.NewVideoRepo()
	videoService := service.New(videoRepo)

	// put some initial content
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
