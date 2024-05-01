package main

import (
	"fmt"
	"github.com/sJones1997/go-restapi-event-manager/internal/routes"
)

const port = ":8080"

func main() {
	server := routes.NewServer()
	err := server.Run(port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server running on port %s\n", port)
}
