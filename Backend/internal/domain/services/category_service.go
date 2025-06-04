package services

import (
	"time"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/interfaces/dto"
)

// CategoryServiceImpl implements CategoryService interface
type CategoryServiceImpl struct {
	categoryRepo models.CategoryRepository
}

// NewCategoryService creates a new category service
func NewCategoryService(categoryRepo models.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		categoryRepo: categoryRepo,
	}
}

// CreateCategory creates a new category
func (s *CategoryServiceImpl) CreateCategory(categoryDTO *dto.CategoryCreateDTO) (*dto.CategoryDTO, error) {
	// Convert DTO to model
	category := categoryDTO.ToModel()

	// Set timestamps
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	// Create category in repository
	if err := s.categoryRepo.Create(category); err != nil {
		return nil, err
	}

	// Convert model back to DTO for response
	return dto.CategoryDTOFromModel(category), nil
}

// GetCategoryByID gets a category by ID
func (s *CategoryServiceImpl) GetCategoryByID(id int) (*dto.CategoryDTO, error) {
	// Get category from repository
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert model to DTO for response
	return dto.CategoryDTOFromModel(category), nil
}

// GetAllCategories gets all categories
func (s *CategoryServiceImpl) GetAllCategories(limit, offset int) ([]*dto.CategoryDTO, error) {
	// Get categories from repository
	categories, err := s.categoryRepo.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert models to DTOs for response
	return dto.CategoryDTOsFromModels(categories), nil
}

// UpdateCategory updates a category
func (s *CategoryServiceImpl) UpdateCategory(id int, categoryDTO *dto.CategoryUpdateDTO) (*dto.CategoryDTO, error) {
	// Get existing category
	existingCategory, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert DTO to model
	updateData := categoryDTO.ToUpdateModel(id)

	// Update fields if provided
	if updateData.Name != "" {
		existingCategory.Name = updateData.Name
	}
	if updateData.Description != "" {
		existingCategory.Description = updateData.Description
	}

	// Update timestamp
	existingCategory.UpdatedAt = time.Now()

	// Update category in repository
	if err := s.categoryRepo.Update(existingCategory); err != nil {
		return nil, err
	}

	// Convert updated model to DTO for response
	return dto.CategoryDTOFromModel(existingCategory), nil
}

// DeleteCategory deletes a category
func (s *CategoryServiceImpl) DeleteCategory(id int) error {
	return s.categoryRepo.Delete(id)
}
