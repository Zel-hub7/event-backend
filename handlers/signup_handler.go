package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Zel-hub7/event-backend/event-management-system/config"
	"github.com/Zel-hub7/event-backend/event-management-system/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Name:     creds.Name,
		Email:    creds.Email,
		Password: string(hashedPassword),
	}

	// Save the user in the database
	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}
