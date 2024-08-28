// config/config.go
package config

import (
	"log"

	"github.com/Zel-hub7/event-backend/event-management-system/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=1234 dbname=eventdb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Migrate the schema
	DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Blacklist{})
}
