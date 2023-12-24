package initializers

import (
	"API/database"
	"API/models"
)

func SyncDatabase() {
	database.DB.AutoMigrate(&models.User{})
}
