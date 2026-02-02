package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

// GENERIC RESPONSE
func respond(c *gin.Context, statusCode int, payload APIResponse) {
	c.JSON(statusCode, payload)
}

func abortRespond(c *gin.Context, statusCode int, payload APIResponse) {
	c.AbortWithStatusJSON(statusCode, payload)
}

// SUCCESS RESPONSES
func OK(c *gin.Context, message string, data any) {
	respond(c, http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, message string, data any) {
	respond(c, http.StatusCreated, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// CLIENT ERROR RESPONSES
func BadRequest(c *gin.Context, message string, err any) {
	respond(c, http.StatusBadRequest, APIResponse{
		Status:  "error",
		Message: message,
		Error:   err,
	})
}

func Unauthorized(c *gin.Context, message string) {
	abortRespond(c, http.StatusUnauthorized, APIResponse{
		Status:  "error",
		Message: message,
	})
}

func Conflict(c *gin.Context, message string) {
	respond(c, http.StatusConflict, APIResponse{
		Status:  "error",
		Message: message,
	})
}

func Forbidden(c *gin.Context, message string) {
	respond(c, http.StatusForbidden, APIResponse{
		Status:  "error",
		Message: message,
	})
}

func NotFound(c *gin.Context, message string) {
	respond(c, http.StatusNotFound, APIResponse{
		Status:  "error",
		Message: message,
	})
}

// SERVER ERROR RESPONSES

func InternalServerError(c *gin.Context, err any, message ...string) {
	msg := "Internal server error"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	respond(c, http.StatusInternalServerError, APIResponse{
		Status:  "error",
		Message: msg,
		Error:   err,
	})
}

func ServiceUnavailable(c *gin.Context, message string) {
	respond(c, http.StatusServiceUnavailable, APIResponse{
		Status:  "error",
		Message: message,
	})
}
