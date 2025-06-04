package repositories

import (
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"
)

// Factory provides all repositories
type Factory struct {
	UserRepository      models.UserRepository
	RoleRepository      models.RoleRepository
	GameRepository      models.GameRepository
	CategoryRepository  models.CategoryRepository
	DeveloperRepository models.DeveloperRepository
	CartRepository      models.CartRepository
	FavoriteRepository  models.FavoriteRepository
	LibraryRepository   models.LibraryRepository
	OrderRepository     models.OrderRepository
	ReviewRepository    models.ReviewRepository
	RestrictRepository  models.RestrictRepository
}

// NewFactory creates a new repository factory
func NewFactory(db *database.Database) *Factory {
	return &Factory{
		UserRepository:      NewUserRepository(db),
		RoleRepository:      NewRoleRepository(db),
		GameRepository:      NewGameRepository(db),
		CategoryRepository:  NewCategoryRepository(db),
		DeveloperRepository: NewDeveloperRepository(db),
		CartRepository:      NewCartRepository(db),
		FavoriteRepository:  NewFavoriteRepository(db),
		LibraryRepository:   NewLibraryRepository(db),
		OrderRepository:     NewOrderRepository(db),
		ReviewRepository:    NewReviewRepository(db),
		RestrictRepository:  NewRestrictRepository(db),
	}
}
