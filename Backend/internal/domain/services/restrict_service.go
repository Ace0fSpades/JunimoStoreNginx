package services

import (
	"time"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/interfaces/dto"
)

// RestrictServiceImpl implements RestrictService interface
type RestrictServiceImpl struct {
	restrictRepo models.RestrictRepository
}

// NewRestrictService creates a new restrict service
func NewRestrictService(restrictRepo models.RestrictRepository) RestrictService {
	return &RestrictServiceImpl{
		restrictRepo: restrictRepo,
	}
}

// CreateRestrict creates a new restrict
func (s *RestrictServiceImpl) CreateRestrict(restrictDTO *dto.RestrictCreateDTO) (*dto.RestrictDTO, error) {
	// Convert DTO to model
	restrict := restrictDTO.ToModel()

	// Set timestamps
	restrict.CreatedAt = time.Now()
	restrict.UpdatedAt = time.Now()

	// Create restrict in repository
	if err := s.restrictRepo.Create(restrict); err != nil {
		return nil, err
	}

	// Convert model back to DTO for response
	return dto.RestrictDTOFromModel(restrict), nil
}

// GetRestrictByID gets a restrict by ID
func (s *RestrictServiceImpl) GetRestrictByID(id int) (*dto.RestrictDTO, error) {
	// Get restrict from repository
	restrict, err := s.restrictRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert model to DTO for response
	return dto.RestrictDTOFromModel(restrict), nil
}

// GetAllRestricts gets all restricts
func (s *RestrictServiceImpl) GetAllRestricts(limit, offset int) ([]*dto.RestrictDTO, error) {
	// Get restricts from repository
	restricts, err := s.restrictRepo.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert models to DTOs for response
	return dto.RestrictDTOsFromModels(restricts), nil
}

// GetRestrictsByGameID gets all restricts for a game
func (s *RestrictServiceImpl) GetRestrictsByGameID(gameID int) ([]*dto.RestrictDTO, error) {
	// Get restricts from repository
	restricts, err := s.restrictRepo.FindByGameID(gameID)
	if err != nil {
		return nil, err
	}

	// Convert models to DTOs for response
	return dto.RestrictDTOsFromModels(restricts), nil
}

// UpdateRestrict updates a restrict
func (s *RestrictServiceImpl) UpdateRestrict(id int, restrictDTO *dto.RestrictUpdateDTO) (*dto.RestrictDTO, error) {
	// Get existing restriction
	existingRestrict, err := s.restrictRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert DTO to model
	updateData := restrictDTO.ToUpdateModel(id)

	// Update fields if provided
	if updateData.Region != "" {
		existingRestrict.Region = updateData.Region
	}

	// Update timestamp
	existingRestrict.UpdatedAt = time.Now()

	// Update restrict in repository
	if err := s.restrictRepo.Update(existingRestrict); err != nil {
		return nil, err
	}

	// Convert updated model to DTO for response
	return dto.RestrictDTOFromModel(existingRestrict), nil
}

// DeleteRestrict deletes a restrict
func (s *RestrictServiceImpl) DeleteRestrict(id int) error {
	return s.restrictRepo.Delete(id)
}
