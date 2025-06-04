package services

import (
	"errors"
	"fmt"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/interfaces/dto"
)

// CartServiceImpl implements CartService interface
type CartServiceImpl struct {
	cartRepo models.CartRepository
	gameRepo models.GameRepository
}

// NewCartService creates a new cart service
func NewCartService(cartRepo models.CartRepository, gameRepo models.GameRepository) CartService {
	return &CartServiceImpl{
		cartRepo: cartRepo,
		gameRepo: gameRepo,
	}
}

// GetCart retrieves a user's shopping cart
func (s *CartServiceImpl) GetCart(userID int) (*dto.CartResponseDTO, error) {
	// Get cart from repository
	cart, err := s.cartRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Get cart items
	cartItems, err := s.cartRepo.GetCartItems(userID)
	if err != nil {
		return nil, err
	}

	// Convert cart items to DTOs
	cartItemDTOs := make([]dto.CartItemDTO, len(cartItems))
	for i, item := range cartItems {
		if item.Game != nil {
			cartItemDTO := dto.CartItemDTOFromModel(item, item.Game)
			cartItemDTOs[i] = *cartItemDTO
		}
	}

	// Create cart response DTO
	return dto.CartResponseDTOFromModel(cart, cartItems, cartItemDTOs), nil
}

// AddGameToCart adds a game to a user's shopping cart
func (s *CartServiceImpl) AddGameToCart(userID int, cartItemDTO *dto.CartItemCreateDTO) error {
	// Check if the game exists
	game, err := s.gameRepo.FindByID(cartItemDTO.GameID)
	if err != nil {
		return fmt.Errorf("game not found: %w", err)
	}

	if game == nil {
		return errors.New("game not found")
	}

	// Add game to cart with specified quantity
	return s.cartRepo.AddGameToCart(userID, cartItemDTO.GameID, cartItemDTO.Quantity)
}

// RemoveGameFromCart removes a game from a user's shopping cart
func (s *CartServiceImpl) RemoveGameFromCart(userID, gameID int) error {
	return s.cartRepo.RemoveGameFromCart(userID, gameID)
}

// ClearCart clears a user's shopping cart
func (s *CartServiceImpl) ClearCart(userID int) error {
	return s.cartRepo.ClearCart(userID)
}

// UpdateCartItemQuantity updates the quantity of a game in a user's shopping cart
func (s *CartServiceImpl) UpdateCartItemQuantity(userID, gameID int, quantityDTO *dto.CartItemUpdateDTO) error {
	if quantityDTO.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	// First remove the game from cart
	if err := s.cartRepo.RemoveGameFromCart(userID, gameID); err != nil {
		return err
	}

	// Then add it back with the new quantity
	return s.cartRepo.AddGameToCart(userID, gameID, quantityDTO.Quantity)
}

// CalculateCartTotal calculates the total cost of a user's shopping cart
func (s *CartServiceImpl) CalculateCartTotal(userID int) (float64, error) {
	// Get cart items
	cartItems, err := s.cartRepo.GetCartItems(userID)
	if err != nil {
		return 0, err
	}

	var total float64
	for _, item := range cartItems {
		if item.Game != nil {
			total += item.Game.Price * float64(item.Quantity)
		}
	}

	return total, nil
}
