package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthorizeAdmin проверяет, является ли пользователь администратором
func AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем роль из контекста (должна быть установлена middleware.Authenticate())
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Роль пользователя не найдена"})
			c.Abort()
			return
		}

		// Проверяем, является ли пользователь администратором
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ разрешен только администраторам"})
			c.Abort()
			return
		}

		c.Next()
	}
}
