package services

import (
	"errors"
	"fmt"
	"jwt-auth/config"
	"jwt-auth/internal/database"
	"jwt-auth/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWT-token generator
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JwtSecret))
}

// Check token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})
	// Проверяем, что токен валиден
	if err != nil || !token.Valid {
		return nil, err
	}

	// Дополнительная проверка (чтобы claims не было nil)
	if claims.Username == "" {
		return nil, errors.New("некорректные claims")
	}

	return claims, nil
}

// Регистрация нового пользователя
func RegisterUser(username, password string) error {
	var existingUser models.User
	result := database.DB.Where("username = ?", username).First(&existingUser)

	// Если пользователь уже существует — ошибка
	if result.RowsAffected > 0 {
		return errors.New("пользователь уже существует")
	}

	// Создаём нового пользователя
	newUser := models.User{Username: username, Password: password}
	if err := newUser.HashPassword(); err != nil {
		return err
	}

	// Сохраняем пользователя в PostgreSQL
	if err := database.DB.Create(&newUser).Error; err != nil {
		fmt.Println("❌ Ошибка сохранения в базу:", err)
		return err
	}

	fmt.Println("✅ Пользователь успешно сохранён в базе:", newUser.Username)
	return nil
}

func AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)

	if result.RowsAffected == 0 {
		return nil, errors.New("пользователь не найден")
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("неверный пароль")
	}

	return &user, nil
}
