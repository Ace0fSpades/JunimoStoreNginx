package services

import (
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/interfaces/dto"
)

// LibraryServiceImpl implements LibraryService interface
type LibraryServiceImpl struct {
	libraryRepo models.LibraryRepository
	gameRepo    models.GameRepository
}

// NewLibraryService creates a new library service
func NewLibraryService(libraryRepo models.LibraryRepository, gameRepo models.GameRepository) LibraryService {
	return &LibraryServiceImpl{
		libraryRepo: libraryRepo,
		gameRepo:    gameRepo,
	}
}

// GetLibrary gets a user's game library
func (s *LibraryServiceImpl) GetLibrary(userID int) (*dto.LibraryResponseDTO, error) {
	// Get library from repository
	library, err := s.libraryRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Get library items
	libraryItems, err := s.libraryRepo.GetLibraryItems(userID)
	if err != nil {
		return nil, err
	}

	// Convert library items to DTOs
	libraryItemDTOs := make([]dto.LibraryItemDTO, 0, len(libraryItems))
	for _, item := range libraryItems {
		if item.Game != nil {
			libraryItemDTO := dto.LibraryItemDTOFromModel(item, item.Game)
			libraryItemDTOs = append(libraryItemDTOs, *libraryItemDTO)
		}
	}

	// Create library response DTO
	return dto.LibraryResponseDTOFromModel(library, libraryItems, libraryItemDTOs), nil
}
