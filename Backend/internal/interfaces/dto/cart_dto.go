package dto

import (
	"uniStore/Backend/internal/domain/models"
)

// CartItemDTO represents a cart item response
type CartItemDTO struct {
	ID       int      `json:"id"`
	Game     *GameDTO `json:"game,omitempty"`
	Quantity int      `json:"quantity"`
}

// CartResponseDTO represents a cart response
type CartResponseDTO struct {
	ID        int              `json:"id"`
	UserID    int              `json:"user_id"`
	User      *UserResponseDTO `json:"user,omitempty"`
	Items     []CartItemDTO    `json:"items"`
	TotalCost float64          `json:"total_cost"`
}

// CartItemCreateDTO represents data for adding an item to cart
type CartItemCreateDTO struct {
	GameID   int `json:"game_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required,min=1"`
}

// CartItemUpdateDTO represents data for updating a cart item
type CartItemUpdateDTO struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

// CartItemDTOFromModel converts a CartItem model and Game model to CartItemDTO
func CartItemDTOFromModel(cartItem *models.CartItem, game *models.Game) *CartItemDTO {
	dto := &CartItemDTO{
		ID:       cartItem.ID,
		Quantity: cartItem.Quantity,
	}

	// Add full game data if available
	if game != nil {
		dto.Game = GameDTOFromModel(game)
	}

	return dto
}

// CartResponseDTOFromModel converts a ShoppingCart model and cart items to CartResponseDTO
func CartResponseDTOFromModel(cart *models.ShoppingCart, cartItems []*models.CartItem, cartItemDTOs []CartItemDTO) *CartResponseDTO {
	// Calculate total cost
	var totalCost float64
	for _, item := range cartItemDTOs {
		if item.Game != nil {
			totalCost += item.Game.Price * float64(item.Quantity)
		}
	}

	dto := &CartResponseDTO{
		ID:        cart.ID,
		UserID:    cart.UserID,
		Items:     cartItemDTOs,
		TotalCost: totalCost,
	}

	// Add user data if available
	if cart.User != nil {
		dto.User = UserResponseDTOFromModel(cart.User)
	}

	return dto
}
