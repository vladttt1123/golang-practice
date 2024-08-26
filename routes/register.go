package routes

import (
	"eventBooking/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	// in path parameters need to pass the event it, and the user who is being registered
	// is being determined by the token
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration for event has been successful!"})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel the registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration has been cancelled successfully!"})

}
