package config

import (
	"log"

	"github.com/spf13/viper"
)

var JwtSecret string
var ServerPort string

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Нет .env файла, используются перменные окружения")
	}

	JwtSecret = viper.GetString("JWT_SECRET")
	ServerPort = viper.GetString("SERVER_PORT")

	if JwtSecret == "" {
		log.Fatal("JWT_SECRET не задан")
	}
}
