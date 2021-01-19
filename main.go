package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/repository"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
	"gitlab.com/pragmaticreviews/golang-gin-poc/docs" 

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutPut() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

		// Swagger 2.0 Meta Information

		docs.SwaggerInfo.Title = "Pacticas Go2 - Video API"
		docs.SwaggerInfo.Description = "Pacticas Go2 "
		docs.SwaggerInfo.Version = "2.0"
		docs.SwaggerInfo.Host = "localhost:8080"
		docs.SwaggerInfo.BasePath = "/api"
		docs.SwaggerInfo.Schemes = []string{"http"}
	defer videoRepository.CloseDB()
	setupLogOutPut()
	server := gin.New()

	//server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	server.Use(gin.Recovery(), gin.Logger())
	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	// Login Endpoint: Authentication + Token creation
/*	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})*/

	apiRoutes := server.Group("/api")
	{	// Login Endpoint: Authentication + Token creation
		server.POST("/login", func(ctx *gin.Context) {
			token := loginController.Login(ctx)
			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, nil)
			}
		})
		apiRoutes.GET("/videos", func(context *gin.Context) {
			context.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(context *gin.Context) {
			err := videoController.Save(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!"})
			}
		})

		apiRoutes.PUT("/videos/:id", func(context *gin.Context) {
			err := videoController.Update(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Video Update is Valid!"})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(context *gin.Context) {
			err := videoController.Delete(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Video Deleting is Valid!"})
			}
		})
	}
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	server.Run(":8080")
}
