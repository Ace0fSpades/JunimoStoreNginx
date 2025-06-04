package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"uniStore/Backend/internal/domain/services"
	"uniStore/Backend/internal/interfaces/dto"
)

// CartHandler handles HTTP requests related to shopping carts
type CartHandler struct {
	cartService services.CartService
}

// NewCartHandler creates a new cart handler
func NewCartHandler(cartService services.CartService) *CartHandler {
	return &CartHandler{
		cartService: cartService,
	}
}

// GetCart handles getting a user's shopping cart
// @Summary Get a user's shopping cart
// @Description Returns a user's shopping cart and its items
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} dto.CartResponseDTO "User's shopping cart"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/cart/{user_id} [get]
func (h *CartHandler) GetCart(c *gin.Context) {
	userID := c.Param("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if the current user can access this cart
	tokenUserID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Only allow users to access their own cart or admin to access any cart
	if tokenUserID.(int) != id && c.GetString("role") != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only access your own cart"})
		return
	}

	// Get cart using the service
	cartResponseDTO, err := h.cartService.GetCart(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total cost
	totalCost, err := h.cartService.CalculateCartTotal(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update total cost in response
	cartResponseDTO.TotalCost = totalCost

	c.JSON(http.StatusOK, cartResponseDTO)
}

// AddGameToCart handles adding a game to a user's shopping cart
// @Summary Add a game to cart
// @Description Adds a game to a user's shopping cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param game_id path int true "Game ID"
// @Success 200 {object} map[string]interface{} "Game added to cart successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 404 {object} map[string]interface{} "Game not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/cart/{user_id}/add/{game_id} [post]
func (h *CartHandler) AddGameToCart(c *gin.Context) {
	userID := c.Param("user_id")
	gameID := c.Param("game_id")

	uid, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	gid, err := strconv.Atoi(gameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	// Check if the current user can modify this cart
	tokenUserID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Only allow users to modify their own cart
	if tokenUserID.(int) != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only modify your own cart"})
		return
	}

	// Create DTO for adding item to cart
	cartItemDTO := &dto.CartItemCreateDTO{
		GameID:   gid,
		Quantity: 1, // Default quantity is 1
	}

	// Add game to cart
	if err := h.cartService.AddGameToCart(uid, cartItemDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game added to cart successfully"})
}

// RemoveGameFromCart handles removing a game from a user's shopping cart
// @Summary Remove a game from cart
// @Description Removes a game from a user's shopping cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param game_id path int true "Game ID"
// @Success 200 {object} map[string]interface{} "Game removed from cart successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/cart/{user_id}/remove/{game_id} [delete]
func (h *CartHandler) RemoveGameFromCart(c *gin.Context) {
	userID := c.Param("user_id")
	gameID := c.Param("game_id")

	uid, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	gid, err := strconv.Atoi(gameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	// Check if the current user can modify this cart
	tokenUserID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Only allow users to modify their own cart
	if tokenUserID.(int) != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only modify your own cart"})
		return
	}

	// Remove game from cart
	if err := h.cartService.RemoveGameFromCart(uid, gid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game removed from cart successfully"})
}

// ClearCart handles clearing a user's shopping cart
// @Summary Clear cart
// @Description Clears all items from a user's shopping cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} map[string]interface{} "Cart cleared successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/cart/{user_id}/clear [delete]
func (h *CartHandler) ClearCart(c *gin.Context) {
	userID := c.Param("user_id")

	uid, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if the current user can modify this cart
	tokenUserID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Only allow users to modify their own cart
	if tokenUserID.(int) != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only modify your own cart"})
		return
	}

	// Clear cart
	if err := h.cartService.ClearCart(uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared successfully"})
}

// UpdateCartItemQuantity handles updating the quantity of a game in a user's shopping cart
// @Summary Update cart item quantity
// @Description Updates the quantity of a game in a user's shopping cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param game_id path int true "Game ID"
// @Param quantity query int true "New quantity"
// @Success 200 {object} map[string]interface{} "Cart item quantity updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/cart/{user_id}/update/{game_id} [patch]
func (h *CartHandler) UpdateCartItemQuantity(c *gin.Context) {
	userID := c.Param("user_id")
	gameID := c.Param("game_id")
	quantityStr := c.Query("quantity")

	uid, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	gid, err := strconv.Atoi(gameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	// Check if the current user can modify this cart
	tokenUserID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Only allow users to modify their own cart
	if tokenUserID.(int) != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only modify your own cart"})
		return
	}

	// Create DTO for updating quantity
	quantityDTO := &dto.CartItemUpdateDTO{
		Quantity: quantity,
	}

	// Update cart item quantity
	if err := h.cartService.UpdateCartItemQuantity(uid, gid, quantityDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item quantity updated successfully"})
}
