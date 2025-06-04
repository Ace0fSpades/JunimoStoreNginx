package repositories

import (
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"
)

// RoleRepositoryImpl implementation
type roleRepositoryImpl struct {
	db *database.Database
}

// NewRoleRepository creates a new role repository
func NewRoleRepository(db *database.Database) models.RoleRepository {
	return &roleRepositoryImpl{db: db}
}

// Create creates a new role
func (r *roleRepositoryImpl) Create(role *models.Role) error {
	return r.db.DB.Create(role).Error
}

// FindByID finds a role by ID
func (r *roleRepositoryImpl) FindByID(id int) (*models.Role, error) {
	var role models.Role
	err := r.db.DB.First(&role, id).Error
	return &role, err
}

// FindByType finds a role by type
func (r *roleRepositoryImpl) FindByType(roleType string) (*models.Role, error) {
	var role models.Role
	err := r.db.DB.Where("type = ?", roleType).First(&role).Error
	return &role, err
}

// Update updates a role
func (r *roleRepositoryImpl) Update(role *models.Role) error {
	return r.db.DB.Save(role).Error
}

// FindAll finds all roles
func (r *roleRepositoryImpl) FindAll() ([]*models.Role, error) {
	var roles []*models.Role
	err := r.db.DB.Find(&roles).Error
	return roles, err
}

// Delete deletes a role by ID
func (r *roleRepositoryImpl) Delete(id int) error {
	return r.db.DB.Delete(&models.Role{}, id).Error
}
