package dto

import (
	"time"

	"uniStore/Backend/internal/domain/models"
)

// FavoriteItemDTO represents a favorite item response
type FavoriteItemDTO struct {
	ID        int       `json:"id"`
	Game      *GameDTO  `json:"game,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// FavoriteResponseDTO represents a favorite response
type FavoriteResponseDTO struct {
	ID        int               `json:"id"`
	UserID    int               `json:"user_id"`
	User      *UserResponseDTO  `json:"user,omitempty"`
	Items     []FavoriteItemDTO `json:"items"`
	CreatedAt time.Time         `json:"created_at"`
}

// FavoriteItemDTOFromModel converts a FavoriteItem model and Game model to FavoriteItemDTO
func FavoriteItemDTOFromModel(favoriteItem *models.FavoriteItem, game *models.Game) *FavoriteItemDTO {
	dto := &FavoriteItemDTO{
		ID:        favoriteItem.ID,
		CreatedAt: favoriteItem.CreatedAt,
	}

	// Add full game data if available
	if game != nil {
		dto.Game = GameDTOFromModel(game)
	}

	return dto
}

// FavoriteResponseDTOFromModel converts a Favorite model and favorite items to FavoriteResponseDTO
func FavoriteResponseDTOFromModel(favorite *models.Favorite, favoriteItems []*models.FavoriteItem, favoriteItemDTOs []FavoriteItemDTO) *FavoriteResponseDTO {
	dto := &FavoriteResponseDTO{
		ID:        favorite.ID,
		UserID:    favorite.UserID,
		Items:     favoriteItemDTOs,
		CreatedAt: favorite.CreatedAt,
	}

	// Add user data if available
	if favorite.User != nil {
		dto.User = UserResponseDTOFromModel(favorite.User)
	}

	return dto
}
