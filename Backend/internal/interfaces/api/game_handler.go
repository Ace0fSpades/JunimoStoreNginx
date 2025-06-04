package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"uniStore/Backend/internal/domain/services"
	"uniStore/Backend/internal/interfaces/dto"
)

// GameHandler handles HTTP requests related to games
type GameHandler struct {
	gameService      services.GameService
	categoryService  services.CategoryService
	developerService services.DeveloperService
	restrictService  services.RestrictService
}

// NewGameHandler creates a new game handler
func NewGameHandler(
	gameService services.GameService,
	categoryService services.CategoryService,
	developerService services.DeveloperService,
	restrictService services.RestrictService,
) *GameHandler {
	return &GameHandler{
		gameService:      gameService,
		categoryService:  categoryService,
		developerService: developerService,
		restrictService:  restrictService,
	}
}

// CreateGame handles creating a new game
// @Summary Create a new game
// @Description Creates a new game with the provided details
// @Tags Games
// @Accept json
// @Produce json
// @Param game body dto.GameCreateDTO true "Game details"
// @Success 201 {object} dto.GameDTO "Game created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/games [post]
func (h *GameHandler) CreateGame(c *gin.Context) {
	var gameDTO dto.GameCreateDTO
	if err := c.BindJSON(&gameDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create game
	createdGame, err := h.gameService.CreateGame(&gameDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdGame)
}

// GetGameByID handles getting a game by ID
// @Summary Get a game by ID
// @Description Returns a game by its ID
// @Tags Games
// @Accept json
// @Produce json
// @Param game_id path int true "Game ID"
// @Success 200 {object} dto.GameDTO "Game details"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 404 {object} map[string]interface{} "Game not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/games/{game_id} [get]
func (h *GameHandler) GetGameByID(c *gin.Context) {
	gameID := c.Param("game_id")
	id, err := strconv.Atoi(gameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	// Get game by ID
	game, err := h.gameService.GetGameByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}

// GetAllGames handles getting all games with pagination
// @Summary Get all games
// @Description Returns a paginated list of games
// @Tags Games
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} dto.GameDTO "List of games"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/games [get]
func (h *GameHandler) GetAllGames(c *gin.Context) {
	// Parse pagination parameters
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	// Get all games with pagination
	games, err := h.gameService.GetAllGames(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}

// UpdateGame handles updating a game
// @Summary Update a game
// @Description Updates a game with the provided details
// @Tags Games
// @Accept json
// @Produce json
// @Param game_id path int true "Game ID"
// @Param game body dto.GameUpdateDTO true "Game details to update"
// @Success 200 {object} dto.GameDTO "Game updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Game not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/games/{game_id} [patch]
func (h *GameHandler) UpdateGame(c *gin.Context) {
	gameID := c.Param("game_id")
	id, err := strconv.Atoi(gameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	var gameDTO dto.GameUpdateDTO
	if err := c.BindJSON(&gameDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update game
	updatedGame, err := h.gameService.UpdateGame(id, &gameDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedGame)
}

// DeleteGame handles deleting a game
// @Summary Delete a game
// @Description Deletes a game by its ID
// @Tags Games
// @Accept json
// @Produce json
// @Param game_id path int true "Game ID"
// @Success 200 {object} map[string]interface{} "Game deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/games/{game_id} [delete]
func (h *GameHandler) DeleteGame(c *gin.Context) {
	gameID := c.Param("game_id")
	id, err := strconv.Atoi(gameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	// Delete game
	if err := h.gameService.DeleteGame(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}

// SearchGamesByTitle handles searching for games by title
// @Summary Search games by title
// @Description Searches for games by title
// @Tags Games
// @Accept json
// @Produce json
// @Param title query string true "Game title to search for"
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} dto.GameDTO "List of games matching the search criteria"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/games/search [get]
func (h *GameHandler) SearchGamesByTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title parameter is required"})
		return
	}

	// Parse pagination parameters
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	// Search games by title
	games, err := h.gameService.SearchGamesByTitle(title, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}

// GetGamesByCategory handles getting games by category
// @Summary Get games by category
// @Description Returns games belonging to a specific category
// @Tags Games
// @Accept json
// @Produce json
// @Param category_id path int true "Category ID"
// @Success 200 {array} dto.GameDTO "List of games in the category"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/games/category/{category_id} [get]
func (h *GameHandler) GetGamesByCategory(c *gin.Context) {
	categoryID := c.Param("category_id")
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	// Get games by category
	games, err := h.gameService.GetGamesByCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}

// GetAllCategories handles getting all categories
// @Summary Get all categories
// @Description Returns a list of all game categories
// @Tags Categories
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(100)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} dto.CategoryDTO "List of categories"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/categories [get]
func (h *GameHandler) GetAllCategories(c *gin.Context) {
	// Parse pagination parameters
	limitStr := c.DefaultQuery("limit", "100")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	// Get all categories
	categories, err := h.categoryService.GetAllCategories(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetAllDevelopers handles getting all developers
// @Summary Get all developers
// @Description Returns a list of all game developers
// @Tags Developers
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(100)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} dto.DeveloperDTO "List of developers"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/developers [get]
func (h *GameHandler) GetAllDevelopers(c *gin.Context) {
	// Parse pagination parameters
	limitStr := c.DefaultQuery("limit", "100")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	// Get all developers
	developers, err := h.developerService.GetAllDevelopers(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, developers)
}

// GetTopSellingGames handles getting top selling games
// @Summary Get top selling games
// @Description Returns a list of top selling games
// @Tags Games
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(3)
// @Success 200 {array} dto.GameDTO "List of top selling games"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/games/top-selling [get]
func (h *GameHandler) GetTopSellingGames(c *gin.Context) {
	// Parse limit parameter
	limitStr := c.DefaultQuery("limit", "3")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	// Get top selling games
	games, err := h.gameService.GetTopSellingGames(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}

// GetDiscountedGames handles getting games with discounts
// @Summary Get discounted games
// @Description Returns a list of games with discounts
// @Tags Games
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(4)
// @Success 200 {array} dto.GameDTO "List of discounted games"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/games/discounted [get]
func (h *GameHandler) GetDiscountedGames(c *gin.Context) {
	// Parse limit parameter
	limitStr := c.DefaultQuery("limit", "4")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	// Get discounted games
	games, err := h.gameService.GetDiscountedGames(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}
