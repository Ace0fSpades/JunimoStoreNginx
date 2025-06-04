package services

import (
	"uniStore/Backend/internal/interfaces/dto"
)

// UserService defines business logic for user operations
type UserService interface {
	Register(userDTO *dto.UserSignupDTO) (*dto.UserResponseDTO, error)
	Login(loginDTO *dto.UserLoginDTO) (*dto.AuthResponseDTO, error)
	GetUserByID(id int) (*dto.UserResponseDTO, error)
	GetAllUsers(limit, offset int) ([]*dto.UserResponseDTO, error)
	UpdateUser(id int, userDTO *dto.UserUpdateDTO) (*dto.UserResponseDTO, error)
	VerifyPassword(password, hashedPassword string) bool
	HashPassword(password string) (string, error)
	GenerateTokens(email, nickname, role string, id int) (string, string, error)
	RefreshToken(refreshToken string) (*dto.AuthResponseDTO, error)
	AddPoints(userID int, points int) (*dto.UserResponseDTO, error)
}

// RoleService defines business logic for role operations
type RoleService interface {
	CreateRole(roleDTO *dto.RoleCreateDTO) (*dto.RoleDTO, error)
	GetRoleByID(id int) (*dto.RoleDTO, error)
	GetAllRoles() ([]*dto.RoleDTO, error)
	UpdateRole(id int, roleDTO *dto.RoleUpdateDTO) (*dto.RoleDTO, error)
	DeleteRole(id int) error
}

// GameService defines business logic for game operations
type GameService interface {
	CreateGame(gameDTO *dto.GameCreateDTO) (*dto.GameDTO, error)
	GetGameByID(id int) (*dto.GameDTO, error)
	GetAllGames(limit, offset int) ([]*dto.GameDTO, error)
	UpdateGame(id int, gameDTO *dto.GameUpdateDTO) (*dto.GameDTO, error)
	DeleteGame(id int) error
	SearchGamesByTitle(title string, limit, offset int) ([]*dto.GameDTO, error)
	GetGamesByCategory(categoryID int) ([]*dto.GameDTO, error)
	GetGamesByDeveloper(developerID int) ([]*dto.GameDTO, error)
	GetTopSellingGames(limit int) ([]*dto.GameDTO, error)
	GetDiscountedGames(limit int) ([]*dto.GameDTO, error)
}

// CategoryService defines business logic for category operations
type CategoryService interface {
	CreateCategory(categoryDTO *dto.CategoryCreateDTO) (*dto.CategoryDTO, error)
	GetCategoryByID(id int) (*dto.CategoryDTO, error)
	GetAllCategories(limit, offset int) ([]*dto.CategoryDTO, error)
	UpdateCategory(id int, categoryDTO *dto.CategoryUpdateDTO) (*dto.CategoryDTO, error)
	DeleteCategory(id int) error
}

// DeveloperService defines business logic for developer operations
type DeveloperService interface {
	CreateDeveloper(developerDTO *dto.DeveloperCreateDTO) (*dto.DeveloperDTO, error)
	GetDeveloperByID(id int) (*dto.DeveloperDTO, error)
	GetAllDevelopers(limit, offset int) ([]*dto.DeveloperDTO, error)
	UpdateDeveloper(id int, developerDTO *dto.DeveloperUpdateDTO) (*dto.DeveloperDTO, error)
	DeleteDeveloper(id int) error
}

// RestrictService defines business logic for restrict operations
type RestrictService interface {
	CreateRestrict(restrictDTO *dto.RestrictCreateDTO) (*dto.RestrictDTO, error)
	GetRestrictByID(id int) (*dto.RestrictDTO, error)
	GetAllRestricts(limit, offset int) ([]*dto.RestrictDTO, error)
	UpdateRestrict(id int, restrictDTO *dto.RestrictUpdateDTO) (*dto.RestrictDTO, error)
	DeleteRestrict(id int) error
	GetRestrictsByGameID(gameID int) ([]*dto.RestrictDTO, error)
}

// CartService defines business logic for cart operations
type CartService interface {
	GetCart(userID int) (*dto.CartResponseDTO, error)
	AddGameToCart(userID int, cartItemDTO *dto.CartItemCreateDTO) error
	RemoveGameFromCart(userID, gameID int) error
	ClearCart(userID int) error
	UpdateCartItemQuantity(userID, gameID int, quantityDTO *dto.CartItemUpdateDTO) error
	CalculateCartTotal(userID int) (float64, error)
}

// FavoriteService defines business logic for favorite operations
type FavoriteService interface {
	GetFavorite(userID int) (*dto.FavoriteResponseDTO, error)
	AddGameToFavorite(userID, gameID int) error
	RemoveGameFromFavorite(userID, gameID int) error
	ClearFavorite(userID int) error
}

// LibraryService defines business logic for library operations
type LibraryService interface {
	GetLibrary(userID int) (*dto.LibraryResponseDTO, error)
}

// OrderService defines business logic for order operations
type OrderService interface {
	CreateOrderFromCart(userID int) (*dto.OrderResponseDTO, error)
	CreateOrder(orderDTO *dto.OrderCreateDTO) (*dto.OrderResponseDTO, error)
	GetOrderByID(id int) (*dto.OrderResponseDTO, error)
	GetUserOrders(userID int) ([]*dto.OrderResponseDTO, error)
	GetAllOrders(limit, offset int) ([]*dto.OrderResponseDTO, error)
	UpdateOrderStatus(id int, statusDTO *dto.OrderUpdateDTO) (*dto.OrderResponseDTO, error)
}

// ReviewService defines business logic for review operations
type ReviewService interface {
	CreateReview(reviewDTO *dto.ReviewCreateDTO) (*dto.ReviewResponseDTO, error)
	GetReviewByID(id int) (*dto.ReviewResponseDTO, error)
	GetReviewsByGameID(gameID int) ([]*dto.ReviewResponseDTO, error)
	UpdateReview(id, userID int, reviewDTO *dto.ReviewUpdateDTO) (*dto.ReviewResponseDTO, error)
	DeleteReview(id, userID int) error
}

// AuthService defines business logic for authentication operations
type AuthService interface {
	VerifyToken(token string) (int, string, error)
	MatchUserTypeToID(userID int, roleType string) error
	RefreshUserToken(token string) (string, string, error)
}
