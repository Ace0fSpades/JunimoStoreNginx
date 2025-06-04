package api

import (
	"github.com/gin-gonic/gin"

	"uniStore/Backend/internal/domain/services"
	"uniStore/Backend/internal/infrastructure/database"
	"uniStore/Backend/internal/infrastructure/repositories"
	"uniStore/Backend/internal/interfaces/middleware"
	"uniStore/Backend/internal/utils"
)

// Server represents the API server
type Server struct {
	DB              *database.Database
	Router          *gin.Engine
	UserHandler     *UserHandler
	GameHandler     *GameHandler
	CartHandler     *CartHandler
	OrderHandler    *OrderHandler
	ReviewHandler   *ReviewHandler
	FavoriteHandler *FavoriteHandler
	LibraryHandler  *LibraryHandler
}

// NewServer creates a new API server
func NewServer(db *database.Database, router *gin.Engine) *Server {
	// Initialize auth utils
	authUtils := utils.NewAuthUtils()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	roleRepo := repositories.NewRoleRepository(db)
	gameRepo := repositories.NewGameRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	developerRepo := repositories.NewDeveloperRepository(db)
	cartRepo := repositories.NewCartRepository(db)
	favoriteRepo := repositories.NewFavoriteRepository(db)
	libraryRepo := repositories.NewLibraryRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	reviewRepo := repositories.NewReviewRepository(db)
	restrictRepo := repositories.NewRestrictRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, roleRepo, authUtils)
	userService := services.NewUserService(userRepo, cartRepo, favoriteRepo, libraryRepo, authUtils)
	roleService := services.NewRoleService(roleRepo)
	gameService := services.NewGameService(gameRepo, categoryRepo, developerRepo, restrictRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	developerService := services.NewDeveloperService(developerRepo)
	restrictService := services.NewRestrictService(restrictRepo)
	cartService := services.NewCartService(cartRepo, gameRepo)
	favoriteService := services.NewFavoriteService(favoriteRepo, gameRepo)
	libraryService := services.NewLibraryService(libraryRepo, gameRepo)
	orderService := services.NewOrderService(orderRepo, cartRepo, gameRepo)
	reviewService := services.NewReviewService(reviewRepo, gameRepo, userRepo)

	// Initialize handlers
	userHandler := NewUserHandler(userService, roleService, authService)
	gameHandler := NewGameHandler(gameService, categoryService, developerService, restrictService)
	cartHandler := NewCartHandler(cartService)
	orderHandler := NewOrderHandler(orderService)
	reviewHandler := NewReviewHandler(reviewService)
	favoriteHandler := NewFavoriteHandler(favoriteService)
	libraryHandler := NewLibraryHandler(libraryService)

	return &Server{
		DB:              db,
		Router:          router,
		UserHandler:     userHandler,
		GameHandler:     gameHandler,
		CartHandler:     cartHandler,
		OrderHandler:    orderHandler,
		ReviewHandler:   reviewHandler,
		FavoriteHandler: favoriteHandler,
		LibraryHandler:  libraryHandler,
	}
}

// SetupRoutes sets up all API routes
func (s *Server) SetupRoutes() {
	// Apply CORS middleware
	s.Router.Use(middleware.CORSMiddleware())

	// API v1 routes
	v1 := s.Router.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", s.UserHandler.Register)
			auth.POST("/login", s.UserHandler.Login)
			auth.POST("/refresh", s.UserHandler.RefreshToken)
		}

		// User routes (some protected)
		users := v1.Group("/users")
		{
			// Admin-only routes
			adminRoutes := users.Group("/")
			adminRoutes.Use(middleware.Authenticate(), middleware.AuthorizeAdmin())
			adminRoutes.GET("/", s.UserHandler.GetAllUsers) // Admin only - get all users

			// User-specific routes (require authentication)
			userRoutes := users.Group("/")
			userRoutes.Use(middleware.Authenticate())
			userRoutes.GET("/:user_id", s.UserHandler.GetUserByID)
			userRoutes.PATCH("/:user_id", s.UserHandler.UpdateUser)
		}

		// Game routes (public)
		games := v1.Group("/games")
		{
			// Public routes
			games.GET("/", s.GameHandler.GetAllGames)
			games.GET("/search", s.GameHandler.SearchGamesByTitle)
			games.GET("/top-selling", s.GameHandler.GetTopSellingGames)
			games.GET("/discounted", s.GameHandler.GetDiscountedGames)
			games.GET("/category/:category_id", s.GameHandler.GetGamesByCategory)
			games.GET("/:game_id", s.GameHandler.GetGameByID)

			// Admin-only routes
			adminRoutes := games.Group("/")
			adminRoutes.Use(middleware.Authenticate(), middleware.AuthorizeAdmin())
			adminRoutes.POST("/", s.GameHandler.CreateGame)
			adminRoutes.PATCH("/:game_id", s.GameHandler.UpdateGame)
			adminRoutes.DELETE("/:game_id", s.GameHandler.DeleteGame)
		}

		// Categories routes (public)
		v1.GET("/categories", s.GameHandler.GetAllCategories)

		// Developers routes (public)
		v1.GET("/developers", s.GameHandler.GetAllDevelopers)

		// Cart routes (public for browsing, protected for saving)
		cart := v1.Group("/cart")
		{
			// Public cart routes (not tied to a user)
			// These would require session-based cart storage in a real implementation

			// Protected cart routes (require login)
			authenticatedCart := cart.Group("/")
			authenticatedCart.Use(middleware.Authenticate())
			authenticatedCart.GET("/:user_id", s.CartHandler.GetCart)
			authenticatedCart.POST("/:user_id/add/:game_id", s.CartHandler.AddGameToCart)
			authenticatedCart.DELETE("/:user_id/remove/:game_id", s.CartHandler.RemoveGameFromCart)
			authenticatedCart.DELETE("/:user_id/clear", s.CartHandler.ClearCart)
			authenticatedCart.PATCH("/:user_id/update/:game_id", s.CartHandler.UpdateCartItemQuantity)
		}

		// Order routes (protected)
		orders := v1.Group("/orders")
		{
			orders.Use(middleware.Authenticate())
			orders.POST("/:user_id/create", s.OrderHandler.CreateOrderFromCart)
			orders.GET("/:order_id", s.OrderHandler.GetOrderByID)
			orders.GET("/user/:user_id", s.OrderHandler.GetUserOrders)

			// Admin-only routes
			adminRoutes := orders.Group("/")
			adminRoutes.Use(middleware.AuthorizeAdmin())
			adminRoutes.GET("/", s.OrderHandler.GetAllOrders) // Admin only
		}

		// Favorite routes (protected - requires login)
		favorite := v1.Group("/favorite")
		{
			favorite.Use(middleware.Authenticate())
			favorite.GET("/:user_id", s.FavoriteHandler.GetFavorite)
			favorite.POST("/:user_id/add/:game_id", s.FavoriteHandler.AddGameToFavorite)
			favorite.DELETE("/:user_id/remove/:game_id", s.FavoriteHandler.RemoveGameFromFavorite)
			favorite.DELETE("/:user_id/clear", s.FavoriteHandler.ClearFavorite)
		}

		// Library routes (protected)
		library := v1.Group("/library")
		{
			library.Use(middleware.Authenticate())
			library.GET("/:user_id", s.LibraryHandler.GetLibrary)
		}

		// Review routes (public for viewing, protected for creating/editing)
		reviews := v1.Group("/reviews")
		{
			// Public routes
			reviews.GET("/:review_id", s.ReviewHandler.GetReviewByID)
			reviews.GET("/game/:game_id", s.ReviewHandler.GetReviewsByGameID)

			// Protected routes
			authenticatedReviews := reviews.Group("/")
			authenticatedReviews.Use(middleware.Authenticate())
			authenticatedReviews.POST("/", s.ReviewHandler.CreateReview)
			authenticatedReviews.PATCH("/:review_id/user/:user_id", s.ReviewHandler.UpdateReview)
			authenticatedReviews.DELETE("/:review_id/user/:user_id", s.ReviewHandler.DeleteReview)
		}
	}
}
