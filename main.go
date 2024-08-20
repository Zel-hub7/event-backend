package main

import (
	"log"
	"net/http"

	"github.com/Zel-hub7/event-backend/event-management-system/config"
	"github.com/Zel-hub7/event-backend/event-management-system/routes"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Set up the router
	router := routes.SetupRouter()

	// Start the server
	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
