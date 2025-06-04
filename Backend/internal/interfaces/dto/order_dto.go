package dto

import (
	"time"

	"uniStore/Backend/internal/domain/models"
)

// OrderItemDTO represents an order item for API responses
type OrderItemDTO struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id"`
	Game      *GameDTO  `json:"game,omitempty"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

// OrderResponseDTO represents an order for API responses
type OrderResponseDTO struct {
	ID        int              `json:"id"`
	UserID    int              `json:"user_id"`
	User      *UserResponseDTO `json:"user,omitempty"`
	TotalCost float64          `json:"total_cost"`
	Status    string           `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Items     []OrderItemDTO   `json:"items"`
}

// OrderCreateDTO represents data needed for creating a new order directly (not from cart)
type OrderCreateDTO struct {
	UserID int `json:"user_id" binding:"required"`
	Items  []struct {
		GameID   int `json:"game_id" binding:"required"`
		Quantity int `json:"quantity" binding:"required,min=1"`
	} `json:"items" binding:"required"`
}

// OrderUpdateDTO represents data needed for updating an order status
type OrderUpdateDTO struct {
	Status string `json:"status" binding:"required"`
}

// OrderItemDTOFromModel converts OrderItem model to OrderItemDTO
func OrderItemDTOFromModel(orderItem *models.OrderItem) *OrderItemDTO {
	dto := &OrderItemDTO{
		ID:        orderItem.ID,
		OrderID:   orderItem.OrderID,
		Price:     orderItem.Price,
		Quantity:  orderItem.Quantity,
		CreatedAt: orderItem.CreatedAt,
	}

	// Add full game data if available
	if orderItem.Game != nil {
		dto.Game = GameDTOFromModel(orderItem.Game)
	}

	return dto
}

// OrderResponseDTOFromModel converts Order model to OrderResponseDTO
func OrderResponseDTOFromModel(order *models.Order, orderItems []*models.OrderItem) *OrderResponseDTO {
	itemDTOs := make([]OrderItemDTO, len(orderItems))
	for i, item := range orderItems {
		itemDTO := OrderItemDTOFromModel(item)
		itemDTOs[i] = *itemDTO
	}

	dto := &OrderResponseDTO{
		ID:        order.ID,
		UserID:    order.UserID,
		TotalCost: order.TotalCost,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		Items:     itemDTOs,
	}

	// Add full user data if available
	if order.User != nil {
		dto.User = UserResponseDTOFromModel(order.User)
	}

	return dto
}

// OrderResponseDTOsFromModels converts a slice of Order models to a slice of OrderResponseDTOs
func OrderResponseDTOsFromModels(orders []*models.Order) []*OrderResponseDTO {
	dtos := make([]*OrderResponseDTO, len(orders))
	for i, order := range orders {
		dtos[i] = OrderResponseDTOFromModel(order, order.OrderItems)
	}
	return dtos
}
