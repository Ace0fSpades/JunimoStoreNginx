package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID           int    `gorm:"primaryKey"`
	Nickname     string `gorm:"not null;unique" validate:"required,min=2,max=100"`
	Email        string `gorm:"unique;not null" validate:"required,email"`
	Password     string `gorm:"not null" validate:"required,min=6"`
	RoleID       int    `gorm:"not null;default:2"`
	Role         *Role  `gorm:"foreignKey:RoleID"`
	Points       int    `gorm:"default:0"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Token        string
	RefreshToken string

	// Relations
	ShoppingCart *ShoppingCart
	Favorites    *Favorite
	Library      *Library
	Orders       []*Order
	Reviews      []*Review
}

// Role represents a user role in the system
type Role struct {
	ID          int    `gorm:"primaryKey"`
	Type        string `gorm:"not null;unique"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Relations
	Users []*User
}

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(user *User) error
	FindByID(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByNickname(nickname string) (*User, error)
	Update(user *User) error
	Delete(id int) error
	FindAll(limit, offset int) ([]*User, error)
}

// RoleRepository defines the interface for role data access
type RoleRepository interface {
	Create(role *Role) error
	FindByID(id int) (*Role, error)
	FindByType(roleType string) (*Role, error)
	Update(role *Role) error
	Delete(id int) error
	FindAll() ([]*Role, error)
}
