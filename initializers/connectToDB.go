package initializers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// Connect to SQLite DB
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	// Catch errors
	if err != nil {
		panic("Failed to connect to database service")

	}
	DB = database
}
