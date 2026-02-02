package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go-project.com/go-project/models"
	"go-project.com/go-project/utils"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		utils.InternalServerError(context, err)
		return
	}
	utils.OK(context, "Events fetched successfully!", events)
}

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		utils.InternalServerError(context, err)
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil && err.Error() == "sql: no rows in result set" {
		utils.NotFound(context, "Event not found by id")
		return
	}

	if err != nil {
		utils.InternalServerError(context, err)
		return
	}

	utils.OK(context, "Event fetched successfully", event)
}

func CreateEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		utils.Unauthorized(context, "Unauthorized")
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		utils.Unauthorized(context, "Unauthorized")
		return
	}

	var event *models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		utils.BadRequest(context, "Could not parse request body", err)
		return
	}

	event.UserID = userId

	err = event.Save()

	if err != nil {
		utils.InternalServerError(context, err, "Failed to create Event")
		return
	}

	utils.Created(context, "Event created successfully", event)
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		utils.InternalServerError(context, err)
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil && err.Error() == "sql: no rows in result set" {
		utils.NotFound(context, "Event not found by id")
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		utils.BadRequest(context, "Could not parse request body", err)
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
		utils.InternalServerError(context, err, "Failed to update Event")
		return
	}

	utils.OK(context, "Event updated successfully", nil)
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		utils.InternalServerError(context, err)
	}

	event, err := models.GetEventByID(eventId)
	if err != nil && err.Error() == "sql: no rows in result set" {
		utils.NotFound(context, "Event not found by id")
		return
	}

	err = event.Delete()

	if err != nil {
		utils.InternalServerError(context, err, "Failed to delete Event")
		return
	}

	utils.OK(context, "Event deleted successfully", nil)
}
