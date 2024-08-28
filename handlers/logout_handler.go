// handlers/logout_handler.go
package handlers

import (
	"net/http"
	"strings"

	"github.com/Zel-hub7/event-backend/event-management-system/config"
	"github.com/Zel-hub7/event-backend/event-management-system/models"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	// Extract the token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Store the token in the blacklist
	blacklistedToken := models.Blacklist{Token: token}
	if err := config.DB.Create(&blacklistedToken).Error; err != nil {
		http.Error(w, "Failed to log out", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully logged out"))
}
