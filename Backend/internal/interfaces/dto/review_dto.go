package dto

import (
	"time"

	"uniStore/Backend/internal/domain/models"
)

// ReviewCreateDTO represents data needed for creating a new review
type ReviewCreateDTO struct {
	UserID  int    `json:"user_id" binding:"required"`
	GameID  int    `json:"game_id" binding:"required"`
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment"`
}

// ReviewUpdateDTO represents data needed for updating a review
type ReviewUpdateDTO struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment"`
}

// ReviewResponseDTO represents a review for API responses
type ReviewResponseDTO struct {
	ID        int              `json:"id"`
	User      *UserResponseDTO `json:"user,omitempty"`
	Game      *GameDTO         `json:"game,omitempty"`
	Rating    int              `json:"rating"`
	Comment   string           `json:"comment"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

// ToModel converts ReviewCreateDTO to Review model
func (dto *ReviewCreateDTO) ToModel() *models.Review {
	return &models.Review{
		UserID:      dto.UserID,
		GameID:      dto.GameID,
		Rating:      dto.Rating,
		Description: dto.Comment,
		Title:       "Review", // Default title
	}
}

// ToUpdateModel converts ReviewUpdateDTO to Review model for updates
func (dto *ReviewUpdateDTO) ToUpdateModel(id int) *models.Review {
	return &models.Review{
		ID:          id,
		Rating:      dto.Rating,
		Description: dto.Comment,
	}
}

// ReviewResponseDTOFromModel converts Review model to ReviewResponseDTO
func ReviewResponseDTOFromModel(review *models.Review) *ReviewResponseDTO {
	dto := &ReviewResponseDTO{
		ID:        review.ID,
		Rating:    review.Rating,
		Comment:   review.Description,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}

	// Add full user data if available
	if review.User != nil {
		dto.User = UserResponseDTOFromModel(review.User)
	}

	// Add full game data if available
	if review.Game != nil {
		dto.Game = GameDTOFromModel(review.Game)
	}

	return dto
}

// ReviewResponseDTOsFromModels converts a slice of Review models to a slice of ReviewResponseDTOs
func ReviewResponseDTOsFromModels(reviews []*models.Review) []*ReviewResponseDTO {
	dtos := make([]*ReviewResponseDTO, len(reviews))
	for i, review := range reviews {
		dtos[i] = ReviewResponseDTOFromModel(review)
	}
	return dtos
}
