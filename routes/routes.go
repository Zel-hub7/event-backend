// routes/routes.go
package routes

import (
	"github.com/Zel-hub7/event-backend/event-management-system/handlers"
	"github.com/Zel-hub7/event-backend/event-management-system/middlewares"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/signup", handlers.Signup).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)
	api.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	api.HandleFunc("/events", handlers.GetEvents).Methods("GET")
	api.HandleFunc("/events", handlers.CreateEvent).Methods("POST")
	api.HandleFunc("/logout", handlers.Logout).Methods("POST") // Logout route

	return router
}
