package repositories

import (
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"
)

// GameRepositoryImpl implementation
type gameRepositoryImpl struct {
	db *database.Database
}

// NewGameRepository creates a new game repository
func NewGameRepository(db *database.Database) models.GameRepository {
	return &gameRepositoryImpl{db: db}
}

// Create implements models.GameRepository.
func (g *gameRepositoryImpl) Create(game *models.Game) error {
	return g.db.DB.Create(game).Error
}

// Delete implements models.GameRepository.
func (g *gameRepositoryImpl) Delete(id int) error {
	return g.db.DB.Delete(&models.Game{}, id).Error
}

// FindAll implements models.GameRepository.
func (g *gameRepositoryImpl) FindAll(limit int, offset int) ([]*models.Game, error) {
	var games []*models.Game
	err := g.db.DB.Limit(limit).Offset(offset).
		Preload("Developer").
		Preload("Category").
		Find(&games).Error
	return games, err
}

// FindByCategory implements models.GameRepository.
func (g *gameRepositoryImpl) FindByCategory(categoryID int) ([]*models.Game, error) {
	var games []*models.Game
	err := g.db.DB.Where("category_id = ?", categoryID).
		Preload("Developer").
		Preload("Category").
		Find(&games).Error
	return games, err
}

// FindByDeveloper implements models.GameRepository.
func (g *gameRepositoryImpl) FindByDeveloper(developerID int) ([]*models.Game, error) {
	var games []*models.Game
	err := g.db.DB.Where("developer_id = ?", developerID).
		Preload("Developer").
		Preload("Category").
		Find(&games).Error
	return games, err
}

// FindByID implements models.GameRepository.
func (g *gameRepositoryImpl) FindByID(id int) (*models.Game, error) {
	var game models.Game
	err := g.db.DB.Preload("Developer").
		Preload("Category").
		First(&game, id).Error
	return &game, err
}

// Update implements models.GameRepository.
func (g *gameRepositoryImpl) Update(game *models.Game) error {
	return g.db.DB.Save(game).Error
}
