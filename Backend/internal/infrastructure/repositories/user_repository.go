package repositories

import (
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"
)

// UserRepositoryImpl implementation
type userRepositoryImpl struct {
	db *database.Database
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *database.Database) models.UserRepository {
	return &userRepositoryImpl{db: db}
}

// Create creates a new user
func (r *userRepositoryImpl) Create(user *models.User) error {
	return r.db.DB.Create(user).Error
}

// FindByID finds a user by ID
func (r *userRepositoryImpl) FindByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.DB.Preload("Role").First(&user, id).Error
	return &user, err
}

// FindByEmail finds a user by email
func (r *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.DB.Preload("Role").Where("email = ?", email).First(&user).Error
	return &user, err
}

// FindByNickname finds a user by nickname
func (r *userRepositoryImpl) FindByNickname(nickname string) (*models.User, error) {
	var user models.User
	err := r.db.DB.Preload("Role").Where("nickname = ?", nickname).First(&user).Error
	return &user, err
}

// Update updates a user
func (r *userRepositoryImpl) Update(user *models.User) error {
	return r.db.DB.Save(user).Error
}

// Delete deletes a user
func (r *userRepositoryImpl) Delete(id int) error {
	return r.db.DB.Delete(&models.User{}, id).Error
}

// FindAll finds all users
func (r *userRepositoryImpl) FindAll(limit, offset int) ([]*models.User, error) {
	var users []*models.User
	err := r.db.DB.Preload("Role").Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}
