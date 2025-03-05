package database

import (
	"fmt"
	"jwt-auth/internal/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to database and migration
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	err = DB.AutoMigrate(&models.Users{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	fmt.Println("✅ База данных успешно инициализирована!")
}
