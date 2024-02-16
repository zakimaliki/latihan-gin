package helpers

import (
	"latihan-gin/src/config"
	"latihan-gin/src/models"
)

func Migration() {
	config.DB.AutoMigrate(&models.Month{})
}
