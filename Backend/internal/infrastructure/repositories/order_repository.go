package repositories

import (
	"errors"
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"

	"gorm.io/gorm"
)

// OrderRepositoryImpl implementation
type orderRepositoryImpl struct {
	db *database.Database
}

// NewOrderRepository creates a new order repository
func NewOrderRepository(db *database.Database) models.OrderRepository {
	return &orderRepositoryImpl{db: db}
}

// Create implements models.OrderRepository.
func (o *orderRepositoryImpl) Create(order *models.Order) error {
	return o.db.DB.Create(order).Error
}

// CreateFromCart implements models.OrderRepository.
func (o *orderRepositoryImpl) CreateFromCart(userID int) (*models.Order, error) {
	// Start a transaction
	tx := o.db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get user's cart
	var cart models.ShoppingCart
	if err := tx.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Get cart items
	var cartItems []*models.CartItem
	if err := tx.Where("shopping_cart_id = ?", cart.ID).Preload("Game").Find(&cartItems).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Check if cart is empty
	if len(cartItems) == 0 {
		tx.Rollback()
		return nil, errors.New("cart is empty")
	}

	// Calculate total cost
	var totalCost float64
	for _, item := range cartItems {
		totalCost += item.Game.Price * float64(item.Quantity)
	}

	// Create new order
	order := &models.Order{
		UserID:    userID,
		TotalCost: totalCost,
		Status:    "pending",
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create order items
	for _, cartItem := range cartItems {
		orderItem := &models.OrderItem{
			OrderID:  order.ID,
			GameID:   cartItem.GameID,
			Price:    cartItem.Game.Price,
			Quantity: cartItem.Quantity,
		}
		if err := tx.Create(orderItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Add games to user's library
	var library models.Library
	result := tx.Where("user_id = ?", userID).First(&library)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Create library if not exists
			library = models.Library{
				UserID: userID,
			}
			if err := tx.Create(&library).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		} else {
			tx.Rollback()
			return nil, result.Error
		}
	}

	// Add games to library
	for _, cartItem := range cartItems {
		// Check if game already in library
		var libraryItem models.LibraryItem
		result := tx.Where("library_id = ? AND game_id = ?", library.ID, cartItem.GameID).First(&libraryItem)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// Add to library
				libraryItem = models.LibraryItem{
					LibraryID: library.ID,
					GameID:    cartItem.GameID,
				}
				if err := tx.Create(&libraryItem).Error; err != nil {
					tx.Rollback()
					return nil, err
				}
			} else {
				tx.Rollback()
				return nil, result.Error
			}
		}
	}

	// Clear cart
	if err := tx.Where("shopping_cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return order, nil
}

// FindAll implements models.OrderRepository.
func (o *orderRepositoryImpl) FindAll(limit int, offset int) ([]*models.Order, error) {
	var orders []*models.Order
	err := o.db.DB.Limit(limit).Offset(offset).
		Preload("OrderItems").
		Preload("OrderItems.Game").
		Find(&orders).Error
	return orders, err
}

// FindByID implements models.OrderRepository.
func (o *orderRepositoryImpl) FindByID(id int) (*models.Order, error) {
	var order models.Order
	err := o.db.DB.Where("id = ?", id).
		Preload("OrderItems").
		Preload("OrderItems.Game").
		First(&order).Error
	return &order, err
}

// FindByUserID implements models.OrderRepository.
func (o *orderRepositoryImpl) FindByUserID(userID int) ([]*models.Order, error) {
	var orders []*models.Order
	err := o.db.DB.Where("user_id = ?", userID).
		Preload("OrderItems").
		Preload("OrderItems.Game").
		Find(&orders).Error
	return orders, err
}

// Update implements models.OrderRepository.
func (o *orderRepositoryImpl) Update(order *models.Order) error {
	return o.db.DB.Save(order).Error
}
