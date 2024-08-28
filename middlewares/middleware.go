// middlewares/auth_middleware.go
package middlewares

import (
	"net/http"
	"strings"

	"github.com/Zel-hub7/event-backend/event-management-system/config"
	"github.com/Zel-hub7/event-backend/event-management-system/models"
	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Check if the token is blacklisted
		var blacklistedToken models.Blacklist
		if err := config.DB.Where("token = ?", tokenString).First(&blacklistedToken).Error; err == nil {
			http.Error(w, "Token is blacklisted", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
