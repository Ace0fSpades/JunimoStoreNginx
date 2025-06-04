package repositories

import (
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"
)

// DeveloperRepositoryImpl implementation
type developerRepositoryImpl struct {
	db *database.Database
}

// NewDeveloperRepository creates a new developer repository
func NewDeveloperRepository(db *database.Database) models.DeveloperRepository {
	return &developerRepositoryImpl{db: db}
}

// Create implements models.DeveloperRepository.
func (d *developerRepositoryImpl) Create(developer *models.Developer) error {
	return d.db.DB.Create(developer).Error
}

// Delete implements models.DeveloperRepository.
func (d *developerRepositoryImpl) Delete(id int) error {
	return d.db.DB.Delete(&models.Developer{}, id).Error
}

// FindAll implements models.DeveloperRepository.
func (d *developerRepositoryImpl) FindAll(limit int, offset int) ([]*models.Developer, error) {
	var developers []*models.Developer
	err := d.db.DB.Limit(limit).Offset(offset).Find(&developers).Error
	return developers, err
}

// FindByID implements models.DeveloperRepository.
func (d *developerRepositoryImpl) FindByID(id int) (*models.Developer, error) {
	var developer models.Developer
	err := d.db.DB.First(&developer, id).Error
	return &developer, err
}

// Update implements models.DeveloperRepository.
func (d *developerRepositoryImpl) Update(developer *models.Developer) error {
	return d.db.DB.Save(developer).Error
}
