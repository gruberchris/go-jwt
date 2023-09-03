package initializers

import "go-jwt/models"

func SyncDatabase() {
	// Sync database here
	// auto migration
	DB.AutoMigrate(&models.User{})
}
