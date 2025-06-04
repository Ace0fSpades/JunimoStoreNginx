package database

import (
	"log"
	"os"
	"time"

	"uniStore/Backend/internal/domain/models"
)

// MockData contains data for populating the database
type MockData struct {
	Roles         []models.Role
	Users         []models.User
	Categories    []models.Category
	Developers    []models.Developer
	Games         []models.Game
	ShoppingCarts []models.ShoppingCart
	CartItems     []models.CartItem
	Favorites     []models.Favorite
	FavoriteItems []models.FavoriteItem
	Libraries     []models.Library
	LibraryItems  []models.LibraryItem
	Orders        []models.Order
	OrderItems    []models.OrderItem
	Reviews       []models.Review
}

// GetMockData returns test data for the database
func GetMockData() MockData {
	now := time.Now()

	// Load default image
	defaultImageData, err := os.ReadFile("internal/assets/images/gameBlankImage.png")
	if err != nil {
		log.Printf("Error loading default image: %v", err)
		defaultImageData = []byte{} // Empty byte array if image loading fails
	}

	// Roles
	adminRole := models.Role{
		Type:        "admin",
		Description: "Administrator with full access",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	userRole := models.Role{
		Type:        "user",
		Description: "Regular user with limited access",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Categories
	actionCategory := models.Category{
		Name:        "Action",
		Description: "Action games focus on challenging the player's reflexes, hand-eye coordination, and reaction time",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	adventureCategory := models.Category{
		Name:        "Adventure",
		Description: "Adventure games focus on exploration, puzzle-solving, and narrative",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	rpgCategory := models.Category{
		Name:        "RPG",
		Description: "Role-playing games where players assume the roles of characters in a fictional setting",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	strategyCategory := models.Category{
		Name:        "Strategy",
		Description: "Strategy games focus on skillful thinking and planning to achieve victory",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Add new categories
	simulationCategory := models.Category{
		Name:        "Simulation",
		Description: "Games designed to simulate real-world activities or fictional scenarios",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	sportsCategory := models.Category{
		Name:        "Sports",
		Description: "Games that simulate traditional sports such as football, basketball, or racing",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	puzzleCategory := models.Category{
		Name:        "Puzzle",
		Description: "Games that emphasize puzzle solving, logic, pattern recognition, and problem solving",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	horrorCategory := models.Category{
		Name:        "Horror",
		Description: "Games designed to scare players through atmosphere, sound design, and psychological terror",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	openWorldCategory := models.Category{
		Name:        "Open World",
		Description: "Games featuring a virtual world in which the player can explore freely",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Developers
	uniStoreDev := models.Developer{
		Name:        "UniStore Games",
		Country:     "United States",
		Description: "Default game developer for the UniStore platform",
		WebsiteURL:  "https://unistore.example.com",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	indie := models.Developer{
		Name:        "Indie Studio",
		Country:     "Canada",
		Description: "Independent game studio creating innovative games",
		WebsiteURL:  "https://indie.example.com",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Well-known developers for realistic data
	valveDev := models.Developer{
		Name:        "Valve Corporation",
		Country:     "United States",
		Description: "Developer of the Half-Life series, Portal, Counter-Strike, and Steam platform",
		WebsiteURL:  "https://www.valvesoftware.com",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	cdProjektDev := models.Developer{
		Name:        "CD Projekt Red",
		Country:     "Poland",
		Description: "Developer of The Witcher series and Cyberpunk 2077",
		WebsiteURL:  "https://www.cdprojektred.com",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Add new developers
	rockstarDev := models.Developer{
		Name:        "Rockstar Games",
		Country:     "United States",
		Description: "Developer of Grand Theft Auto and Red Dead Redemption series",
		WebsiteURL:  "https://www.rockstargames.com",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	ubisoftDev := models.Developer{
		Name:        "Ubisoft",
		Country:     "France",
		Description: "Developer of Assassin's Creed, Far Cry, and Watch Dogs series",
		WebsiteURL:  "https://www.ubisoft.com",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	eaDev := models.Developer{
		Name:        "Electronic Arts",
		Country:     "United States",
		Description: "Developer of FIFA, Battlefield, and The Sims series",
		WebsiteURL:  "https://www.ea.com",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	mojangDev := models.Developer{
		Name:        "Mojang Studios",
		Country:     "Sweden",
		Description: "Developer of Minecraft",
		WebsiteURL:  "https://www.minecraft.net",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	fromSoftwareDev := models.Developer{
		Name:        "FromSoftware",
		Country:     "Japan",
		Description: "Developer of Dark Souls, Bloodborne, and Elden Ring",
		WebsiteURL:  "https://www.fromsoftware.jp",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Users
	adminUser := models.User{
		Nickname:     "AdminUser",
		Email:        "admin@example.com",
		Password:     "admin123", // Will be hashed before saving
		RoleID:       1,          // Admin role ID
		Points:       1000,
		CreatedAt:    now,
		UpdatedAt:    now,
		Token:        "admin-jwt-token",
		RefreshToken: "admin-refresh-token",
	}

	regularUser := models.User{
		Nickname:     "UserGamer",
		Email:        "user@example.com",
		Password:     "user123", // Will be hashed before saving
		RoleID:       2,         // User role ID
		Points:       500,
		CreatedAt:    now,
		UpdatedAt:    now,
		Token:        "user-jwt-token",
		RefreshToken: "user-refresh-token",
	}

	// Set IDs before creating relationships
	actionCategory.ID = 1
	adventureCategory.ID = 2
	rpgCategory.ID = 3
	strategyCategory.ID = 4
	simulationCategory.ID = 5
	sportsCategory.ID = 6
	puzzleCategory.ID = 7
	horrorCategory.ID = 8
	openWorldCategory.ID = 9

	uniStoreDev.ID = 1
	indie.ID = 2
	valveDev.ID = 3
	cdProjektDev.ID = 4
	rockstarDev.ID = 5
	ubisoftDev.ID = 6
	eaDev.ID = 7
	mojangDev.ID = 8
	fromSoftwareDev.ID = 9

	// Set user IDs
	adminUser.ID = 1
	regularUser.ID = 2

	// Games with correctly established relationships
	demoGame := models.Game{
		Title:       "Demo Game",
		Description: "This is a demo game for testing purposes",
		Price:       29.99,
		ReleaseDate: now,
		DeveloperID: 1,               // UniStore developer ID
		Developer:   &uniStoreDev,    // Related object
		CategoryID:  1,               // Action category ID
		Category:    &actionCategory, // Related object
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	adventureGame := models.Game{
		Title:       "Adventure Quest",
		Description: "Embark on an epic journey to save the world",
		Price:       39.99,
		ReleaseDate: now,
		DeveloperID: 2,                  // Indie developer ID
		Developer:   &indie,             // Related object
		CategoryID:  2,                  // Adventure category ID
		Category:    &adventureCategory, // Related object
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	rpgGame := models.Game{
		Title:       "Fantasy World",
		Description: "Immerse yourself in a rich fantasy world with deep lore and compelling characters",
		Price:       49.99,
		ReleaseDate: now.AddDate(0, -3, 0), // 3 months ago
		DeveloperID: 3,                     // ID developer CD Projekt Red
		Developer:   &valveDev,             // Related object
		CategoryID:  3,                     // ID category RPG
		Category:    &rpgCategory,          // Related object
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	strategyGame := models.Game{
		Title:       "Empire Builder",
		Description: "Build and manage your own empire through the ages",
		Price:       34.99,
		ReleaseDate: now.AddDate(0, -1, 0), // 1 month ago
		DeveloperID: 2,                     // ID developer Indie
		Developer:   &indie,                // Related object
		CategoryID:  4,                     // ID category Strategy
		Category:    &strategyCategory,     // Related object
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	shooterGame := models.Game{
		Title:       "Tactical Force",
		Description: "Fast-paced first-person shooter with tactical elements",
		Price:       19.99,
		ReleaseDate: now.AddDate(0, -6, 0), // 6 months ago
		DeveloperID: 1,                     // ID developer UniStore
		Developer:   &uniStoreDev,          // Related object
		CategoryID:  1,                     // ID category Action
		Category:    &actionCategory,       // Related object
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Add new games based on real ones
	witcher3Game := models.Game{
		Title:       "The Witcher 3: Wild Hunt",
		Description: "An open-world RPG set in a dark fantasy universe, following the adventures of monster hunter Geralt of Rivia",
		Price:       39.99,
		ReleaseDate: time.Date(2015, 5, 19, 0, 0, 0, 0, time.UTC),
		DeveloperID: 4, // CD Projekt Red
		Developer:   &cdProjektDev,
		CategoryID:  3, // RPG
		Category:    &rpgCategory,
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// ID: 7
	gtaVGame := models.Game{
		Title:       "Grand Theft Auto V",
		Description: "An action-adventure game set in the fictional state of San Andreas, following three criminals and their efforts to commit heists",
		Price:       29.99,
		ReleaseDate: time.Date(2013, 9, 17, 0, 0, 0, 0, time.UTC),
		DeveloperID: 5, // Rockstar Games
		Developer:   &rockstarDev,
		CategoryID:  9, // Open World
		Category:    &openWorldCategory,
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// ID: 8
	minecraftGame := models.Game{
		Title:       "Minecraft",
		Description: "A sandbox game focused on exploration, building, and survival in a procedurally generated 3D world",
		Price:       26.99,
		ReleaseDate: time.Date(2011, 11, 18, 0, 0, 0, 0, time.UTC),
		DeveloperID: 8, // Mojang Studios
		Developer:   &mojangDev,
		CategoryID:  5, // Simulation
		Category:    &simulationCategory,
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// ID: 9
	csgoGame := models.Game{
		Title:       "Counter-Strike: Global Offensive",
		Description: "A competitive first-person shooter pitting two teams against each other: Terrorists and Counter-Terrorists",
		Price:       0.00, // Free to play
		ReleaseDate: time.Date(2012, 8, 21, 0, 0, 0, 0, time.UTC),
		DeveloperID: 3, // Valve Corporation
		Developer:   &valveDev,
		CategoryID:  1, // Action
		Category:    &actionCategory,
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// ID: 10
	fifa23Game := models.Game{
		Title:       "FIFA 23",
		Description: "The latest installment in the FIFA series, featuring realistic football gameplay and official teams",
		Price:       59.99,
		ReleaseDate: time.Date(2022, 9, 30, 0, 0, 0, 0, time.UTC),
		DeveloperID: 7, // Electronic Arts
		Developer:   &eaDev,
		CategoryID:  6, // Sports
		Category:    &sportsCategory,
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// ID: 11
	redDeadRedemption2Game := models.Game{
		Title:       "Red Dead Redemption 2",
		Description: "An epic tale of life in America's unforgiving heartland, following outlaw Arthur Morgan and the Van der Linde gang",
		Price:       59.99,
		ReleaseDate: time.Date(2018, 10, 26, 0, 0, 0, 0, time.UTC),
		DeveloperID: 5, // Rockstar Games
		Developer:   &rockstarDev,
		CategoryID:  9, // Open World
		Category:    &openWorldCategory,
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// ID: 12
	cyberpunk2077Game := models.Game{
		Title:       "Cyberpunk 2077",
		Description: "An open-world, action-adventure story set in Night City, a megalopolis obsessed with power, glamour, and body modification",
		Price:       49.99,
		ReleaseDate: time.Date(2020, 12, 10, 0, 0, 0, 0, time.UTC),
		DeveloperID: 4, // CD Projekt Red
		Developer:   &cdProjektDev,
		CategoryID:  9, // Open World
		Category:    &openWorldCategory,
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// ID: 13
	acValhalla := models.Game{
		Title:       "Assassin's Creed Valhalla",
		Description: "Become a legendary Viking warrior on a quest for glory, exploring England's Dark Ages",
		Price:       59.99,
		ReleaseDate: time.Date(2020, 11, 10, 0, 0, 0, 0, time.UTC),
		DeveloperID: 6, // Ubisoft
		Developer:   &ubisoftDev,
		CategoryID:  9, // Open World
		Category:    &openWorldCategory,
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// ID: 14
	eldenRingGame := models.Game{
		Title:       "Elden Ring",
		Description: "An action RPG set in a fantasy world created by Hidetaka Miyazaki and George R. R. Martin",
		Price:       59.99,
		ReleaseDate: time.Date(2022, 2, 25, 0, 0, 0, 0, time.UTC),
		DeveloperID: 9, // FromSoftware
		Developer:   &fromSoftwareDev,
		CategoryID:  3, // RPG
		Category:    &rpgCategory,
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// ID: 15
	portal2Game := models.Game{
		Title:       "Portal 2",
		Description: "A first-person puzzle-platform game featuring cooperative gameplay and mind-bending physics",
		Price:       19.99,
		ReleaseDate: time.Date(2011, 4, 19, 0, 0, 0, 0, time.UTC),
		DeveloperID: 3, // Valve Corporation
		Developer:   &valveDev,
		CategoryID:  7, // Puzzle
		Category:    &puzzleCategory,
		ImageData:   defaultImageData,
		ImageName:   "gameBlankImage.png",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Cart for the first user
	userCart := models.ShoppingCart{
		UserID:    2, // Regular user ID
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Cart items
	cartItem1 := models.CartItem{
		ShoppingCartID: 1,         // User's cart ID
		GameID:         1,         // Demo Game ID
		Game:           &demoGame, // Related object
		Quantity:       1,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	cartItem2 := models.CartItem{
		ShoppingCartID: 1,        // User's cart ID
		GameID:         3,        // Fantasy World ID
		Game:           &rpgGame, // Related object
		Quantity:       2,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	// Favorite for the user
	userFavorite := models.Favorite{
		UserID:    2, // Regular user ID
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Favorite items
	favoriteItem1 := models.FavoriteItem{
		FavoriteID: 1,              // User's favorites ID
		GameID:     2,              // Adventure Quest ID
		Game:       &adventureGame, // Related object
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	favoriteItem2 := models.FavoriteItem{
		FavoriteID: 1,             // User's favorites ID
		GameID:     4,             // Empire Builder ID
		Game:       &strategyGame, // Related object
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	// Library for the user
	userLibrary := models.Library{
		UserID:    2, // Regular user ID
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Library items
	libraryItem1 := models.LibraryItem{
		LibraryID: 1,            // User's library ID
		GameID:    5,            // Tactical Force ID
		Game:      &shooterGame, // Related object
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Orders
	completedOrder := models.Order{
		UserID:    2, // Regular user ID
		TotalCost: 49.99,
		Status:    "completed",
		CreatedAt: now.AddDate(0, -1, 0), // 1 month ago
		UpdatedAt: now.AddDate(0, -1, 0),
	}

	pendingOrder := models.Order{
		UserID:    2, // Regular user ID
		TotalCost: 39.99,
		Status:    "pending",
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Order items
	orderItem1 := models.OrderItem{
		OrderID:   1,              // Completed order ID
		GameID:    2,              // Adventure Quest ID
		Game:      &adventureGame, // Related object
		Price:     39.99,
		Quantity:  1,
		CreatedAt: now.AddDate(0, -1, 0),
		UpdatedAt: now.AddDate(0, -1, 0),
	}

	orderItem2 := models.OrderItem{
		OrderID:   1,            // Completed order ID
		GameID:    5,            // Tactical Force ID
		Game:      &shooterGame, // Related object
		Price:     19.99,
		Quantity:  1,
		CreatedAt: now.AddDate(0, -1, 0),
		UpdatedAt: now.AddDate(0, -1, 0),
	}

	orderItem3 := models.OrderItem{
		OrderID:   2,        // Pending order ID
		GameID:    3,        // Fantasy World ID
		Game:      &rpgGame, // Related object
		Price:     49.99,
		Quantity:  1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Reviews
	review1 := models.Review{
		GameID:      5,            // ID for Tactical Force game
		Game:        &shooterGame, // Related object
		UserID:      2,            // ID for regular user
		User:        &regularUser, // Related object
		Title:       "Great game!",
		Description: "Very engaging shooter with good graphics and gameplay.",
		Rating:      5,
		CreatedAt:   now.AddDate(0, 0, -15), // 15 days ago
		UpdatedAt:   now.AddDate(0, 0, -15),
	}

	review2 := models.Review{
		GameID:      2,              // ID for Adventure Quest game
		Game:        &adventureGame, // Related object
		UserID:      2,              // ID for regular user
		User:        &regularUser,   // Related object
		Title:       "Nice adventure game",
		Description: "Interesting plot, but a bit boring at times. Overall positive impression.",
		Rating:      4,
		CreatedAt:   now.AddDate(0, 0, -7), // 7 days ago
		UpdatedAt:   now.AddDate(0, 0, -7),
	}

	return MockData{
		Roles: []models.Role{
			adminRole,
			userRole,
		},
		Categories: []models.Category{
			actionCategory,
			adventureCategory,
			rpgCategory,
			strategyCategory,
			simulationCategory,
			sportsCategory,
			puzzleCategory,
			horrorCategory,
			openWorldCategory,
		},
		Developers: []models.Developer{
			uniStoreDev,
			indie,
			valveDev,
			cdProjektDev,
			rockstarDev,
			ubisoftDev,
			eaDev,
			mojangDev,
			fromSoftwareDev,
		},
		Games: []models.Game{
			demoGame,
			adventureGame,
			rpgGame,
			strategyGame,
			shooterGame,
			witcher3Game,
			gtaVGame,
			minecraftGame,
			csgoGame,
			fifa23Game,
			redDeadRedemption2Game,
			cyberpunk2077Game,
			acValhalla,
			eldenRingGame,
			portal2Game,
		},
		Users: []models.User{
			adminUser,
			regularUser,
		},
		ShoppingCarts: []models.ShoppingCart{
			userCart,
		},
		CartItems: []models.CartItem{
			cartItem1,
			cartItem2,
		},
		Favorites: []models.Favorite{
			userFavorite,
		},
		FavoriteItems: []models.FavoriteItem{
			favoriteItem1,
			favoriteItem2,
		},
		Libraries: []models.Library{
			userLibrary,
		},
		LibraryItems: []models.LibraryItem{
			libraryItem1,
		},
		Orders: []models.Order{
			completedOrder,
			pendingOrder,
		},
		OrderItems: []models.OrderItem{
			orderItem1,
			orderItem2,
			orderItem3,
		},
		Reviews: []models.Review{
			review1,
			review2,
		},
	}
}
