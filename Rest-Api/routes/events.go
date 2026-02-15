package routes

import (
	"net/http"
	"strconv"

	"example.org/rest-api/models"
	"github.com/gin-gonic/gin"
)

func deleteById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "error fetching data", "error": err})
		return
	}

	userId := context.GetInt64("userId")
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "you are not allowed to delete this event"})
		return
	}

	err = event.DeleteById()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the data", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}
	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching data", "error": err})
		return
	}

	userId := context.GetInt64("userId")
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "you are not allowed to update this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	updatedEvent.Id = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event is updated"})
}

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching data", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	event.UserId = context.GetInt64("userId")
	createdEvent, err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error while creating data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": createdEvent})
}

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error while getting data"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": events})
}
