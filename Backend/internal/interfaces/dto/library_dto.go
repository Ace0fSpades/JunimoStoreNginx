package dto

import (
	"time"

	"uniStore/Backend/internal/domain/models"
)

// LibraryItemDTO represents a library item response
type LibraryItemDTO struct {
	ID        int       `json:"id"`
	Game      *GameDTO  `json:"game,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// LibraryResponseDTO represents a library response
type LibraryResponseDTO struct {
	ID        int              `json:"id"`
	UserID    int              `json:"user_id"`
	User      *UserResponseDTO `json:"user,omitempty"`
	Items     []LibraryItemDTO `json:"items"`
	CreatedAt time.Time        `json:"created_at"`
}

// LibraryItemDTOFromModel converts a LibraryItem model and Game model to LibraryItemDTO
func LibraryItemDTOFromModel(libraryItem *models.LibraryItem, game *models.Game) *LibraryItemDTO {
	dto := &LibraryItemDTO{
		ID:        libraryItem.ID,
		CreatedAt: libraryItem.CreatedAt,
	}

	// Add full game data if available
	if game != nil {
		dto.Game = GameDTOFromModel(game)
	}

	return dto
}

// LibraryResponseDTOFromModel converts a Library model and library items to LibraryResponseDTO
func LibraryResponseDTOFromModel(library *models.Library, libraryItems []*models.LibraryItem, libraryItemDTOs []LibraryItemDTO) *LibraryResponseDTO {
	dto := &LibraryResponseDTO{
		ID:        library.ID,
		UserID:    library.UserID,
		Items:     libraryItemDTOs,
		CreatedAt: library.CreatedAt,
	}

	// Add user data if available
	if library.User != nil {
		dto.User = UserResponseDTOFromModel(library.User)
	}

	return dto
}
