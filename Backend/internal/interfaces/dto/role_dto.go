package dto

import (
	"uniStore/Backend/internal/domain/models"
)

// RoleCreateDTO represents data needed for creating a new role
type RoleCreateDTO struct {
	Type        string `json:"type" binding:"required"`
	Description string `json:"description"`
}

// RoleUpdateDTO represents data needed for updating a role
type RoleUpdateDTO struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// ToModel converts RoleCreateDTO to Role model
func (dto *RoleCreateDTO) ToModel() *models.Role {
	return &models.Role{
		Type:        dto.Type,
		Description: dto.Description,
	}
}

// ToUpdateModel converts RoleUpdateDTO to Role model for updates
func (dto *RoleUpdateDTO) ToUpdateModel(id int) *models.Role {
	return &models.Role{
		ID:          id,
		Type:        dto.Type,
		Description: dto.Description,
	}
}
