package repositories

import (
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"
)

// ReviewRepositoryImpl implementation
type reviewRepositoryImpl struct {
	db *database.Database
}

// NewReviewRepository creates a new review repository
func NewReviewRepository(db *database.Database) models.ReviewRepository {
	return &reviewRepositoryImpl{db: db}
}

// Create implements models.ReviewRepository.
func (r *reviewRepositoryImpl) Create(review *models.Review) error {
	return r.db.DB.Create(review).Error
}

// Delete implements models.ReviewRepository.
func (r *reviewRepositoryImpl) Delete(id int) error {
	return r.db.DB.Delete(&models.Review{}, id).Error
}

// FindByGameID implements models.ReviewRepository.
func (r *reviewRepositoryImpl) FindByGameID(gameID int) ([]*models.Review, error) {
	var reviews []*models.Review
	err := r.db.DB.Where("game_id = ?", gameID).
		Preload("User").
		Find(&reviews).Error
	return reviews, err
}

// FindByID implements models.ReviewRepository.
func (r *reviewRepositoryImpl) FindByID(id int) (*models.Review, error) {
	var review models.Review
	err := r.db.DB.Where("id = ?", id).
		Preload("User").
		Preload("Game").
		First(&review).Error
	return &review, err
}

// FindByUserID implements models.ReviewRepository.
func (r *reviewRepositoryImpl) FindByUserID(userID int) ([]*models.Review, error) {
	var reviews []*models.Review
	err := r.db.DB.Where("user_id = ?", userID).
		Preload("Game").
		Find(&reviews).Error
	return reviews, err
}

// Update implements models.ReviewRepository.
func (r *reviewRepositoryImpl) Update(review *models.Review) error {
	return r.db.DB.Save(review).Error
}
