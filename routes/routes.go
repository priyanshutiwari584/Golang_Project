package routes

import (
	"github.com/gin-gonic/gin"
	"go-project.com/go-project/controllers"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.POST("/events", controllers.CreateEvent)
	server.GET("/events/:id", controllers.GetEvent)
	server.PUT("/events/:id", controllers.UpdateEvent)
	server.DELETE("/events/:id", controllers.DeleteEvent)
}

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/users/singup", controllers.Singup)
	server.POST("/users/login", controllers.Login)
}
