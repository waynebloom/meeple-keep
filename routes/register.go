package routes

import (
	"golearn/first-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Path parameter 'id' has an invalid value."})
		return
	}

	event, err := model.GetEvent(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No event exists with that id."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User was registered for the event."})
}

func deregisterForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	var event model.Event
	event.ID = eventId

	err = event.Deregister(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Deregistering user for event failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User was deregistered from the event."})
}

func getEventRoster(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Path parameter 'id' has an invalid value."})
		return
	}

	event, err := model.GetEvent(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No event exists with that id."})
		return
	}

	if event.ID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User is not authorized to view this event's roster."})
		return
	}

	roster, err := model.GetRoster(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve a roster for this event."})
		return
	}

	c.JSON(http.StatusOK, roster)
}
