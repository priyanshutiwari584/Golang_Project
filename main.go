package main

import (
	"go-project/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func SuccessResponse[T any](message string, data T) APIResponse[T] {
	return APIResponse[T]{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err any) APIResponse[any] {
	return APIResponse[any]{
		Status:  "error",
		Message: message,
		Error:   err,
	}
}

func main() {
	// Define a new Gin router server
	var server = gin.Default()

	// Define routes
	server.GET("/", func(c *gin.Context) {
		resp := SuccessResponse("Server is running successfully!", any(nil))
		c.JSON(http.StatusOK, resp)
	})

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":3000") // listens on 0.0.0.0:8080 by default

	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, SuccessResponse("Events Fetched successfully!", events))
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse("Unable to create Event", err))
	}

	event.ID = 1
	event.UserID = 1
	context.JSON(http.StatusCreated, SuccessResponse("Event created successfully", event))
}
