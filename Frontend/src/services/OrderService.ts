import api from './api';
import { Order } from '../models/Order';

export const OrderService = {
  createOrder: async (userId: number) => {
    const response = await api.post<Order>(`/orders/${userId}/create/`);
    return response.data;
  },

  getOrderById: async (orderId: number) => {
    const response = await api.get<Order>(`/orders/${orderId}/`);
    return response.data;
  },

  getUserOrders: async (userId: number) => {
    const response = await api.get<Order[]>(`/orders/user/${userId}/`);
    return response.data;
  }
}; 