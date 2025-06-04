package dto

import (
	"uniStore/Backend/internal/domain/models"
)

// DeveloperCreateDTO represents data needed for creating a new developer
type DeveloperCreateDTO struct {
	Name        string `json:"name" binding:"required"`
	Country     string `json:"country"`
	Description string `json:"description"`
	WebsiteURL  string `json:"website_url"`
}

// DeveloperUpdateDTO represents data needed for updating a developer
type DeveloperUpdateDTO struct {
	Name        string `json:"name"`
	Country     string `json:"country"`
	Description string `json:"description"`
	WebsiteURL  string `json:"website_url"`
}

// ToModel converts DeveloperCreateDTO to Developer model
func (dto *DeveloperCreateDTO) ToModel() *models.Developer {
	return &models.Developer{
		Name:        dto.Name,
		Country:     dto.Country,
		Description: dto.Description,
		WebsiteURL:  dto.WebsiteURL,
	}
}

// ToUpdateModel converts DeveloperUpdateDTO to Developer model for updates
func (dto *DeveloperUpdateDTO) ToUpdateModel(id int) *models.Developer {
	return &models.Developer{
		ID:          id,
		Name:        dto.Name,
		Country:     dto.Country,
		Description: dto.Description,
		WebsiteURL:  dto.WebsiteURL,
	}
}
