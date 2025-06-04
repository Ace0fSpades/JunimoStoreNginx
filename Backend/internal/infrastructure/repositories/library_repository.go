package repositories

import (
	"errors"
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"

	"gorm.io/gorm"
)

// LibraryRepositoryImpl implementation
type libraryRepositoryImpl struct {
	db *database.Database
}

// NewLibraryRepository creates a new library repository
func NewLibraryRepository(db *database.Database) models.LibraryRepository {
	return &libraryRepositoryImpl{db: db}
}

// Create creates a new library
func (l *libraryRepositoryImpl) Create(library *models.Library) error {
	return l.db.DB.Create(library).Error
}

// AddGameToLibrary implements models.LibraryRepository.
func (l *libraryRepositoryImpl) AddGameToLibrary(userID int, gameID int) error {
	library, err := l.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create a new library if not exists
			library = &models.Library{
				UserID: userID,
			}
			if err := l.db.DB.Create(library).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Check if the game already exists in the library
	var libraryItem models.LibraryItem
	result := l.db.DB.Where("library_id = ? AND game_id = ?", library.ID, gameID).First(&libraryItem)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Add new item to library
			libraryItem = models.LibraryItem{
				LibraryID: library.ID,
				GameID:    gameID,
			}
			return l.db.DB.Create(&libraryItem).Error
		}
		return result.Error
	}

	// Item already exists, nothing to do
	return nil
}

// FindByUserID implements models.LibraryRepository.
func (l *libraryRepositoryImpl) FindByUserID(userID int) (*models.Library, error) {
	var library models.Library
	err := l.db.DB.Where("user_id = ?", userID).First(&library).Error
	return &library, err
}

// GetLibraryItems implements models.LibraryRepository.
func (l *libraryRepositoryImpl) GetLibraryItems(userID int) ([]*models.LibraryItem, error) {
	library, err := l.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var libraryItems []*models.LibraryItem
	err = l.db.DB.Where("library_id = ?", library.ID).
		Preload("Game").
		Find(&libraryItems).Error

	return libraryItems, err
}
