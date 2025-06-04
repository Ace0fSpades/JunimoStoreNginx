import api from './api';
import { Library } from '../models/Cart';

export const LibraryService = {
  getLibrary: async (userId: number) => {
    const response = await api.get<Library>(`/library/${userId}/`);
    return response.data;
  },
  
  // Эти методы могут быть реализованы на бэкенде в будущем
  addToLibrary: async (userId: number, gameId: number) => {
    const response = await api.post<Library>(`/library/${userId}/add/${gameId}/`);
    return response.data;
  },
  
  removeFromLibrary: async (userId: number, gameId: number) => {
    const response = await api.delete<{ success: boolean }>(`/library/${userId}/remove/${gameId}/`);
    return response.data;
  }
}; 