package dto

import (
	"time"

	"uniStore/Backend/internal/domain/models"
)

// RestrictDTO represents restriction data for API response
type RestrictDTO struct {
	ID        int       `json:"id"`
	GameID    int       `json:"game_id"`
	GameTitle string    `json:"game_title,omitempty"`
	Region    string    `json:"region"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RestrictCreateDTO represents data needed for creating a new restriction
type RestrictCreateDTO struct {
	GameID int    `json:"game_id" binding:"required"`
	Region string `json:"region" binding:"required"`
}

// RestrictUpdateDTO represents data needed for updating a restriction
type RestrictUpdateDTO struct {
	Region string `json:"region" binding:"required"`
}

// ToModel converts RestrictCreateDTO to Restrict model
func (dto *RestrictCreateDTO) ToModel() *models.Restrict {
	return &models.Restrict{
		GameID: dto.GameID,
		Region: dto.Region,
	}
}

// ToUpdateModel converts RestrictUpdateDTO to Restrict model for updates
func (dto *RestrictUpdateDTO) ToUpdateModel(id int) *models.Restrict {
	return &models.Restrict{
		ID:     id,
		Region: dto.Region,
	}
}

// RestrictDTOFromModel converts Restrict model to RestrictDTO
func RestrictDTOFromModel(restrict *models.Restrict) *RestrictDTO {
	dto := &RestrictDTO{
		ID:        restrict.ID,
		GameID:    restrict.GameID,
		Region:    restrict.Region,
		CreatedAt: restrict.CreatedAt,
		UpdatedAt: restrict.UpdatedAt,
	}

	// Add game title if available
	if restrict.Game != nil {
		dto.GameTitle = restrict.Game.Title
	}

	return dto
}

// RestrictDTOsFromModels converts a slice of Restrict models to a slice of RestrictDTOs
func RestrictDTOsFromModels(restricts []*models.Restrict) []*RestrictDTO {
	dtos := make([]*RestrictDTO, len(restricts))
	for i, restrict := range restricts {
		dtos[i] = RestrictDTOFromModel(restrict)
	}
	return dtos
}
