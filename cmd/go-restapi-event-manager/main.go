package main

import (
	"fmt"
	"github.com/sJones1997/go-restapi-event-manager/internal/events"
)

const port = ":8080"

func main() {
	server := events.NewServer()
	err := server.Run(port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server running on port %s\n", port)
}
