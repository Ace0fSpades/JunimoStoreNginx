package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/interfaces/dto"
	"uniStore/Backend/internal/utils"
)

// UserServiceImpl implements the UserService interface
type UserServiceImpl struct {
	userRepo     models.UserRepository
	cartRepo     models.CartRepository
	favoriteRepo models.FavoriteRepository
	libraryRepo  models.LibraryRepository
	authUtils    *utils.AuthUtils
}

// NewUserService creates a new instance of UserService
func NewUserService(
	userRepo models.UserRepository,
	cartRepo models.CartRepository,
	favoriteRepo models.FavoriteRepository,
	libraryRepo models.LibraryRepository,
	authUtils *utils.AuthUtils,
) UserService {
	return &UserServiceImpl{
		userRepo:     userRepo,
		cartRepo:     cartRepo,
		favoriteRepo: favoriteRepo,
		libraryRepo:  libraryRepo,
		authUtils:    authUtils,
	}
}

// Register registers a new user
func (s *UserServiceImpl) Register(userDTO *dto.UserSignupDTO) (*dto.UserResponseDTO, error) {
	// Check if nickname already exists
	existingUser, err := s.userRepo.FindByNickname(userDTO.Nickname)
	if err == nil && existingUser != nil {
		return nil, errors.New("nickname already in use")
	}

	// Check if email already exists
	existingUser, err = s.userRepo.FindByEmail(userDTO.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already in use")
	}

	// Convert DTO to model
	user := userDTO.ToModel()

	// Set timestamps
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Hash the password
	hashedPassword, err := s.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	// Create user
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Initialize related entities
	if err := s.initializeUserData(user.ID); err != nil {
		return nil, err
	}

	// Convert model to DTO for response
	return dto.UserResponseDTOFromModel(user), nil
}

// initializeUserData creates empty cart, favorites and library for a new user
func (s *UserServiceImpl) initializeUserData(userID int) error {
	// Create empty shopping cart
	if err := s.cartRepo.Create(&models.ShoppingCart{
		UserID:    userID,
		CartItems: []*models.CartItem{},
	}); err != nil {
		return err
	}

	// Create empty favorites
	if err := s.favoriteRepo.Create(&models.Favorite{
		UserID:        userID,
		FavoriteItems: []*models.FavoriteItem{},
	}); err != nil {
		return err
	}

	// Create empty library
	if err := s.libraryRepo.Create(&models.Library{
		UserID:       userID,
		LibraryItems: []*models.LibraryItem{},
	}); err != nil {
		return err
	}

	return nil
}

// Login authenticates a user
func (s *UserServiceImpl) Login(loginDTO *dto.UserLoginDTO) (*dto.AuthResponseDTO, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(loginDTO.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email or password is incorrect")
		}
		return nil, err
	}

	// Verify password
	if !s.VerifyPassword(loginDTO.Password, user.Password) {
		return nil, errors.New("email or password is incorrect")
	}

	// Check for nil before accessing Role field
	var roleType string
	if user.Role != nil {
		roleType = user.Role.Type
	} else {
		// Get role from database if it wasn't loaded
		role, err := s.userRepo.FindByID(user.ID)
		if err == nil && role != nil && role.Role != nil {
			roleType = role.Role.Type
		} else {
			// Use default role value if unable to retrieve
			roleType = "user"
		}
	}

	// Generate tokens
	token, refreshToken, err := s.GenerateTokens(user.Email, user.Nickname, roleType, user.ID)
	if err != nil {
		return nil, err
	}

	// Update tokens in database
	user.Token = token
	user.RefreshToken = refreshToken
	user.UpdatedAt = time.Now()
	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	// Create response DTO
	authResponseDTO := dto.AuthResponseDTOFromModel(user)
	return authResponseDTO, nil
}

// GetUserByID gets a user by ID
func (s *UserServiceImpl) GetUserByID(id int) (*dto.UserResponseDTO, error) {
	// Get user from repository
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert model to DTO for response
	return dto.UserResponseDTOFromModel(user), nil
}

// GetAllUsers gets all users with pagination
func (s *UserServiceImpl) GetAllUsers(limit, offset int) ([]*dto.UserResponseDTO, error) {
	// Get users from repository
	users, err := s.userRepo.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert models to DTOs for response
	userDTOs := make([]*dto.UserResponseDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.UserResponseDTOFromModel(user)
	}

	return userDTOs, nil
}

// UpdateUser updates a user
func (s *UserServiceImpl) UpdateUser(id int, userDTO *dto.UserUpdateDTO) (*dto.UserResponseDTO, error) {
	// Get existing user
	existingUser, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert DTO to model
	updateData := userDTO.ToUpdateModel(id)

	// Update fields if provided
	if updateData.Nickname != "" {
		existingUser.Nickname = updateData.Nickname
	}
	if updateData.Email != "" {
		existingUser.Email = updateData.Email
	}
	if updateData.Password != "" {
		hashedPassword, err := s.HashPassword(updateData.Password)
		if err != nil {
			return nil, err
		}
		existingUser.Password = hashedPassword
	}
	if updateData.RoleID != 0 {
		existingUser.RoleID = updateData.RoleID
	}

	existingUser.UpdatedAt = time.Now()

	// Update user in repository
	if err := s.userRepo.Update(existingUser); err != nil {
		return nil, err
	}

	// Convert updated model to DTO for response
	return dto.UserResponseDTOFromModel(existingUser), nil
}

// AddPoints adds points to a user's account
func (s *UserServiceImpl) AddPoints(userID int, points int) (*dto.UserResponseDTO, error) {
	// Get user from repository
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	// Update points
	user.Points += points
	user.UpdatedAt = time.Now()

	// Update user in repository
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	// Convert updated model to DTO for response
	return dto.UserResponseDTOFromModel(user), nil
}

// VerifyPassword verifies a password against a hash
func (s *UserServiceImpl) VerifyPassword(password, hashedPassword string) bool {
	return utils.CheckPassword(password, hashedPassword)
}

// HashPassword hashes a password
func (s *UserServiceImpl) HashPassword(password string) (string, error) {
	return utils.HashPassword(password)
}

// GenerateTokens generates JWT tokens
func (s *UserServiceImpl) GenerateTokens(email, nickname, role string, id int) (string, string, error) {
	return s.authUtils.GenerateToken(email, nickname, role, id)
}

// GetUserIDFromToken extracts user ID from token
func (s *UserServiceImpl) GetUserIDFromToken(tokenString string) (int, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Get the secret key for validation
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return 0, err
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract user_id from claims
		if userID, ok := claims["user_id"].(float64); ok {
			return int(userID), nil
		}
		return 0, errors.New("user_id not found in token")
	}

	return 0, errors.New("invalid token")
}

// RefreshToken refreshes a user's token
func (s *UserServiceImpl) RefreshToken(refreshToken string) (*dto.AuthResponseDTO, error) {
	// Refresh token
	token, newRefreshToken, err := s.authUtils.RefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// Get user ID from token
	userID, err := s.GetUserIDFromToken(token)
	if err != nil {
		return nil, err
	}

	// Update tokens in database
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	user.Token = token
	user.RefreshToken = newRefreshToken
	user.UpdatedAt = time.Now()

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	// Create response DTO
	return dto.AuthResponseDTOFromModel(user), nil
}
