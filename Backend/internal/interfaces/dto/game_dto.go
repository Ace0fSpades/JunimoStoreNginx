package dto

import (
	"encoding/base64"
	"time"

	"uniStore/Backend/internal/domain/models"
)

// GameCreateDTO represents data needed for creating a new game
type GameCreateDTO struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Price       float64   `json:"price" binding:"required,gte=0"`
	ReleaseDate time.Time `json:"release_date"`
	DeveloperID int       `json:"developer_id" binding:"required"`
	CategoryID  int       `json:"category_id" binding:"required"`
	ImageData   string    `json:"image_data"`
	ImageName   string    `json:"image_name"`
}

// GameUpdateDTO represents data needed for updating a game
type GameUpdateDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price" binding:"omitempty,gte=0"`
	ReleaseDate time.Time `json:"release_date"`
	DeveloperID int       `json:"developer_id"`
	CategoryID  int       `json:"category_id"`
	ImageData   string    `json:"image_data"`
	ImageName   string    `json:"image_name"`
}

// GameSearchDTO represents search criteria for games
type GameSearchDTO struct {
	Title      string `form:"title"`
	CategoryID int    `form:"category_id"`
	Limit      int    `form:"limit,default=10"`
	Offset     int    `form:"offset,default=0"`
}

// GameDTO represents full game data for API response
type GameDTO struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	ReleaseDate time.Time      `json:"release_date"`
	Developer   *DeveloperDTO  `json:"developer,omitempty"`
	Category    *CategoryDTO   `json:"category,omitempty"`
	ImageData   string         `json:"image_data"` // Base64-encoded image data
	ImageName   string         `json:"image_name"`
	Restricts   []*RestrictDTO `json:"restricts,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// DeveloperDTO represents developer data for API response
type DeveloperDTO struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	Description string    `json:"description"`
	WebsiteURL  string    `json:"website_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CategoryDTO represents category data for API response
type CategoryDTO struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ToModel converts GameCreateDTO to Game model
func (dto *GameCreateDTO) ToModel() *models.Game {
	// Process image data if needed
	var imageData []byte
	if dto.ImageData != "" {
		imageData = []byte(dto.ImageData)
	}

	return &models.Game{
		Title:       dto.Title,
		Description: dto.Description,
		Price:       dto.Price,
		ReleaseDate: dto.ReleaseDate,
		DeveloperID: dto.DeveloperID,
		CategoryID:  dto.CategoryID,
		ImageData:   imageData,
		ImageName:   dto.ImageName,
	}
}

// ToUpdateModel converts GameUpdateDTO to Game model
func (dto *GameUpdateDTO) ToUpdateModel(id int) *models.Game {
	// Process image data if needed
	var imageData []byte
	if dto.ImageData != "" {
		imageData = []byte(dto.ImageData)
	}

	return &models.Game{
		ID:          id,
		Title:       dto.Title,
		Description: dto.Description,
		Price:       dto.Price,
		ReleaseDate: dto.ReleaseDate,
		DeveloperID: dto.DeveloperID,
		CategoryID:  dto.CategoryID,
		ImageData:   imageData,
		ImageName:   dto.ImageName,
	}
}

// FromModel converts Game model to GameDTO
func GameDTOFromModel(model *models.Game) *GameDTO {
	var imageDataStr string
	if model.ImageData != nil && len(model.ImageData) > 0 {
		imageDataStr = base64.StdEncoding.EncodeToString(model.ImageData)
	}

	dto := &GameDTO{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		Price:       model.Price,
		ReleaseDate: model.ReleaseDate,
		ImageData:   imageDataStr,
		ImageName:   model.ImageName,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

	// Add developer data if available
	if model.Developer != nil {
		dto.Developer = DeveloperDTOFromModel(model.Developer)
	}

	// Add category data if available
	if model.Category != nil {
		dto.Category = CategoryDTOFromModel(model.Category)
	}

	// Add restrictions if available
	if model.Restricts != nil && len(model.Restricts) > 0 {
		dto.Restricts = RestrictDTOsFromModels(model.Restricts)
	}

	return dto
}

// FromModels converts a slice of Game models to a slice of GameDTOs
func GameDTOsFromModels(models []*models.Game) []*GameDTO {
	dtos := make([]*GameDTO, len(models))
	for i, model := range models {
		dtos[i] = GameDTOFromModel(model)
	}
	return dtos
}

// FromModel converts Developer model to DeveloperDTO
func DeveloperDTOFromModel(model *models.Developer) *DeveloperDTO {
	return &DeveloperDTO{
		ID:          model.ID,
		Name:        model.Name,
		Country:     model.Country,
		Description: model.Description,
		WebsiteURL:  model.WebsiteURL,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

// FromModels converts a slice of Developer models to a slice of DeveloperDTOs
func DeveloperDTOsFromModels(models []*models.Developer) []*DeveloperDTO {
	dtos := make([]*DeveloperDTO, len(models))
	for i, model := range models {
		dtos[i] = DeveloperDTOFromModel(model)
	}
	return dtos
}

// FromModel converts Category model to CategoryDTO
func CategoryDTOFromModel(model *models.Category) *CategoryDTO {
	return &CategoryDTO{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

// FromModels converts a slice of Category models to a slice of CategoryDTOs
func CategoryDTOsFromModels(models []*models.Category) []*CategoryDTO {
	dtos := make([]*CategoryDTO, len(models))
	for i, model := range models {
		dtos[i] = CategoryDTOFromModel(model)
	}
	return dtos
}
