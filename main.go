package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/docs"
	"gitlab.com/pragmaticreviews/golang-gin-poc/repository"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

var (
	exerciseRepository repository.ExerciseRepository  = repository.NewExerciseRepository()
	exerciseService    service.ExerciseService        = service.New(exerciseRepository)
	exerciseController controller.ExcerciseController = controller.New(exerciseService)
	loginService       service.LoginService           = service.NewLoginService()
	jwtService         service.JWTService             = service.NewJWTService()
	loginController    controller.LoginController     = controller.NewLoginController(loginService, jwtService)
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
	defer exerciseRepository.CloseDB()
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
	{ // Login Endpoint: Authentication + Token creation
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
		apiRoutes.GET("/exercises", func(context *gin.Context) {
			context.JSON(200, exerciseController.FindAll())
		})

		apiRoutes.POST("/exercises", func(context *gin.Context) {
			err := exerciseController.Save(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Excercise Input is Valid!"})
			}
		})

		apiRoutes.PUT("/exercises/:id", func(context *gin.Context) {
			err := exerciseController.Update(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Excercise Update is Valid!"})
			}
		})

		apiRoutes.DELETE("/exercises/:id", func(context *gin.Context) {
			err := exerciseController.Delete(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Excercise Deleting is Valid!"})
			}
		})
	}
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/exercises", exerciseController.ShowAll)
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8080")
}
