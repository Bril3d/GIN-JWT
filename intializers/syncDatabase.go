package intializers

import "github.com/bril3d/gin-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}