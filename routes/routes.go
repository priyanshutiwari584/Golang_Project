package routes

import (
	"github.com/gin-gonic/gin"
	"go-project.com/go-project/controllers"
	"go-project.com/go-project/middlewares"
)

func registerEventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)

	// authenticated routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) //add middleware
	// protected routes
	authenticated.POST("/events", controllers.CreateEvent)
	authenticated.PUT("/events/:id", controllers.UpdateEvent)
	authenticated.DELETE("/events/:id", controllers.DeleteEvent)
	authenticated.POST("/events/:id/register", controllers.RegisterForEvent)
	authenticated.DELETE("/events/:id/cancel", controllers.CancellationForEvent)
}

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/users/singup", controllers.Singup)
	server.POST("/users/login", controllers.Login)
}

func registerRegistrationRoutes(server *gin.Engine) {
	server.GET("/registrations", controllers.GetAllRegisteredEvents)
}

func RegisterRoutes(server *gin.Engine) {
	registerEventRoutes(server)
	RegisterUserRoutes(server)
	registerRegistrationRoutes(server)
}
