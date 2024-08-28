// models/blacklist.go
package models

import "gorm.io/gorm"

// Blacklist represents the structure of a blacklisted token.
type Blacklist struct {
	gorm.Model
	Token string `gorm:"not null;unique"`
}
