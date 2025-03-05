package middleware

import (
	"jwt-auth/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует"})
			c.Abort()
			return
		}

		// Парсим и проверяем токен
		claims, err := services.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный или просроченный токен"})
			c.Abort()
			return
		}

		// Проверяем, что claims не nil
		if claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка обработки токена"})
			c.Abort()
			return
		}

		// Сохраняем имя пользователя в контексте запроса
		c.Set("username", claims.Username)
		c.Next()
	}
}
