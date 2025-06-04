package services

import (
	"errors"
	"fmt"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/interfaces/dto"
)

// FavoriteServiceImpl implements FavoriteService interface
type FavoriteServiceImpl struct {
	favoriteRepo models.FavoriteRepository
	gameRepo     models.GameRepository
}

// NewFavoriteService creates a new favorite service
func NewFavoriteService(favoriteRepo models.FavoriteRepository, gameRepo models.GameRepository) FavoriteService {
	return &FavoriteServiceImpl{
		favoriteRepo: favoriteRepo,
		gameRepo:     gameRepo,
	}
}

// GetFavorite gets a user's favorite games
func (s *FavoriteServiceImpl) GetFavorite(userID int) (*dto.FavoriteResponseDTO, error) {
	// Get favorite from repository
	favorite, err := s.favoriteRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Get favorite items
	favoriteItems, err := s.favoriteRepo.GetFavoriteItems(userID)
	if err != nil {
		return nil, err
	}

	// Convert favorite items to DTOs
	favoriteItemDTOs := make([]dto.FavoriteItemDTO, 0, len(favoriteItems))
	for _, item := range favoriteItems {
		if item.Game != nil {
			favoriteItemDTO := dto.FavoriteItemDTOFromModel(item, item.Game)
			favoriteItemDTOs = append(favoriteItemDTOs, *favoriteItemDTO)
		}
	}

	// Create favorite response DTO
	return dto.FavoriteResponseDTOFromModel(favorite, favoriteItems, favoriteItemDTOs), nil
}

// AddGameToFavorite adds a game to a user's favorites
func (s *FavoriteServiceImpl) AddGameToFavorite(userID, gameID int) error {
	// Check if the game exists
	game, err := s.gameRepo.FindByID(gameID)
	if err != nil {
		return fmt.Errorf("game not found: %w", err)
	}

	if game == nil {
		return errors.New("game not found")
	}

	// Add game to favorites
	return s.favoriteRepo.AddGameToFavorite(userID, gameID)
}

// RemoveGameFromFavorite removes a game from a user's favorites
func (s *FavoriteServiceImpl) RemoveGameFromFavorite(userID, gameID int) error {
	return s.favoriteRepo.RemoveGameFromFavorite(userID, gameID)
}

// ClearFavorite clears a user's favorites
func (s *FavoriteServiceImpl) ClearFavorite(userID int) error {
	return s.favoriteRepo.ClearFavorite(userID)
}
