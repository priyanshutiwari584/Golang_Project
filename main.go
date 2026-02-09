package main

import (
	"log"

	"go-project.com/go-project/db"
	"go-project.com/go-project/routes"
	"go-project.com/go-project/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inintialize the database
	db.InitDB()

	// Define a new Gin router server
	var server = gin.Default()

	// Define routes
	server.GET("/", func(c *gin.Context) {
		utils.OK(c, "Sever is running successfully!", nil)
	})

	// Register routes
	routes.RegisterRoutes(server)

	err := server.Run(":3000") // listens on 0.0.0.0:8080 by default if not specified

	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
