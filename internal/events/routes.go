package events

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sJones1997/go-restapi-event-manager/internal/utils/HTTPError"
	"net/http"
	"strconv"
)

func NewServer() *gin.Engine {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)
	server.POST("/events", createEvent)
	server.DELETE("/event/:id", deleteEvent)
	server.PUT("/event/:id", updateEvent)

	return server
}

func getEvents(c *gin.Context) {
	events, err := GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	event, err := GetEvent(id)

	var httpErr *HTTPError.HTTPError
	if errors.As(err, &httpErr) {
		c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	event := Event{}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.UserID = 1

	event, err := event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func deleteEvent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	err = DeleteEvent(id)

	var httpErr *HTTPError.HTTPError
	if errors.As(err, &httpErr) {
		c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted!"})

}

func updateEvent(c *gin.Context) {

	event := Event{}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := UpdateEvent(event)

	var httpErr *HTTPError.HTTPError
	if errors.As(err, &httpErr) {
		c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event updated", "event": event})

}
