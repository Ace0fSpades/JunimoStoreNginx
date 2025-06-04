package repositories

import (
	"errors"
	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/infrastructure/database"

	"gorm.io/gorm"
)

// CartRepositoryImpl implementation
type cartRepositoryImpl struct {
	db *database.Database
}

// NewCartRepository creates a new cart repository
func NewCartRepository(db *database.Database) models.CartRepository {
	return &cartRepositoryImpl{db: db}
}

// Create creates a new shopping cart
func (c *cartRepositoryImpl) Create(cart *models.ShoppingCart) error {
	return c.db.DB.Create(cart).Error
}

// AddGameToCart implements models.CartRepository.
func (c *cartRepositoryImpl) AddGameToCart(userID int, gameID int, quantity int) error {
	cart, err := c.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create a new cart if not exists
			cart = &models.ShoppingCart{
				UserID: userID,
			}
			if err := c.db.DB.Create(cart).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Check if the game already exists in the cart
	var cartItem models.CartItem
	result := c.db.DB.Where("shopping_cart_id = ? AND game_id = ?", cart.ID, gameID).First(&cartItem)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Add new item to cart
			cartItem = models.CartItem{
				ShoppingCartID: cart.ID,
				GameID:         gameID,
				Quantity:       quantity,
			}
			return c.db.DB.Create(&cartItem).Error
		}
		return result.Error
	}

	// Update quantity of existing item
	cartItem.Quantity = quantity
	return c.db.DB.Save(&cartItem).Error
}

// ClearCart implements models.CartRepository.
func (c *cartRepositoryImpl) ClearCart(userID int) error {
	cart, err := c.FindByUserID(userID)
	if err != nil {
		return err
	}

	return c.db.DB.Where("shopping_cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error
}

// FindByUserID implements models.CartRepository.
func (c *cartRepositoryImpl) FindByUserID(userID int) (*models.ShoppingCart, error) {
	var cart models.ShoppingCart
	err := c.db.DB.Where("user_id = ?", userID).First(&cart).Error
	return &cart, err
}

// GetCartItems implements models.CartRepository.
func (c *cartRepositoryImpl) GetCartItems(userID int) ([]*models.CartItem, error) {
	cart, err := c.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var cartItems []*models.CartItem
	err = c.db.DB.Where("shopping_cart_id = ?", cart.ID).
		Preload("Game").
		Find(&cartItems).Error

	return cartItems, err
}

// RemoveGameFromCart implements models.CartRepository.
func (c *cartRepositoryImpl) RemoveGameFromCart(userID int, gameID int) error {
	cart, err := c.FindByUserID(userID)
	if err != nil {
		return err
	}

	return c.db.DB.Where("shopping_cart_id = ? AND game_id = ?", cart.ID, gameID).
		Delete(&models.CartItem{}).Error
}
