package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// AuthUtils provides utilities for authentication
type AuthUtils struct {
	secretKey string
}

// NewAuthUtils creates a new AuthUtils instance
func NewAuthUtils() *AuthUtils {
	return &AuthUtils{
		secretKey: "your-secret-key", // In production, this should be loaded from an environment variable
	}
}

// GenerateToken generates a JWT token for a user
func (a *AuthUtils) GenerateToken(email, nickname, role string, id int) (string, string, error) {
	// Generate access token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["nickname"] = nickname
	claims["role"] = role
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24 hours expiration

	accessToken, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["email"] = email
	refreshClaims["user_id"] = id
	refreshClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // 7 days expiration for refresh token

	refreshTokenString, err := refreshToken.SignedString([]byte(a.secretKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshTokenString, nil
}

// VerifyToken verifies a JWT token
func (a *AuthUtils) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.secretKey), nil
	})

	return token, err
}

// RefreshToken creates a new token if the current token is still valid
func (a *AuthUtils) RefreshToken(tokenString string) (string, string, error) {
	// Verify the token
	token, err := a.VerifyToken(tokenString)
	if err != nil {
		return "", "", err
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("failed to extract claims")
	}

	// Extract user information
	email, ok := claims["email"].(string)
	if !ok {
		return "", "", errors.New("failed to extract email")
	}

	nickname, ok := claims["nickname"].(string)
	if !ok {
		nickname = "" // Allow for backward compatibility
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", "", errors.New("failed to extract role")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return "", "", errors.New("failed to extract user ID")
	}

	// Generate new tokens
	return a.GenerateToken(email, nickname, role, int(userID))
}

// IsProd checks if the application is running in production mode
func IsProd() bool {
	return os.Getenv("APP_ENV") == "production"
}
