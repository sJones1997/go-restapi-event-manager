package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const port = ":8080"

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(port)
}

func getEvents(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})

}
