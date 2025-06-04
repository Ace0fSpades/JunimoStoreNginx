import axios, { AxiosInstance } from 'axios';
import apiConfig, { shouldUseMocks } from '@/config/apiConfig';

// Логируем режим работы при старте приложения
console.log(`API инициализирован. Режим моков: ${shouldUseMocks() ? 'включен' : 'выключен'}`);

// Создаем экземпляр axios с настройками по умолчанию
const api: AxiosInstance = axios.create({
  baseURL: apiConfig.baseUrl,
  timeout: apiConfig.timeout,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  }
});

// Функция для добавления токена авторизации к запросу
export const setAuthToken = (token: string | null): void => {
  if (token) {
    api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
    localStorage.setItem('token', token);
  } else {
    delete api.defaults.headers.common['Authorization'];
    localStorage.removeItem('token');
  }
};

// Инициализация токена из localStorage при загрузке приложения
const token = localStorage.getItem('token');
if (token) {
  setAuthToken(token);
}

export default api; 