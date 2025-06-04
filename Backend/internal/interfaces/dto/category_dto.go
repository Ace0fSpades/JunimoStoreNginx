package dto

import (
	"uniStore/Backend/internal/domain/models"
)

// CategoryCreateDTO represents data needed for creating a new category
type CategoryCreateDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// CategoryUpdateDTO represents data needed for updating a category
type CategoryUpdateDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ToModel converts CategoryCreateDTO to Category model
func (dto *CategoryCreateDTO) ToModel() *models.Category {
	return &models.Category{
		Name:        dto.Name,
		Description: dto.Description,
	}
}

// ToUpdateModel converts CategoryUpdateDTO to Category model for updates
func (dto *CategoryUpdateDTO) ToUpdateModel(id int) *models.Category {
	return &models.Category{
		ID:          id,
		Name:        dto.Name,
		Description: dto.Description,
	}
}
