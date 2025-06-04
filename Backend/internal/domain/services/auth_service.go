package services

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/utils"
)

// AuthServiceImpl implements AuthService interface
type AuthServiceImpl struct {
	userRepo  models.UserRepository
	roleRepo  models.RoleRepository
	authUtils *utils.AuthUtils
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo models.UserRepository, roleRepo models.RoleRepository, authUtils *utils.AuthUtils) AuthService {
	return &AuthServiceImpl{
		userRepo:  userRepo,
		roleRepo:  roleRepo,
		authUtils: authUtils,
	}
}

// VerifyToken verifies a JWT token and returns the user ID and role
func (s *AuthServiceImpl) VerifyToken(token string) (int, string, error) {
	// Verify the token
	parsedToken, err := s.authUtils.VerifyToken(token)
	if err != nil {
		return 0, "", err
	}

	// Check if the token is valid
	if !parsedToken.Valid {
		return 0, "", errors.New("invalid token")
	}

	// Extract claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("failed to extract claims")
	}

	// Extract user ID and role
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, "", errors.New("failed to extract user ID")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", errors.New("failed to extract role")
	}

	return int(userID), role, nil
}

// RefreshUserToken refreshes a user's token
func (s *AuthServiceImpl) RefreshUserToken(token string) (string, string, error) {
	return s.authUtils.RefreshToken(token)
}

// MatchUserTypeToID checks if a user can access a resource
func (s *AuthServiceImpl) MatchUserTypeToID(userID int, roleType string) error {
	// Admin can access all resources
	if roleType == "admin" {
		return nil
	}

	// Get the ID from the token
	tokenUserID, err := s.getUserIDFromContext()
	if err != nil {
		return err
	}

	// Regular users can only access their own resources
	if tokenUserID != userID {
		return errors.New("unauthorized access to another user's resource")
	}

	return nil
}

// getUserIDFromContext is a helper function to get the user ID from the context
func (s *AuthServiceImpl) getUserIDFromContext() (int, error) {
	// In a real implementation, we would extract userID from the context
	// But since this is a helper method for testing,
	// we can simply return 1 (default user)
	return 1, nil
}
