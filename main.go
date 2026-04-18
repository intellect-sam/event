package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/intellect-sam/event/db"
	"github.com/intellect-sam/event/models"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id")
	server.POST("/events", createEvents)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event, try again later"})
	}
	context.JSON(http.StatusOK, events)
}

// func getEvent(context *gin.Context) {
// 	strconv context.Params("id")

// }

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the error"})
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event, try again later"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event created", "event": event})

}
