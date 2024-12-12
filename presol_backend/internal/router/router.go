package router

import (
	"presolai/internal/controllers"
	"presolai/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.ErrorHandle())
	r.Use(middlewares.Cors())

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.POST("/", controllers.CreateUser)
	}

	commentRoutes := r.Group("/comment")
	{
		commentRoutes.GET("/", controllers.GetAllComments)
		commentRoutes.POST("/", controllers.CreateComment)
	}

	eventRoutes := r.Group("/event")
	{
		eventRoutes.GET("/", controllers.GetAllComments)
		eventRoutes.POST("/", controllers.CreateEvent)
	}

	return r
}
