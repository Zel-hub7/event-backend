package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Zel-hub7/event-backend/event-management-system/config"
	"github.com/Zel-hub7/event-backend/event-management-system/models"
)

// GetUsers handles the HTTP request to fetch users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	// Fetch users from the database
	if err := config.DB.Find(&users).Error; err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	// Encode users to JSON and write to response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetEvents handles the HTTP request to fetch events
func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	// Fetch events from the database
	if err := config.DB.Find(&events).Error; err != nil {
		http.Error(w, "Error fetching events", http.StatusInternalServerError)
		return
	}
	// Encode events to JSON and write to response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
