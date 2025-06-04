package services

import (
	"errors"
	"time"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/interfaces/dto"
)

// ReviewServiceImpl implements ReviewService interface
type ReviewServiceImpl struct {
	reviewRepo models.ReviewRepository
	gameRepo   models.GameRepository
	userRepo   models.UserRepository
}

// NewReviewService creates a new review service
func NewReviewService(reviewRepo models.ReviewRepository, gameRepo models.GameRepository, userRepo models.UserRepository) ReviewService {
	return &ReviewServiceImpl{
		reviewRepo: reviewRepo,
		gameRepo:   gameRepo,
		userRepo:   userRepo,
	}
}

// CreateReview creates a new review
func (s *ReviewServiceImpl) CreateReview(reviewDTO *dto.ReviewCreateDTO) (*dto.ReviewResponseDTO, error) {
	// Check if the game exists
	game, err := s.gameRepo.FindByID(reviewDTO.GameID)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return nil, errors.New("game not found")
	}

	// Check if the user exists
	user, err := s.userRepo.FindByID(reviewDTO.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Create model from DTO
	review := &models.Review{
		GameID:      reviewDTO.GameID,
		UserID:      reviewDTO.UserID,
		Title:       "Review", // Default title
		Description: reviewDTO.Comment,
		Rating:      reviewDTO.Rating,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Game:        game,
		User:        user,
	}

	// Create review in repository
	if err := s.reviewRepo.Create(review); err != nil {
		return nil, err
	}

	// Convert model to DTO for response
	return dto.ReviewResponseDTOFromModel(review), nil
}

// GetReviewByID gets a review by ID
func (s *ReviewServiceImpl) GetReviewByID(id int) (*dto.ReviewResponseDTO, error) {
	// Get review from repository
	review, err := s.reviewRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert model to DTO for response
	return dto.ReviewResponseDTOFromModel(review), nil
}

// GetReviewsByGameID gets all reviews for a game
func (s *ReviewServiceImpl) GetReviewsByGameID(gameID int) ([]*dto.ReviewResponseDTO, error) {
	// Get reviews from repository
	reviews, err := s.reviewRepo.FindByGameID(gameID)
	if err != nil {
		return nil, err
	}

	// Convert models to DTOs for response
	reviewDTOs := make([]*dto.ReviewResponseDTO, len(reviews))
	for i, review := range reviews {
		reviewDTOs[i] = dto.ReviewResponseDTOFromModel(review)
	}

	return reviewDTOs, nil
}

// UpdateReview updates a review
func (s *ReviewServiceImpl) UpdateReview(id, userID int, reviewDTO *dto.ReviewUpdateDTO) (*dto.ReviewResponseDTO, error) {
	// Get existing review
	existingReview, err := s.reviewRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Check if the review belongs to the user
	if existingReview.UserID != userID {
		return nil, errors.New("review belongs to another user")
	}

	// Update fields if provided
	existingReview.Rating = reviewDTO.Rating
	existingReview.Description = reviewDTO.Comment

	// Update timestamp
	existingReview.UpdatedAt = time.Now()

	// Update review in repository
	if err := s.reviewRepo.Update(existingReview); err != nil {
		return nil, err
	}

	// Convert updated model to DTO for response
	return dto.ReviewResponseDTOFromModel(existingReview), nil
}

// DeleteReview deletes a review
func (s *ReviewServiceImpl) DeleteReview(id, userID int) error {
	// Get existing review
	existingReview, err := s.reviewRepo.FindByID(id)
	if err != nil {
		return err
	}

	// Check if the review belongs to the user
	if existingReview.UserID != userID {
		return errors.New("review belongs to another user")
	}

	return s.reviewRepo.Delete(id)
}
