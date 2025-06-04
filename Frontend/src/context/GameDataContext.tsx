import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { GameService } from '@/services/GameService';
import { Category, Developer } from '@/models/Game';

/**
 * Примечание:
 * В будущем для управления состоянием данных игры лучше использовать Redux с Redux-Saga,
 * что позволит лучше организовать асинхронные запросы и кэширование данных.
 * 
 * Пример структуры Redux store:
 * {
 *   games: {
 *     items: [], // список всех игр
 *     categories: [], // список всех категорий
 *     developers: [], // список всех разработчиков
 *     loading: false,
 *     error: null
 *   }
 * }
 * 
 * В этом случае, вместо этого контекста можно будет использовать Redux actions:
 * - FETCH_CATEGORIES_REQUEST
 * - FETCH_CATEGORIES_SUCCESS
 * - FETCH_CATEGORIES_FAILURE
 * 
 * И соответствующие саги для выполнения API-запросов.
 */

interface GameDataContextType {
  categories: Category[];
  developers: Developer[];
  loading: boolean;
  error: string | null;
}

const defaultContextValue: GameDataContextType = {
  categories: [],
  developers: [],
  loading: true,
  error: null
};

const GameDataContext = createContext<GameDataContextType>(defaultContextValue);

export const useGameData = () => useContext(GameDataContext);

interface GameDataProviderProps {
  children: ReactNode;
}

export const GameDataProvider: React.FC<GameDataProviderProps> = ({ children }) => {
  const [categories, setCategories] = useState<Category[]>([]);
  const [developers, setDevelopers] = useState<Developer[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        
        // Загружаем категории и разработчиков напрямую с сервера
        const [categoriesResponse, developersResponse] = await Promise.all([
          GameService.getCategories(),
          GameService.getDevelopers()
        ]);
        
        setCategories(categoriesResponse);
        setDevelopers(developersResponse);
        setError(null);
      } catch (err) {
        console.error('Ошибка при загрузке данных:', err);
        setError('Не удалось загрузить данные. Пожалуйста, попробуйте позже.');
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  return (
    <GameDataContext.Provider value={{
      categories,
      developers,
      loading,
      error
    }}>
      {children}
    </GameDataContext.Provider>
  );
};

export default GameDataContext; 