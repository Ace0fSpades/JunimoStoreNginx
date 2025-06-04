package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"uniStore/Backend/internal/domain/services"
	"uniStore/Backend/internal/utils"
)

// AuthMiddleware handles authentication middleware
type AuthMiddleware struct {
	authService services.AuthService
}

// NewAuthMiddleware creates a new auth middleware
func NewAuthMiddleware(authService services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// Authenticate is a middleware for authenticating requests
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize auth utils
		authUtils := utils.NewAuthUtils()

		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		// The format should be "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer <token>"})
			c.Abort()
			return
		}

		// Verify the token
		token, err := authUtils.VerifyToken(parts[1])
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to extract claims"})
			c.Abort()
			return
		}

		// Extract user ID and role
		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to extract user ID"})
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to extract role"})
			c.Abort()
			return
		}

		// Set user ID and role in the context
		c.Set("userID", int(userID))
		c.Set("role", role)

		c.Next()
	}
}
