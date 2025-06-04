package models

import (
	"time"

	"gorm.io/gorm"
)

// Game represents a video game in the system
type Game struct {
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(255);not null" validate:"required"`
	Description string
	Price       float64 `gorm:"not null" validate:"required,gte=0"`
	ReleaseDate time.Time
	DeveloperID int        `gorm:"not null" validate:"required"`
	Developer   *Developer `gorm:"foreignKey:DeveloperID"`
	CategoryID  int        `gorm:"not null" validate:"required"`
	Category    *Category  `gorm:"foreignKey:CategoryID"`
	ImageData   []byte     `gorm:"type:bytea"`
	ImageName   string     `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Relations
	CartItems     []*CartItem
	FavoriteItems []*FavoriteItem
	LibraryItems  []*LibraryItem
	OrderItems    []*OrderItem
	Reviews       []*Review
	Restricts     []*Restrict
}

// Developer represents a game developer
type Developer struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null;unique" validate:"required"`
	Country     string
	Description string
	WebsiteURL  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Relations
	Games []*Game
}

// Category represents a game category/genre
type Category struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null;unique" validate:"required"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Relations
	Games []*Game
}

// Restrict represents regional restrictions for games
type Restrict struct {
	ID        int    `gorm:"primaryKey"`
	GameID    int    `gorm:"not null" validate:"required"`
	Game      *Game  `gorm:"foreignKey:GameID"`
	Region    string `gorm:"not null" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Review represents a user review for a game
type Review struct {
	ID          int    `gorm:"primaryKey"`
	GameID      int    `gorm:"not null" validate:"required"`
	Game        *Game  `gorm:"foreignKey:GameID"`
	UserID      int    `gorm:"not null" validate:"required"`
	User        *User  `gorm:"foreignKey:UserID"`
	Title       string `gorm:"not null" validate:"required"`
	Description string
	Rating      int `gorm:"not null" validate:"required,min=1,max=5"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// GameRepository defines the interface for game data access
type GameRepository interface {
	Create(game *Game) error
	FindByID(id int) (*Game, error)
	Update(game *Game) error
	Delete(id int) error
	FindAll(limit, offset int) ([]*Game, error)
	FindByCategory(categoryID int) ([]*Game, error)
	FindByDeveloper(developerID int) ([]*Game, error)
}

// DeveloperRepository defines the interface for developer data access
type DeveloperRepository interface {
	Create(developer *Developer) error
	FindByID(id int) (*Developer, error)
	Update(developer *Developer) error
	Delete(id int) error
	FindAll(limit, offset int) ([]*Developer, error)
}

// CategoryRepository defines the interface for category data access
type CategoryRepository interface {
	Create(category *Category) error
	FindByID(id int) (*Category, error)
	Update(category *Category) error
	Delete(id int) error
	FindAll(limit, offset int) ([]*Category, error)
}

// RestrictRepository defines the interface for restrict data access
type RestrictRepository interface {
	Create(restrict *Restrict) error
	FindByID(id int) (*Restrict, error)
	FindByGameID(gameID int) ([]*Restrict, error)
	Update(restrict *Restrict) error
	Delete(id int) error
	FindAll(limit, offset int) ([]*Restrict, error)
}

// ReviewRepository defines the interface for review data access
type ReviewRepository interface {
	Create(review *Review) error
	FindByID(id int) (*Review, error)
	FindByGameID(gameID int) ([]*Review, error)
	FindByUserID(userID int) ([]*Review, error)
	Update(review *Review) error
	Delete(id int) error
}
