package routes

import (
	"net/http"
	"strconv"

	"example.org/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}
	event, err := models.GetEventById(eventId)
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "you are not allowed to update this event"})
		return
	}

	err = event.CreateRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "can not create a registration"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "registration created"})
}
func cancleRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}
	var event models.Event
	event.Id = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "can not cancel the registration"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "registration cancelled"})
}
