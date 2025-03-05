package database

import (
	"fmt"
	"jwt-auth/internal/models"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to database and migration
func InitDB() {
	viper.AutomaticEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Ошибка подключения к базе данных:", err)
	}

	fmt.Println("✅ Подключение к PostgreSQL успешно!")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ Ошибка миграции:", err)
	}

	fmt.Println("✅ Таблица users инициализирована!")
}
