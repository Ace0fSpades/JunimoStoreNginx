import api from './api';
import { ShoppingCart } from '../models/Cart';

export const CartService = {
  getCart: async (userId: number) => {
    const response = await api.get<ShoppingCart>(`/cart/${userId}/`);
    return response.data;
  },

  addToCart: async (userId: number, gameId: number, quantity = 1) => {
    const response = await api.post<ShoppingCart>(`/cart/${userId}/add/${gameId}/`, { quantity });
    return response.data;
  },

  removeFromCart: async (userId: number, gameId: number) => {
    const response = await api.delete<{ success: boolean }>(`/cart/${userId}/remove/${gameId}/`);
    return response.data;
  },

  updateQuantity: async (userId: number, gameId: number, quantity: number) => {
    const response = await api.patch<ShoppingCart>(`/cart/${userId}/update/${gameId}/`, { quantity });
    return response.data;
  },

  clearCart: async (userId: number) => {
    const response = await api.delete<{ success: boolean }>(`/cart/${userId}/clear/`);
    return response.data;
  }
}; 