package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Zel-hub7/event-backend/event-management-system/config"
	"github.com/Zel-hub7/event-backend/event-management-system/models"
)

// GetEvents handles the HTTP request to fetch events
func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	if err := config.DB.Find(&events).Error; err != nil {
		http.Error(w, "Error fetching events", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// CreateEvent handles the HTTP POST request to add a new event
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := config.DB.Create(&event).Error; err != nil {
		http.Error(w, "Error creating event", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}
