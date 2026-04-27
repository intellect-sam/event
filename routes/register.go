package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/intellect-sam/event/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for event"})
}

func unregisterFromEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.Unregister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to unregister from event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully unregistered from event"})

}
