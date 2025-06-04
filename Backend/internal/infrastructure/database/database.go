package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/utils"
)

// Database represents the database connection and operations
type Database struct {
	DB *gorm.DB
}

// NewDatabase creates a new database connection
func NewDatabase() (*Database, error) {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// Retrieve connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Default values if not provided
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbUser == "" {
		dbUser = "postgres"
	}
	if dbName == "" {
		dbName = "unistore"
	}

	// Create connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Use singular table names
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Database{DB: db}, nil
}

// Migrate performs database migrations
func (d *Database) Migrate() error {
	// Check if tables already exist in the database
	var tableCount int64
	d.DB.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public'").Count(&tableCount)

	// If the database is empty, use clean migration
	if tableCount == 0 {
		log.Println("Performing clean database migration...")
		return d.cleanMigrate()
	}

	// Otherwise perform regular migration
	log.Println("Performing database migration with existing tables...")
	return d.regularMigrate()
}

// cleanMigrate performs migration on a clean database and fills it with mock data
func (d *Database) cleanMigrate() error {
	// Set migration settings
	migrator := d.DB.Session(&gorm.Session{
		SkipDefaultTransaction: true,
		AllowGlobalUpdate:      true,
	})

	// Migrate models in the correct order
	modelGroups := [][]interface{}{
		// Base models
		{&models.Role{}, &models.Developer{}, &models.Category{}},
		// Models with dependencies
		{&models.User{}, &models.Game{}},
		// Relationship models
		{&models.ShoppingCart{}, &models.Favorite{}, &models.Library{}, &models.Order{}, &models.Restrict{}, &models.Review{}},
		// Join tables
		{&models.CartItem{}, &models.FavoriteItem{}, &models.LibraryItem{}, &models.OrderItem{}},
	}

	for i, group := range modelGroups {
		for _, model := range group {
			if err := migrator.AutoMigrate(model); err != nil {
				return fmt.Errorf("failed to migrate group %d: %w", i, err)
			}
		}
	}

	return nil
}

// regularMigrate performs regular migration on an existing database
func (d *Database) regularMigrate() error {
	// Set migration settings
	migrator := d.DB.Session(&gorm.Session{
		SkipDefaultTransaction: true,
		AllowGlobalUpdate:      true,
	})

	// Migrate models in the correct order
	modelGroups := [][]interface{}{
		// Base models
		{&models.Role{}, &models.Developer{}, &models.Category{}},
		// Models with dependencies
		{&models.User{}, &models.Game{}},
		// Relationship models
		{&models.ShoppingCart{}, &models.Favorite{}, &models.Library{}, &models.Order{}, &models.Restrict{}, &models.Review{}},
		// Join tables
		{&models.CartItem{}, &models.FavoriteItem{}, &models.LibraryItem{}, &models.OrderItem{}},
	}

	for i, group := range modelGroups {
		for _, model := range group {
			if err := migrator.AutoMigrate(model); err != nil {
				return fmt.Errorf("failed to migrate group %d: %w", i, err)
			}
		}
	}

	// After schema migration, add missing data from mocks
	return d.updateDataFromMocks()
}

// updateDataFromMocks updates the database data, adding new records from mocks
func (d *Database) updateDataFromMocks() error {
	// Get data from mocks
	mockData := GetMockData()

	return d.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Add missing roles
		for _, role := range mockData.Roles {
			var existingRole models.Role
			if err := tx.Where("type = ?", role.Type).First(&existingRole).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// Role doesn't exist, create a new one
					if err := tx.Create(&role).Error; err != nil {
						return fmt.Errorf("failed to create role %s: %w", role.Type, err)
					}
					log.Printf("Created new role: %s", role.Type)
				} else {
					return fmt.Errorf("error checking role %s: %w", role.Type, err)
				}
			}
		}

		// 2. Add missing categories
		for _, category := range mockData.Categories {
			var existingCategory models.Category
			if err := tx.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// Category doesn't exist, create a new one
					if err := tx.Create(&category).Error; err != nil {
						return fmt.Errorf("failed to create category %s: %w", category.Name, err)
					}
					log.Printf("Created new category: %s", category.Name)
				} else {
					return fmt.Errorf("error checking category %s: %w", category.Name, err)
				}
			}
		}

		// 3. Add missing developers
		for _, developer := range mockData.Developers {
			var existingDeveloper models.Developer
			if err := tx.Where("name = ?", developer.Name).First(&existingDeveloper).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// Developer doesn't exist, create a new one
					if err := tx.Create(&developer).Error; err != nil {
						return fmt.Errorf("failed to create developer %s: %w", developer.Name, err)
					}
					log.Printf("Created new developer: %s", developer.Name)
				} else {
					return fmt.Errorf("error checking developer %s: %w", developer.Name, err)
				}
			}
		}

		// 4. Update games
		for _, game := range mockData.Games {
			var existingGame models.Game
			if err := tx.Where("title = ?", game.Title).First(&existingGame).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// Get correct IDs for category and developer
					var category models.Category
					if err := tx.Where("name = ?", game.Category.Name).First(&category).Error; err != nil {
						return fmt.Errorf("failed to find category for game %s: %w", game.Title, err)
					}

					var developer models.Developer
					if err := tx.Where("name = ?", game.Developer.Name).First(&developer).Error; err != nil {
						return fmt.Errorf("failed to find developer for game %s: %w", game.Title, err)
					}

					// Set correct IDs
					game.CategoryID = category.ID
					game.DeveloperID = developer.ID

					// Game doesn't exist, create a new one
					if err := tx.Create(&game).Error; err != nil {
						return fmt.Errorf("failed to create game %s: %w", game.Title, err)
					}
					log.Printf("Created new game: %s", game.Title)
				} else {
					return fmt.Errorf("error checking game %s: %w", game.Title, err)
				}
			}
		}

		// Don't add users, carts, favorites, libraries, orders, and reviews,
		// as they can be created by users dynamically
		// and must maintain data consistency

		return nil
	})
}

// CheckAdminAndRoles ensures default roles exist
func (d *Database) CheckAdminAndRoles() error {
	// Use transaction for ensuring data consistency
	return d.DB.Transaction(func(tx *gorm.DB) error {
		// Create default roles if they don't exist
		var adminRole models.Role
		if err := tx.First(&adminRole, "type = ?", "admin").Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				adminRole = models.Role{
					Type:        "admin",
					Description: "Administrator with full access",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				if err := tx.Create(&adminRole).Error; err != nil {
					return fmt.Errorf("failed to create admin role: %w", err)
				}
			} else {
				return fmt.Errorf("error checking admin role: %w", err)
			}
		}

		var userRole models.Role
		if err := tx.First(&userRole, "type = ?", "user").Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				userRole = models.Role{
					Type:        "user",
					Description: "Regular user with limited access",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				if err := tx.Create(&userRole).Error; err != nil {
					return fmt.Errorf("failed to create user role: %w", err)
				}
			} else {
				return fmt.Errorf("error checking user role: %w", err)
			}
		}

		// Check if admin user exists
		var adminCount int64
		if err := tx.Model(&models.User{}).Where("role_id = ?", adminRole.ID).Count(&adminCount).Error; err != nil {
			return fmt.Errorf("failed to count admin users: %w", err)
		}

		// Create default admin if none exists
		if adminCount == 0 {
			adminPassword := os.Getenv("ADMIN_PASSWORD")
			if adminPassword == "" {
				adminPassword = "admin123" // Default password if not provided
			}

			hashedPassword, err := utils.HashPassword(adminPassword)
			if err != nil {
				return fmt.Errorf("failed to hash admin password: %w", err)
			}

			adminUser := models.User{
				Nickname:  "AdminUser",
				Email:     "admin@example.com",
				Password:  hashedPassword,
				RoleID:    adminRole.ID,
				Points:    0,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			if err := tx.Create(&adminUser).Error; err != nil {
				return fmt.Errorf("failed to create admin user: %w", err)
			}

			// Create shopping cart for admin
			cart := models.ShoppingCart{
				UserID:    adminUser.ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := tx.Create(&cart).Error; err != nil {
				return fmt.Errorf("failed to create admin shopping cart: %w", err)
			}

			// Create favorite list for admin
			favorite := models.Favorite{
				UserID:    adminUser.ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := tx.Create(&favorite).Error; err != nil {
				return fmt.Errorf("failed to create admin favorite: %w", err)
			}

			// Create library for admin
			library := models.Library{
				UserID:    adminUser.ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := tx.Create(&library).Error; err != nil {
				return fmt.Errorf("failed to create admin library: %w", err)
			}
		}

		// Ensure we have at least one basic category
		var categoryCount int64
		if err := tx.Model(&models.Category{}).Count(&categoryCount).Error; err != nil {
			return fmt.Errorf("failed to count categories: %w", err)
		}

		var actionCategory models.Category
		var defaultCategories []models.Category

		if categoryCount == 0 {
			defaultCategories = []models.Category{
				{
					Name:        "Action",
					Description: "Action games focus on challenging the player's reflexes, hand-eye coordination, and reaction time",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					Name:        "Adventure",
					Description: "Adventure games focus on exploration, puzzle-solving, and narrative",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					Name:        "RPG",
					Description: "Role-playing games where players assume the roles of characters in a fictional setting",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					Name:        "Strategy",
					Description: "Strategy games focus on skillful thinking and planning to achieve victory",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			}

			for _, category := range defaultCategories {
				if err := tx.Create(&category).Error; err != nil {
					return fmt.Errorf("failed to create default category %s: %w", category.Name, err)
				}

				if category.Name == "Action" {
					actionCategory = category
				}
			}
		} else {
			// Find Action category for demo game
			if err := tx.Where("name = ?", "Action").First(&actionCategory).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// If Action category doesn't exist, use the first category
					if err := tx.First(&actionCategory).Error; err != nil {
						return fmt.Errorf("failed to find any category: %w", err)
					}
				} else {
					return fmt.Errorf("error finding Action category: %w", err)
				}
			}
		}

		// Ensure we have at least one developer
		var developerCount int64
		if err := tx.Model(&models.Developer{}).Count(&developerCount).Error; err != nil {
			return fmt.Errorf("failed to count developers: %w", err)
		}

		var defaultDeveloper models.Developer

		if developerCount == 0 {
			defaultDeveloper = models.Developer{
				Name:        "UniStore Games",
				Country:     "United States",
				Description: "Default game developer for the UniStore platform",
				WebsiteURL:  "https://unistore.example.com",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			if err := tx.Create(&defaultDeveloper).Error; err != nil {
				return fmt.Errorf("failed to create default developer: %w", err)
			}
		} else {
			// Get first developer for demo game
			if err := tx.First(&defaultDeveloper).Error; err != nil {
				return fmt.Errorf("failed to get developer: %w", err)
			}
		}

		// Check if we have any games
		var gameCount int64
		if err := tx.Model(&models.Game{}).Count(&gameCount).Error; err != nil {
			return fmt.Errorf("failed to count games: %w", err)
		}

		if gameCount == 0 {
			// Create a demo game
			demoGame := models.Game{
				Title:       "Demo Game",
				Description: "This is a demo game for testing purposes",
				Price:       29.99,
				ReleaseDate: time.Now(),
				DeveloperID: defaultDeveloper.ID,
				CategoryID:  actionCategory.ID,
				ImageData:   []byte{}, // Пустой массив байт для изображения
				ImageName:   "gameBlankImage.png",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			if err := tx.Create(&demoGame).Error; err != nil {
				return fmt.Errorf("failed to create demo game: %w", err)
			}
		}

		return nil
	})
}
