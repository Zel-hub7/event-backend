package routes

import (
	"github.com/Zel-hub7/event-backend/event-management-system/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/events", handlers.GetEvents).Methods("GET")
	return router
}
