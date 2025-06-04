import api, { setAuthToken } from './api';
import { AuthResponse, LoginRequest, RegisterRequest, User } from '../models/User';

interface UserUpdateData {
  name?: string;
  secondName?: string;
  thirdName?: string;
  email?: string;
  phoneNumber?: string;
}

export const AuthService = {
  login: async (credentials: LoginRequest) => {
    const response = await api.post<AuthResponse>('/auth/login', credentials);
    
    // Сохраняем токен и данные пользователя в localStorage
    if (response.data.token) {
      setAuthToken(response.data.token);
      localStorage.setItem('refreshToken', response.data.refreshToken);
      localStorage.setItem('user', JSON.stringify(response.data.user));
    }
    
    return response.data;
  },

  register: async (userData: RegisterRequest) => {
    const response = await api.post<AuthResponse>('/auth/signup', userData);
    
    // Сохраняем токен и данные пользователя в localStorage
    if (response.data.token) {
      setAuthToken(response.data.token);
      localStorage.setItem('refreshToken', response.data.refreshToken);
      localStorage.setItem('user', JSON.stringify(response.data.user));
    }
    
    return response.data;
  },

  refreshToken: async () => {
    const refreshToken = localStorage.getItem('refreshToken');
    if (!refreshToken) {
      throw new Error('No refresh token found');
    }
    
    const response = await api.post('/auth/refresh', { refresh_token: refreshToken });
    
    // Обновляем токен в localStorage и заголовках
    if (response.data.token) {
      setAuthToken(response.data.token);
      localStorage.setItem('refreshToken', response.data.refresh_token);
    }
    
    return response.data;
  },

  getCurrentUser: async () => {
    const userData = localStorage.getItem('user');
    if (userData) {
      const user = JSON.parse(userData);
      return user as User;
    }
    return null;
  },

  updateUser: async (userId: number, userData: UserUpdateData) => {
    const response = await api.patch<User>(`/users/${userId}/`, userData);
    return response.data;
  },

  logout: () => {
    localStorage.removeItem('user');
    localStorage.removeItem('token');
    localStorage.removeItem('refreshToken');
    setAuthToken(null);
  }
}; 