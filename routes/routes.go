package routes

import (
	"github.com/gin-gonic/gin"
	"go-project.com/go-project/controllers"
	"go-project.com/go-project/middlewares"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)

	// authenticated routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) //add middleware
	// protected routes
	authenticated.POST("/events", controllers.CreateEvent)
	authenticated.PUT("/events/:id", controllers.UpdateEvent)
	authenticated.DELETE("/events/:id", controllers.DeleteEvent)
}

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/users/singup", controllers.Singup)
	server.POST("/users/login", controllers.Login)
}
