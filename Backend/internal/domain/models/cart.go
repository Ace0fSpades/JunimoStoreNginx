package models

import (
	"time"

	"gorm.io/gorm"
)

// ShoppingCart represents a user's shopping cart
type ShoppingCart struct {
	ID        int   `gorm:"primaryKey"`
	UserID    int   `gorm:"not null;uniqueIndex" validate:"required"`
	User      *User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relations
	CartItems []*CartItem
}

// CartItem represents an item in a shopping cart
type CartItem struct {
	ID             int           `gorm:"primaryKey"`
	ShoppingCartID int           `gorm:"not null" validate:"required"`
	ShoppingCart   *ShoppingCart `gorm:"foreignKey:ShoppingCartID"`
	GameID         int           `gorm:"not null" validate:"required"`
	Game           *Game         `gorm:"foreignKey:GameID"`
	Quantity       int           `gorm:"not null;default:1" validate:"required,min=1"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

// Favorite represents a user's favorite games list
type Favorite struct {
	ID        int   `gorm:"primaryKey"`
	UserID    int   `gorm:"not null;uniqueIndex" validate:"required"`
	User      *User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relations
	FavoriteItems []*FavoriteItem
}

// FavoriteItem represents a game in a user's favorite list
type FavoriteItem struct {
	ID         int       `gorm:"primaryKey"`
	FavoriteID int       `gorm:"not null" validate:"required"`
	Favorite   *Favorite `gorm:"foreignKey:FavoriteID"`
	GameID     int       `gorm:"not null" validate:"required"`
	Game       *Game     `gorm:"foreignKey:GameID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// Library represents a user's game library
type Library struct {
	ID        int   `gorm:"primaryKey"`
	UserID    int   `gorm:"not null;uniqueIndex" validate:"required"`
	User      *User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relations
	LibraryItems []*LibraryItem
}

// LibraryItem represents a game in a user's library
type LibraryItem struct {
	ID        int      `gorm:"primaryKey"`
	LibraryID int      `gorm:"not null" validate:"required"`
	Library   *Library `gorm:"foreignKey:LibraryID"`
	GameID    int      `gorm:"not null" validate:"required"`
	Game      *Game    `gorm:"foreignKey:GameID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CartRepository defines the interface for cart data access
type CartRepository interface {
	Create(cart *ShoppingCart) error
	FindByUserID(userID int) (*ShoppingCart, error)
	AddGameToCart(userID, gameID int, quantity int) error
	RemoveGameFromCart(userID, gameID int) error
	ClearCart(userID int) error
	GetCartItems(userID int) ([]*CartItem, error)
}

// FavoriteRepository defines the interface for favorite data access
type FavoriteRepository interface {
	Create(favorite *Favorite) error
	FindByUserID(userID int) (*Favorite, error)
	AddGameToFavorite(userID, gameID int) error
	RemoveGameFromFavorite(userID, gameID int) error
	ClearFavorite(userID int) error
	GetFavoriteItems(userID int) ([]*FavoriteItem, error)
}

// LibraryRepository defines the interface for library data access
type LibraryRepository interface {
	Create(library *Library) error
	FindByUserID(userID int) (*Library, error)
	AddGameToLibrary(userID, gameID int) error
	GetLibraryItems(userID int) ([]*LibraryItem, error)
}
