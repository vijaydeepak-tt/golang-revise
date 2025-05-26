package routes

import (
	"net/http"
	"strconv"

	"example.com/go_events_api/pkgs/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not able to parse the ID"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not able to fetch the event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not able to register the event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event is Registered!!!"})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not able to parse the ID"})
		return
	}

	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not able to cancel the registered event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event is Cancelled!!!"})

}
