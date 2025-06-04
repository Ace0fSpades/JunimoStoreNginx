import api from './api';
import { Favorite } from '../models/Cart';

export const FavoriteService = {
  getFavorites: async (userId: number) => {
    const response = await api.get<Favorite>(`/favorite/${userId}/`);
    return response.data;
  },

  addToFavorites: async (userId: number, gameId: number) => {
    const response = await api.post<Favorite>(`/favorite/${userId}/add/${gameId}/`);
    return response.data;
  },

  removeFromFavorites: async (userId: number, gameId: number) => {
    const response = await api.delete<{ success: boolean }>(`/favorite/${userId}/remove/${gameId}/`);
    return response.data;
  },

  clearFavorites: async (userId: number) => {
    const response = await api.delete<{ success: boolean }>(`/favorite/${userId}/clear/`);
    return response.data;
  }
}; 