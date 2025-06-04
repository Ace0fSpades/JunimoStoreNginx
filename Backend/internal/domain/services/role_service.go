package services

import (
	"time"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/interfaces/dto"
)

// RoleServiceImpl implements RoleService interface
type RoleServiceImpl struct {
	roleRepo models.RoleRepository
}

// NewRoleService creates a new role service
func NewRoleService(roleRepo models.RoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepo: roleRepo,
	}
}

// CreateRole creates a new role
func (s *RoleServiceImpl) CreateRole(roleDTO *dto.RoleCreateDTO) (*dto.RoleDTO, error) {
	// Convert DTO to model
	role := roleDTO.ToModel()

	// Set timestamps
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()

	// Create role in repository
	if err := s.roleRepo.Create(role); err != nil {
		return nil, err
	}

	// Convert model back to DTO for response
	return dto.RoleDTOFromModel(role), nil
}

// GetRoleByID gets a role by ID
func (s *RoleServiceImpl) GetRoleByID(id int) (*dto.RoleDTO, error) {
	// Get role from repository
	role, err := s.roleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert model to DTO for response
	return dto.RoleDTOFromModel(role), nil
}

// GetAllRoles gets all roles
func (s *RoleServiceImpl) GetAllRoles() ([]*dto.RoleDTO, error) {
	// Get roles from repository
	roles, err := s.roleRepo.FindAll()
	if err != nil {
		return nil, err
	}

	// Convert models to DTOs for response
	return dto.RoleDTOsFromModels(roles), nil
}

// UpdateRole updates a role
func (s *RoleServiceImpl) UpdateRole(id int, roleDTO *dto.RoleUpdateDTO) (*dto.RoleDTO, error) {
	// Get existing role
	existingRole, err := s.roleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert DTO to model
	updateData := roleDTO.ToUpdateModel(id)

	// Update fields if provided
	if updateData.Type != "" {
		existingRole.Type = updateData.Type
	}
	if updateData.Description != "" {
		existingRole.Description = updateData.Description
	}

	// Update timestamp
	existingRole.UpdatedAt = time.Now()

	// Update role in repository
	if err := s.roleRepo.Update(existingRole); err != nil {
		return nil, err
	}

	// Convert updated model to DTO for response
	return dto.RoleDTOFromModel(existingRole), nil
}

// DeleteRole deletes a role
func (s *RoleServiceImpl) DeleteRole(id int) error {
	return s.roleRepo.Delete(id)
}
