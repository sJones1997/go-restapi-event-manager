package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sJones1997/go-restapi-event-manager/internal/models"
	"net/http"
)

func NewServer() *gin.Engine {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	return server
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	event := models.Event{}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1

	c.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}