package dto

import (
	"time"

	"uniStore/Backend/internal/domain/models"
)

// UserLoginDTO represents data needed for user login
type UserLoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserSignupDTO represents data needed for user registration
type UserSignupDTO struct {
	Nickname        string `json:"nickname" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

// UserUpdateDTO represents data needed for user update
type UserUpdateDTO struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}

// UserResponseDTO represents a user for API responses (no sensitive data)
type UserResponseDTO struct {
	ID        int       `json:"id"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Role      *RoleDTO  `json:"role,omitempty"`
	Points    int       `json:"points"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuthResponseDTO represents the response for login/signup operations
type AuthResponseDTO struct {
	User         UserResponseDTO `json:"user"`
	Token        string          `json:"token"`
	RefreshToken string          `json:"refresh_token"`
}

// RoleDTO represents a user role
type RoleDTO struct {
	ID          int       `json:"id"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ToModel converts UserSignupDTO to User model
func (dto *UserSignupDTO) ToModel() *models.User {
	return &models.User{
		Nickname: dto.Nickname,
		Email:    dto.Email,
		Password: dto.Password, // Password will be hashed in service layer
		RoleID:   2,            // Default role ID (user)
	}
}

// ToModel converts UserUpdateDTO to User model for updates
func (dto *UserUpdateDTO) ToUpdateModel(id int) *models.User {
	return &models.User{
		ID:       id,
		Nickname: dto.Nickname,
		Email:    dto.Email,
		Password: dto.Password, // Password will be hashed in service layer
		RoleID:   dto.RoleID,
	}
}

// UserResponseDTOFromModel converts User model to UserResponseDTO
func UserResponseDTOFromModel(user *models.User) *UserResponseDTO {
	dto := &UserResponseDTO{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Points:    user.Points,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	// Add full role data if role is loaded
	if user.Role != nil {
		dto.Role = RoleDTOFromModel(user.Role)
	}

	return dto
}

// AuthResponseDTOFromModel creates AuthResponseDTO from User model
func AuthResponseDTOFromModel(user *models.User) *AuthResponseDTO {
	return &AuthResponseDTO{
		User:         *UserResponseDTOFromModel(user),
		Token:        user.Token,
		RefreshToken: user.RefreshToken,
	}
}

// RoleDTOFromModel converts Role model to RoleDTO
func RoleDTOFromModel(role *models.Role) *RoleDTO {
	return &RoleDTO{
		ID:          role.ID,
		Type:        role.Type,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}

// RoleDTOsFromModels converts a slice of Role models to a slice of RoleDTOs
func RoleDTOsFromModels(roles []*models.Role) []*RoleDTO {
	dtos := make([]*RoleDTO, len(roles))
	for i, role := range roles {
		dtos[i] = RoleDTOFromModel(role)
	}
	return dtos
}
