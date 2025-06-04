package services

import (
	"time"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/interfaces/dto"
)

// DeveloperServiceImpl implements DeveloperService interface
type DeveloperServiceImpl struct {
	developerRepo models.DeveloperRepository
}

// NewDeveloperService creates a new developer service
func NewDeveloperService(developerRepo models.DeveloperRepository) DeveloperService {
	return &DeveloperServiceImpl{
		developerRepo: developerRepo,
	}
}

// CreateDeveloper creates a new developer
func (s *DeveloperServiceImpl) CreateDeveloper(developerDTO *dto.DeveloperCreateDTO) (*dto.DeveloperDTO, error) {
	// Convert DTO to model
	developer := developerDTO.ToModel()

	// Set timestamps
	developer.CreatedAt = time.Now()
	developer.UpdatedAt = time.Now()

	// Create developer in repository
	if err := s.developerRepo.Create(developer); err != nil {
		return nil, err
	}

	// Convert model back to DTO for response
	return dto.DeveloperDTOFromModel(developer), nil
}

// GetDeveloperByID gets a developer by ID
func (s *DeveloperServiceImpl) GetDeveloperByID(id int) (*dto.DeveloperDTO, error) {
	// Get developer from repository
	developer, err := s.developerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert model to DTO for response
	return dto.DeveloperDTOFromModel(developer), nil
}

// GetAllDevelopers gets all developers
func (s *DeveloperServiceImpl) GetAllDevelopers(limit, offset int) ([]*dto.DeveloperDTO, error) {
	// Get developers from repository
	developers, err := s.developerRepo.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert models to DTOs for response
	return dto.DeveloperDTOsFromModels(developers), nil
}

// UpdateDeveloper updates a developer
func (s *DeveloperServiceImpl) UpdateDeveloper(id int, developerDTO *dto.DeveloperUpdateDTO) (*dto.DeveloperDTO, error) {
	// Get existing developer
	existingDeveloper, err := s.developerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert DTO to model
	updateData := developerDTO.ToUpdateModel(id)

	// Update fields if provided
	if updateData.Name != "" {
		existingDeveloper.Name = updateData.Name
	}
	if updateData.Description != "" {
		existingDeveloper.Description = updateData.Description
	}
	if updateData.Country != "" {
		existingDeveloper.Country = updateData.Country
	}
	if updateData.WebsiteURL != "" {
		existingDeveloper.WebsiteURL = updateData.WebsiteURL
	}

	// Update timestamp
	existingDeveloper.UpdatedAt = time.Now()

	// Update developer in repository
	if err := s.developerRepo.Update(existingDeveloper); err != nil {
		return nil, err
	}

	// Convert updated model to DTO for response
	return dto.DeveloperDTOFromModel(existingDeveloper), nil
}

// DeleteDeveloper deletes a developer
func (s *DeveloperServiceImpl) DeleteDeveloper(id int) error {
	return s.developerRepo.Delete(id)
}
