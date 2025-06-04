package repositories

import (
	"errors"
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"

	"gorm.io/gorm"
)

// FavoriteRepositoryImpl implementation
type favoriteRepositoryImpl struct {
	db *database.Database
}

// NewFavoriteRepository creates a new favorite repository
func NewFavoriteRepository(db *database.Database) models.FavoriteRepository {
	return &favoriteRepositoryImpl{db: db}
}

// Create creates a new favorite list
func (f *favoriteRepositoryImpl) Create(favorite *models.Favorite) error {
	return f.db.DB.Create(favorite).Error
}

// AddGameToFavorite implements models.FavoriteRepository.
func (f *favoriteRepositoryImpl) AddGameToFavorite(userID int, gameID int) error {
	favorite, err := f.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create a new favorite list if not exists
			favorite = &models.Favorite{
				UserID: userID,
			}
			if err := f.db.DB.Create(favorite).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Check if the game already exists in favorites
	var favoriteItem models.FavoriteItem
	result := f.db.DB.Where("favorite_id = ? AND game_id = ?", favorite.ID, gameID).First(&favoriteItem)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Add new item to favorites
			favoriteItem = models.FavoriteItem{
				FavoriteID: favorite.ID,
				GameID:     gameID,
			}
			return f.db.DB.Create(&favoriteItem).Error
		}
		return result.Error
	}

	// Item already exists, nothing to do
	return nil
}

// ClearFavorite implements models.FavoriteRepository.
func (f *favoriteRepositoryImpl) ClearFavorite(userID int) error {
	favorite, err := f.FindByUserID(userID)
	if err != nil {
		return err
	}

	return f.db.DB.Where("favorite_id = ?", favorite.ID).Delete(&models.FavoriteItem{}).Error
}

// FindByUserID implements models.FavoriteRepository.
func (f *favoriteRepositoryImpl) FindByUserID(userID int) (*models.Favorite, error) {
	var favorite models.Favorite
	err := f.db.DB.Where("user_id = ?", userID).First(&favorite).Error
	return &favorite, err
}

// GetFavoriteItems implements models.FavoriteRepository.
func (f *favoriteRepositoryImpl) GetFavoriteItems(userID int) ([]*models.FavoriteItem, error) {
	favorite, err := f.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var favoriteItems []*models.FavoriteItem
	err = f.db.DB.Where("favorite_id = ?", favorite.ID).
		Preload("Game").
		Find(&favoriteItems).Error

	return favoriteItems, err
}

// RemoveGameFromFavorite implements models.FavoriteRepository.
func (f *favoriteRepositoryImpl) RemoveGameFromFavorite(userID int, gameID int) error {
	favorite, err := f.FindByUserID(userID)
	if err != nil {
		return err
	}

	return f.db.DB.Where("favorite_id = ? AND game_id = ?", favorite.ID, gameID).
		Delete(&models.FavoriteItem{}).Error
}
