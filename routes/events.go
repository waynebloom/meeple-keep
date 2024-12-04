package routes

import (
	"golearn/first-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Path parameter 'id' has an invalid value."})
		return
	}

	event, err := model.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No event exists with that id."})
		return
	}

	c.JSON(http.StatusOK, event)
}

func getEvents(c *gin.Context) {
	events, err := model.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func postEvent(c *gin.Context) {
	var event model.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	event.UserID = c.GetInt64("userId")
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "the event was successfully created", "event": event})
}

func putEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Path parameter 'id' has an invalid value."})
		return
	}

	userId := c.GetInt64("userId")
	event, err := model.GetEvent(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are attempting to modify an event you do not own."})
		return
	}

	var updatedEvent model.Event
	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request body is malformed."})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request body is malformed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Path parameter 'id' has an invalid value."})
		return
	}

	userId := c.GetInt64("userId")
	event, err := model.GetEvent(id)

	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are attempting to delete an event you do not own."})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
		return
	}

	err = event.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}
