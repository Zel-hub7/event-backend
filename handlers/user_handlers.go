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
	if err := config.DB.Find(&users).Error; err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser handles the HTTP POST request to add a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert the new user into the database
	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
