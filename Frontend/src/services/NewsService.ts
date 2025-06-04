import { News } from '../models/News';
import { shouldUseMocks } from '../config/apiConfig';
import { getMockNews, getMockNewsById, getMockLatestNews } from '../mocks/news';

export const NewsService = {
  getAllNews: async (): Promise<News[]> => {
    if (shouldUseMocks()) {
      return getMockNews();
    }
    
    // В будущем здесь будет реальный API-вызов
    // Пока используем моки в любом случае
    return getMockNews();
  },

  getNewsById: async (id: number): Promise<News | null> => {
    if (shouldUseMocks()) {
      return getMockNewsById(id);
    }
    
    // В будущем здесь будет реальный API-вызов
    // Пока используем моки в любом случае
    return getMockNewsById(id);
  },

  getLatestNews: async (limit: number = 3): Promise<News[]> => {
    if (shouldUseMocks()) {
      return getMockLatestNews(limit);
    }
    
    // В будущем здесь будет реальный API-вызов
    // Пока используем моки в любом случае
    return getMockLatestNews(limit);
  },

  getNewsByCategory: async (category: string): Promise<News[]> => {
    const allNews = await NewsService.getAllNews();
    return allNews.filter(news => news.category === category);
  }
}; 