package services

import (
	"errors"
	"strings"
	"time"

	"uniStore/Backend/internal/domain/models"
	"uniStore/Backend/internal/interfaces/dto"
)

// GameServiceImpl implements GameService interface
type GameServiceImpl struct {
	gameRepo      models.GameRepository
	categoryRepo  models.CategoryRepository
	developerRepo models.DeveloperRepository
	restrictRepo  models.RestrictRepository
}

// NewGameService creates a new game service
func NewGameService(gameRepo models.GameRepository, categoryRepo models.CategoryRepository, developerRepo models.DeveloperRepository, restrictRepo models.RestrictRepository) GameService {
	return &GameServiceImpl{
		gameRepo:      gameRepo,
		categoryRepo:  categoryRepo,
		developerRepo: developerRepo,
		restrictRepo:  restrictRepo,
	}
}

// CreateGame creates a new game
func (s *GameServiceImpl) CreateGame(gameDTO *dto.GameCreateDTO) (*dto.GameDTO, error) {
	// Convert DTO to model
	game := gameDTO.ToModel()

	// Set timestamps
	game.CreatedAt = time.Now()
	game.UpdatedAt = time.Now()

	// Load related models for validation if needed
	if game.DeveloperID != 0 {
		developer, err := s.developerRepo.FindByID(game.DeveloperID)
		if err != nil {
			return nil, err
		}
		game.Developer = developer
	}

	if game.CategoryID != 0 {
		category, err := s.categoryRepo.FindByID(game.CategoryID)
		if err != nil {
			return nil, err
		}
		game.Category = category
	}

	// Create game in repository
	if err := s.gameRepo.Create(game); err != nil {
		return nil, err
	}

	// Convert model back to DTO for response
	return dto.GameDTOFromModel(game), nil
}

// GetGameByID retrieves a game by ID
func (s *GameServiceImpl) GetGameByID(id int) (*dto.GameDTO, error) {
	// Get game from repository
	game, err := s.gameRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Load developer if not loaded
	if game.Developer == nil && game.DeveloperID != 0 {
		developer, err := s.developerRepo.FindByID(game.DeveloperID)
		if err != nil {
			return nil, err
		}
		game.Developer = developer
	}

	// Load category if not loaded
	if game.Category == nil && game.CategoryID != 0 {
		category, err := s.categoryRepo.FindByID(game.CategoryID)
		if err != nil {
			return nil, err
		}
		game.Category = category
	}

	// Load restrictions
	restrictions, err := s.restrictRepo.FindByGameID(game.ID)
	if err != nil {
		return nil, err
	}
	game.Restricts = restrictions

	// Convert model to DTO for response
	return dto.GameDTOFromModel(game), nil
}

// GetAllGames retrieves all games with pagination
func (s *GameServiceImpl) GetAllGames(limit, offset int) ([]*dto.GameDTO, error) {
	// Get games from repository
	games, err := s.gameRepo.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	// Load related data for all games
	for _, game := range games {
		// Load developer if not loaded
		if game.Developer == nil && game.DeveloperID != 0 {
			developer, err := s.developerRepo.FindByID(game.DeveloperID)
			if err != nil {
				return nil, err
			}
			game.Developer = developer
		}

		// Load category if not loaded
		if game.Category == nil && game.CategoryID != 0 {
			category, err := s.categoryRepo.FindByID(game.CategoryID)
			if err != nil {
				return nil, err
			}
			game.Category = category
		}

		// Load restrictions
		restrictions, err := s.restrictRepo.FindByGameID(game.ID)
		if err != nil {
			return nil, err
		}
		game.Restricts = restrictions
	}

	// Convert models to DTOs for response
	return dto.GameDTOsFromModels(games), nil
}

// UpdateGame updates a game
func (s *GameServiceImpl) UpdateGame(id int, gameDTO *dto.GameUpdateDTO) (*dto.GameDTO, error) {
	// Get existing game
	existingGame, err := s.gameRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert DTO to model
	updateData := gameDTO.ToUpdateModel(id)

	// Update fields if provided
	if updateData.Title != "" {
		existingGame.Title = updateData.Title
	}
	if updateData.Description != "" {
		existingGame.Description = updateData.Description
	}
	if updateData.Price > 0 {
		existingGame.Price = updateData.Price
	}
	if !updateData.ReleaseDate.IsZero() {
		existingGame.ReleaseDate = updateData.ReleaseDate
	}
	if updateData.DeveloperID != 0 {
		existingGame.DeveloperID = updateData.DeveloperID
	}
	if updateData.CategoryID != 0 {
		existingGame.CategoryID = updateData.CategoryID
	}
	if updateData.ImageData != nil && len(updateData.ImageData) > 0 {
		existingGame.ImageData = updateData.ImageData
	}
	if updateData.ImageName != "" {
		existingGame.ImageName = updateData.ImageName
	}

	// Update timestamp
	existingGame.UpdatedAt = time.Now()

	// Update game in repository
	if err := s.gameRepo.Update(existingGame); err != nil {
		return nil, err
	}

	// Load related data for response
	// Load developer if not loaded
	if existingGame.Developer == nil && existingGame.DeveloperID != 0 {
		developer, err := s.developerRepo.FindByID(existingGame.DeveloperID)
		if err != nil {
			return nil, err
		}
		existingGame.Developer = developer
	}

	// Load category if not loaded
	if existingGame.Category == nil && existingGame.CategoryID != 0 {
		category, err := s.categoryRepo.FindByID(existingGame.CategoryID)
		if err != nil {
			return nil, err
		}
		existingGame.Category = category
	}

	// Load restrictions
	restrictions, err := s.restrictRepo.FindByGameID(existingGame.ID)
	if err != nil {
		return nil, err
	}
	existingGame.Restricts = restrictions

	// Convert updated model to DTO for response
	return dto.GameDTOFromModel(existingGame), nil
}

// DeleteGame deletes a game
func (s *GameServiceImpl) DeleteGame(id int) error {
	return s.gameRepo.Delete(id)
}

// SearchGamesByTitle searches for games by title
func (s *GameServiceImpl) SearchGamesByTitle(title string, limit, offset int) ([]*dto.GameDTO, error) {
	if title == "" {
		return nil, errors.New("search title cannot be empty")
	}

	// Get all games from repository
	games, err := s.gameRepo.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	// Filter games by title
	var result []*models.Game
	lowerTitle := strings.ToLower(title)

	for _, game := range games {
		if strings.Contains(strings.ToLower(game.Title), lowerTitle) {
			result = append(result, game)
		}
	}

	// Load related data for all games
	for _, game := range result {
		// Load developer if not loaded
		if game.Developer == nil && game.DeveloperID != 0 {
			developer, err := s.developerRepo.FindByID(game.DeveloperID)
			if err != nil {
				return nil, err
			}
			game.Developer = developer
		}

		// Load category if not loaded
		if game.Category == nil && game.CategoryID != 0 {
			category, err := s.categoryRepo.FindByID(game.CategoryID)
			if err != nil {
				return nil, err
			}
			game.Category = category
		}

		// Load restrictions
		restrictions, err := s.restrictRepo.FindByGameID(game.ID)
		if err != nil {
			return nil, err
		}
		game.Restricts = restrictions
	}

	// Convert filtered models to DTOs for response
	return dto.GameDTOsFromModels(result), nil
}

// GetGamesByCategory retrieves games by category
func (s *GameServiceImpl) GetGamesByCategory(categoryID int) ([]*dto.GameDTO, error) {
	// Get games from repository
	games, err := s.gameRepo.FindByCategory(categoryID)
	if err != nil {
		return nil, err
	}

	// Load related data for all games
	for _, game := range games {
		// Load developer if not loaded
		if game.Developer == nil && game.DeveloperID != 0 {
			developer, err := s.developerRepo.FindByID(game.DeveloperID)
			if err != nil {
				return nil, err
			}
			game.Developer = developer
		}

		// Load category if not loaded
		if game.Category == nil && game.CategoryID != 0 {
			category, err := s.categoryRepo.FindByID(game.CategoryID)
			if err != nil {
				return nil, err
			}
			game.Category = category
		}

		// Load restrictions
		restrictions, err := s.restrictRepo.FindByGameID(game.ID)
		if err != nil {
			return nil, err
		}
		game.Restricts = restrictions
	}

	// Convert models to DTOs for response
	return dto.GameDTOsFromModels(games), nil
}

// GetGamesByDeveloper retrieves games by developer
func (s *GameServiceImpl) GetGamesByDeveloper(developerID int) ([]*dto.GameDTO, error) {
	// Get games from repository
	games, err := s.gameRepo.FindByDeveloper(developerID)
	if err != nil {
		return nil, err
	}

	// Load related data for all games
	for _, game := range games {
		// Load developer if not loaded
		if game.Developer == nil && game.DeveloperID != 0 {
			developer, err := s.developerRepo.FindByID(game.DeveloperID)
			if err != nil {
				return nil, err
			}
			game.Developer = developer
		}

		// Load category if not loaded
		if game.Category == nil && game.CategoryID != 0 {
			category, err := s.categoryRepo.FindByID(game.CategoryID)
			if err != nil {
				return nil, err
			}
			game.Category = category
		}

		// Load restrictions
		restrictions, err := s.restrictRepo.FindByGameID(game.ID)
		if err != nil {
			return nil, err
		}
		game.Restricts = restrictions
	}

	// Convert models to DTOs for response
	return dto.GameDTOsFromModels(games), nil
}

// GetTopSellingGames retrieves top selling games
func (s *GameServiceImpl) GetTopSellingGames(limit int) ([]*dto.GameDTO, error) {
	// Ограничиваем количество возвращаемых игр
	if limit <= 0 {
		limit = 10 // дефолтное ограничение
	}

	// Получаем все игры
	games, err := s.gameRepo.FindAll(100, 0) // Получаем достаточное количество игр для выборки
	if err != nil {
		return nil, err
	}

	// Имитируем сортировку по продажам (в реальном приложении здесь был бы запрос к репозиторию)
	// Для демонстрации просто берем первые N игр и считаем их "топовыми"
	if len(games) > limit {
		games = games[:limit]
	}

	// Загрузим связанные данные для всех игр
	for _, game := range games {
		// Загружаем разработчика, если не загружен
		if game.Developer == nil && game.DeveloperID != 0 {
			developer, err := s.developerRepo.FindByID(game.DeveloperID)
			if err != nil {
				return nil, err
			}
			game.Developer = developer
		}

		// Загружаем категорию, если не загружена
		if game.Category == nil && game.CategoryID != 0 {
			category, err := s.categoryRepo.FindByID(game.CategoryID)
			if err != nil {
				return nil, err
			}
			game.Category = category
		}

		// Загружаем ограничения
		restrictions, err := s.restrictRepo.FindByGameID(game.ID)
		if err != nil {
			return nil, err
		}
		game.Restricts = restrictions
	}

	// Преобразуем модели в DTO для ответа
	return dto.GameDTOsFromModels(games), nil
}

// GetDiscountedGames retrieves games with discounts
func (s *GameServiceImpl) GetDiscountedGames(limit int) ([]*dto.GameDTO, error) {
	// Ограничиваем количество возвращаемых игр
	if limit <= 0 {
		limit = 10 // дефолтное ограничение
	}

	// Получаем все игры
	games, err := s.gameRepo.FindAll(100, 0) // Получаем достаточное количество игр для выборки
	if err != nil {
		return nil, err
	}

	// Имитируем фильтрацию по скидкам (в реальном приложении здесь был бы специальный запрос к репозиторию)
	// Для демонстрации просто берем первые N игр и считаем их "со скидками"
	if len(games) > limit {
		games = games[:limit]
	}

	// Загрузим связанные данные для всех игр
	for _, game := range games {
		// Загружаем разработчика, если не загружен
		if game.Developer == nil && game.DeveloperID != 0 {
			developer, err := s.developerRepo.FindByID(game.DeveloperID)
			if err != nil {
				return nil, err
			}
			game.Developer = developer
		}

		// Загружаем категорию, если не загружена
		if game.Category == nil && game.CategoryID != 0 {
			category, err := s.categoryRepo.FindByID(game.CategoryID)
			if err != nil {
				return nil, err
			}
			game.Category = category
		}

		// Загружаем ограничения
		restrictions, err := s.restrictRepo.FindByGameID(game.ID)
		if err != nil {
			return nil, err
		}
		game.Restricts = restrictions
	}

	// Преобразуем модели в DTO для ответа
	return dto.GameDTOsFromModels(games), nil
}
