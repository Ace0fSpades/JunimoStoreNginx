package repositories

import (
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"
)

// CategoryRepositoryImpl implementation
type categoryRepositoryImpl struct {
	db *database.Database
}

// NewCategoryRepository creates a new category repository
func NewCategoryRepository(db *database.Database) models.CategoryRepository {
	return &categoryRepositoryImpl{db: db}
}

// Create implements models.CategoryRepository.
func (c *categoryRepositoryImpl) Create(category *models.Category) error {
	return c.db.DB.Create(category).Error
}

// Delete implements models.CategoryRepository.
func (c *categoryRepositoryImpl) Delete(id int) error {
	return c.db.DB.Delete(&models.Category{}, id).Error
}

// FindAll implements models.CategoryRepository.
func (c *categoryRepositoryImpl) FindAll(limit int, offset int) ([]*models.Category, error) {
	var categories []*models.Category
	err := c.db.DB.Limit(limit).Offset(offset).Find(&categories).Error
	return categories, err
}

// FindByID implements models.CategoryRepository.
func (c *categoryRepositoryImpl) FindByID(id int) (*models.Category, error) {
	var category models.Category
	err := c.db.DB.First(&category, id).Error
	return &category, err
}

// Update implements models.CategoryRepository.
func (c *categoryRepositoryImpl) Update(category *models.Category) error {
	return c.db.DB.Save(category).Error
}
