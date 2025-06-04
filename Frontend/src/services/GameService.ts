import api from './api';
import { Game, Category, Developer } from '@/models/Game';
import { shouldUseMocks } from '@/config/apiConfig';
import { 
  getMockGames, 
  getMockTopSellingGames, 
  getMockDiscountedGames,
  getMockCategories,
  getMockDevelopers
} from '@/mocks/games';
import { createGamePlaceholderImage } from '@/utils/createPlaceholderImage';

// Функция для добавления плейсхолдеров изображений, если их нет
const addImagePlaceholders = (games: Game[]): Game[] => {
  return games.map(game => {
    if (!game.imageURL && !game.image_data) {
      return {
        ...game,
        imageURL: createGamePlaceholderImage(300, 150, game.title)
      };
    }
    return game;
  });
};

export class GameService {
  /**
   * Получает список всех игр
   */
  static async getAllGames(limit = 10, offset = 0): Promise<Game[]> {
    // Если используем моки
    if (shouldUseMocks()) {
      console.log('[MOCK] Получение списка всех игр');
      const mockData = await getMockGames();
      const paginatedData = mockData.slice(offset, offset + limit);
      return addImagePlaceholders(paginatedData);
    }
    
    // В противном случае используем API
    try {
      const response = await api.get<Game[]>(`/games/?limit=${limit}&offset=${offset}`);
      return response.data;
    } catch (error) {
      console.error('Ошибка при получении всех игр:', error);
      throw error;
    }
  }

  /**
   * Получает детальную информацию об игре по ID
   */
  static async getGameById(id: number): Promise<Game> {
    // Если используем моки
    if (shouldUseMocks()) {
      console.log(`[MOCK] Получение игры с ID=${id}`);
      const mockGames = await getMockGames();
      const game = mockGames.find(g => g.id === id);
      
      if (!game) {
        throw new Error(`Игра с ID=${id} не найдена`);
      }
      
      return addImagePlaceholders([game])[0];
    }
    
    // В противном случае используем API
    try {
      const response = await api.get<Game>(`/games/${id}`);
      return response.data;
    } catch (error) {
      console.error(`Ошибка при получении игры с ID=${id}:`, error);
      throw error;
    }
  }

  /**
   * Получает список игр по категории
   */
  static async getGamesByCategory(categoryId: number): Promise<Game[]> {
    // Если используем моки
    if (shouldUseMocks()) {
      console.log(`[MOCK] Получение игр для категории ${categoryId}`);
      const mockGames = await getMockGames();
      const filteredGames = mockGames.filter(g => g.category.id === categoryId);
      return addImagePlaceholders(filteredGames);
    }
    
    // В противном случае используем API
    try {
      const response = await api.get<Game[]>(`/games/category/${categoryId}`);
      return response.data;
    } catch (error) {
      console.error(`Ошибка при получении игр для категории ${categoryId}:`, error);
      throw error;
    }
  }

  /**
   * Получает список хитов продаж
   */
  static async getTopSellingGames(limit = 5): Promise<Game[]> {
    // Если используем моки
    if (shouldUseMocks()) {
      console.log('[MOCK] Получение хитов продаж');
      const topGames = await getMockTopSellingGames(limit);
      return addImagePlaceholders(topGames);
    }
    
    // В противном случае используем API
    try {
      const response = await api.get<Game[]>(`/games/top-selling?limit=${limit}`);
      return response.data;
    } catch (error) {
      console.error('Ошибка при получении хитов продаж:', error);
      throw error;
    }
  }

  /**
   * Получает список игр со скидками
   */
  static async getDiscountedGames(limit = 5): Promise<Game[]> {
    // Если используем моки
    if (shouldUseMocks()) {
      console.log('[MOCK] Получение игр со скидками');
      const discountedGames = await getMockDiscountedGames(limit);
      return addImagePlaceholders(discountedGames);
    }
    
    // В противном случае используем API
    try {
      const response = await api.get<Game[]>(`/games/discounted?limit=${limit}`);
      return response.data;
    } catch (error) {
      console.error('Ошибка при получении игр со скидками:', error);
      throw error;
    }
  }

  /**
   * Получает список всех категорий
   */
  static async getCategories(): Promise<Category[]> {
    // Если используем моки
    if (shouldUseMocks()) {
      console.log('[MOCK] Получение списка категорий');
      const categories = await getMockCategories();
      return categories;
    }
    
    // В противном случае используем API
    try {
      const response = await api.get<Category[]>('/categories');
      return response.data;
    } catch (error) {
      console.error('Ошибка при получении категорий:', error);
      throw error;
    }
  }

  /**
   * Получает список всех разработчиков
   */
  static async getDevelopers(): Promise<Developer[]> {
    // Если используем моки
    if (shouldUseMocks()) {
      console.log('[MOCK] Получение списка разработчиков');
      const developers = await getMockDevelopers();
      return developers;
    }
    
    // В противном случае используем API
    try {
      const response = await api.get<Developer[]>('/developers');
      return response.data;
    } catch (error) {
      console.error('Ошибка при получении разработчиков:', error);
      throw error;
    }
  }

  /**
   * Получает список последних добавленных игр
   */
  static async getLatestGames(limit = 5): Promise<Game[]> {
    // Если используем моки
    if (shouldUseMocks()) {
      console.log('[MOCK] Получение последних добавленных игр');
      const mockGames = await getMockGames();
      // Сортируем игры по убыванию ID (предполагаем, что последние добавленные имеют больший ID)
      const latestGames = [...mockGames].sort((a, b) => b.id - a.id).slice(0, limit);
      return addImagePlaceholders(latestGames);
    }
    
    // В противном случае используем API
    try {
      const response = await api.get<Game[]>(`/games/latest?limit=${limit}`);
      return response.data;
    } catch (error) {
      console.error('Ошибка при получении последних игр:', error);
      throw error;
    }
  }

  /**
   * Поиск игр по названию
   */
  static async searchGames(title: string, limit = 10, offset = 0): Promise<Game[]> {
    // Если используем моки
    if (shouldUseMocks()) {
      console.log(`[MOCK] Поиск игр по запросу "${title}"`);
      const mockGames = await getMockGames();
      const filteredGames = mockGames.filter(g => 
        g.title.toLowerCase().includes(title.toLowerCase())
      );
      return addImagePlaceholders(filteredGames.slice(offset, offset + limit));
    }
    
    // В противном случае используем API
    try {
      const response = await api.get<Game[]>(`/games/search?title=${encodeURIComponent(title)}&limit=${limit}&offset=${offset}`);
      return response.data;
    } catch (error) {
      console.error(`Ошибка при поиске игр по запросу "${title}":`, error);
      throw error;
    }
  }
} 