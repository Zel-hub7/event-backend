package routes

import (
	"net/http"

	"github.com/Zel-hub7/event-backend/event-management-system/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupRouter() http.Handler {
	router := mux.NewRouter()

	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow your frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With", "Accept"},
		AllowCredentials: true,
	})

	// Define your routes
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	api.HandleFunc("/events", handlers.GetEvents).Methods("GET")
	api.HandleFunc("/events", handlers.CreateEvent).Methods("POST")

	// Apply CORS middleware to the router
	return c.Handler(router)
}
