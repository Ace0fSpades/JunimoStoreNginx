package services

import (
	"errors"
	"time"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/interfaces/dto"
)

// OrderServiceImpl implements OrderService interface
type OrderServiceImpl struct {
	orderRepo models.OrderRepository
	cartRepo  models.CartRepository
	gameRepo  models.GameRepository
}

// NewOrderService creates a new order service
func NewOrderService(orderRepo models.OrderRepository, cartRepo models.CartRepository, gameRepo models.GameRepository) OrderService {
	return &OrderServiceImpl{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
		gameRepo:  gameRepo,
	}
}

// CreateOrderFromCart creates an order from a user's cart
func (s *OrderServiceImpl) CreateOrderFromCart(userID int) (*dto.OrderResponseDTO, error) {
	// Get cart items
	cartItems, err := s.cartRepo.GetCartItems(userID)
	if err != nil {
		return nil, err
	}

	// Check that the cart is not empty
	if len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	// Calculate the total cost
	var total float64
	orderItems := make([]*models.OrderItem, 0, len(cartItems))

	for _, item := range cartItems {
		// Get the current price of the game
		game, err := s.gameRepo.FindByID(item.GameID)
		if err != nil {
			return nil, err
		}

		// Create an order item
		orderItem := &models.OrderItem{
			GameID:   item.GameID,
			Game:     game,
			Quantity: item.Quantity,
			Price:    game.Price, // Lock in the current price
		}

		orderItems = append(orderItems, orderItem)
		total += game.Price * float64(item.Quantity)
	}

	// Create the order
	order := &models.Order{
		UserID:     userID,
		OrderItems: orderItems,
		TotalCost:  total,
		Status:     "new",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Save the order
	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	// Clear the cart
	if err := s.cartRepo.ClearCart(userID); err != nil {
		return nil, err
	}

	// Convert to DTO for response
	return dto.OrderResponseDTOFromModel(order, orderItems), nil
}

// CreateOrder creates a new order
func (s *OrderServiceImpl) CreateOrder(orderDTO *dto.OrderCreateDTO) (*dto.OrderResponseDTO, error) {
	// Validate order items
	if len(orderDTO.Items) == 0 {
		return nil, errors.New("order must have at least one item")
	}

	// Calculate total and create order items
	var total float64
	orderItems := make([]*models.OrderItem, 0, len(orderDTO.Items))

	for _, item := range orderDTO.Items {
		// Get the current price of the game
		game, err := s.gameRepo.FindByID(item.GameID)
		if err != nil {
			return nil, err
		}

		// Create an order item
		orderItem := &models.OrderItem{
			GameID:   item.GameID,
			Game:     game,
			Quantity: item.Quantity,
			Price:    game.Price, // Lock in the current price
		}

		orderItems = append(orderItems, orderItem)
		total += game.Price * float64(item.Quantity)
	}

	// Create the order
	order := &models.Order{
		UserID:     orderDTO.UserID,
		OrderItems: orderItems,
		TotalCost:  total,
		Status:     "new",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Save the order
	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	// Convert to DTO for response
	return dto.OrderResponseDTOFromModel(order, orderItems), nil
}

// GetOrderByID gets an order by ID
func (s *OrderServiceImpl) GetOrderByID(id int) (*dto.OrderResponseDTO, error) {
	// Get order from repository
	order, err := s.orderRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert to DTO for response
	return dto.OrderResponseDTOFromModel(order, order.OrderItems), nil
}

// GetUserOrders gets all orders for a user
func (s *OrderServiceImpl) GetUserOrders(userID int) ([]*dto.OrderResponseDTO, error) {
	// Get orders from repository
	orders, err := s.orderRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Convert to DTOs for response
	return dto.OrderResponseDTOsFromModels(orders), nil
}

// GetAllOrders gets all orders
func (s *OrderServiceImpl) GetAllOrders(limit, offset int) ([]*dto.OrderResponseDTO, error) {
	// Get orders from repository
	orders, err := s.orderRepo.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert to DTOs for response
	return dto.OrderResponseDTOsFromModels(orders), nil
}

// UpdateOrderStatus updates an order's status
func (s *OrderServiceImpl) UpdateOrderStatus(id int, statusDTO *dto.OrderUpdateDTO) (*dto.OrderResponseDTO, error) {
	// Get existing order
	order, err := s.orderRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update status
	order.Status = statusDTO.Status
	order.UpdatedAt = time.Now()

	// Save changes
	if err := s.orderRepo.Update(order); err != nil {
		return nil, err
	}

	// Convert to DTO for response
	return dto.OrderResponseDTOFromModel(order, order.OrderItems), nil
}
