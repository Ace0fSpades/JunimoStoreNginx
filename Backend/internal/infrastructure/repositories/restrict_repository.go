package repositories

import (
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"
)

// RestrictRepositoryImpl implementation
type restrictRepositoryImpl struct {
	db *database.Database
}

// NewRestrictRepository creates a new restrict repository
func NewRestrictRepository(db *database.Database) models.RestrictRepository {
	return &restrictRepositoryImpl{db: db}
}

// Create implements models.RestrictRepository.
func (r *restrictRepositoryImpl) Create(restrict *models.Restrict) error {
	return r.db.DB.Create(restrict).Error
}

// Delete implements models.RestrictRepository.
func (r *restrictRepositoryImpl) Delete(id int) error {
	return r.db.DB.Delete(&models.Restrict{}, id).Error
}

// FindAll implements models.RestrictRepository.
func (r *restrictRepositoryImpl) FindAll(limit int, offset int) ([]*models.Restrict, error) {
	var restricts []*models.Restrict
	err := r.db.DB.Limit(limit).Offset(offset).
		Preload("Game").
		Find(&restricts).Error
	return restricts, err
}

// FindByGameID implements models.RestrictRepository.
func (r *restrictRepositoryImpl) FindByGameID(gameID int) ([]*models.Restrict, error) {
	var restricts []*models.Restrict
	err := r.db.DB.Where("game_id = ?", gameID).Find(&restricts).Error
	return restricts, err
}

// FindByID implements models.RestrictRepository.
func (r *restrictRepositoryImpl) FindByID(id int) (*models.Restrict, error) {
	var restrict models.Restrict
	err := r.db.DB.First(&restrict, id).Error
	return &restrict, err
}

// Update implements models.RestrictRepository.
func (r *restrictRepositoryImpl) Update(restrict *models.Restrict) error {
	return r.db.DB.Save(restrict).Error
}
