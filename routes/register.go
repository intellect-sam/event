package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/intellect-sam/event/models"
)

// @Summary      Register for an event
// @Description  Registers the authenticated user for a specific event
// @Tags         Events
// @Produce      json
// @Param        id path int true "Event ID"
// @Success      201 "Successfully registered for event"
// @Failure      400  "Invalid event ID"
// @Failure      404  "Event not found"
// @Failure      500  "Failed to register for event"
// @Security     BearerAuth
// @Router       /events/{id}/register [post]
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

// @Summary      Unregister from an event
// @Description  Removes the authenticated user's registration from a specific event
// @Tags         Events
// @Produce      json
// @Param        id path int true "Event ID"
// @Success      200 {object} map[string]string "Successfully unregistered from event"
// @Failure      400 {object} map[string]string "Invalid event ID"
// @Failure      500 {object} map[string]string "Failed to unregister from event"
// @Security     BearerAuth
// @Router       /events/{id}/unregister [delete]
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
