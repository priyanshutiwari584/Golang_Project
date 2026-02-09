package controllers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go-project.com/go-project/models"
	"go-project.com/go-project/utils"
)

func RegisterForEvent(context *gin.Context) {
	var registrations models.Registration

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		utils.BadRequest(context, "Could not parse request body", err)
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			utils.NotFound(context, "Event not found by id")
			return
		}

		utils.InternalServerError(context, err, "Could not get event")
		return
	}

	registrations.EventID = eventId
	registrations.UserID = userId
	err = registrations.Save()

	if err != nil {
		utils.InternalServerError(context, err, "Could not register for event")
		return
	}

	utils.OK(context, "User registered for event successfully", event)
}

func GetAllRegisteredEvents(context *gin.Context) {
	registrations, err := models.GetAllRegistrations()

	if err != nil {
		utils.InternalServerError(context, err, "Could not get Events")
		return
	}

	utils.OK(context, "Registration fetched successfully!", registrations)
}

func CancellationForEvent(context *gin.Context) {
	var registrations models.Registration

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		utils.BadRequest(context, "Could not parse request body", err)
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			utils.NotFound(context, "Event not found by id")
			return
		}

		utils.InternalServerError(context, err, "Could not get event")
		return
	}

	registrations.EventID = eventId
	registrations.UserID = userId
	err = registrations.Delete()

	if err != nil {
		utils.InternalServerError(context, err, "Could not Delete registration for event")
		return
	}

	utils.OK(context, "User cancelled for event successfully", event)
}
