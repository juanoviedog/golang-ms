package initializers

import "github.com/juanoviedog/golang-ms/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
