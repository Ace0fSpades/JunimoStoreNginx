package models

import (
	"time"

	"gorm.io/gorm"
)

// Order represents a user's order
type Order struct {
	ID        int     `gorm:"primaryKey"`
	UserID    int     `gorm:"not null" validate:"required"`
	User      *User   `gorm:"foreignKey:UserID"`
	TotalCost float64 `gorm:"not null" validate:"required,gte=0"`
	Status    string  `gorm:"not null;default:'pending'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relations
	OrderItems []*OrderItem
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        int     `gorm:"primaryKey"`
	OrderID   int     `gorm:"not null" validate:"required"`
	Order     *Order  `gorm:"foreignKey:OrderID"`
	GameID    int     `gorm:"not null" validate:"required"`
	Game      *Game   `gorm:"foreignKey:GameID"`
	Price     float64 `gorm:"not null" validate:"required,gte=0"`
	Quantity  int     `gorm:"not null;default:1" validate:"required,min=1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// OrderRepository defines the interface for order data access
type OrderRepository interface {
	Create(order *Order) error
	FindByID(id int) (*Order, error)
	FindByUserID(userID int) ([]*Order, error)
	Update(order *Order) error
	FindAll(limit, offset int) ([]*Order, error)
	CreateFromCart(userID int) (*Order, error)
}
