package main

import (
	"fmt"
	"log"
	"radiant/server"
)

func main() {
	s, err := server.NewServer(":8080") // Listen on port 8080.
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	fmt.Println("Server started...")

	// Prevent the main function from exiting.
	// In a more realistic application, you might instead start a
	// control loop here that can handle SIGINT and other signals
	// to gracefully shut down the server.
	select {}
}
