/**
 * Конфигурация API и режима работы приложения
 */

interface ApiConfig {
  // Базовый URL для API запросов
  baseUrl: string;
  
  // Флаг, определяющий, нужно ли использовать моки вместо реальных API запросов
  useMocks: boolean;
  
  // Таймаут для запросов (в миллисекундах)
  timeout: number;
  
  // Заголовки по умолчанию для запросов
  headers: Record<string, string>;
}

// Вспомогательная функция для получения переменных окружения из Vite
const getEnvVariable = (key: string, defaultValue: string): string => {
  // Сначала проверяем import.meta.env (Vite)
  // @ts-ignore
  const viteValue = import.meta?.env?.[`VITE_${key}`] || null;
  
  // Потом проверяем process.env (для совместимости)
  const processValue = process.env?.[key] || null;
  
  return viteValue || processValue || defaultValue;
};

// Определяем режим работы с API: 'mocks' или 'api'
const API_MODE = getEnvVariable('API_MODE', 'mocks'); // Используем моки по умолчанию для GitHub Pages

// Определяем, используем ли мы моки
const USE_MOCKS = API_MODE === 'mocks';

// Определяем, находимся ли мы на GitHub Pages
const isGitHubPages = window.location.hostname.includes('github.io');

// Получаем хост и протокол из окружения
const host = window.location.hostname;
const protocol = window.location.protocol;

// Порт для API - по умолчанию 9090, если не указан иначе
const API_PORT = getEnvVariable('BACKEND_PORT', '9090');

// Базовый URL для API - для GitHub Pages всегда используем моки
// Для Docker и локальной разработки используем путь к API через указанный порт
const API_BASE_URL = isGitHubPages 
  ? '' 
  : `${protocol}//${host}:${API_PORT}/api/v1`; // Используем правильный порт и путь /api/v1/

// Логируем конфигурацию API
const logApiConfig = (config: ApiConfig) => {
  console.log('API Config:', {
    baseUrl: config.baseUrl,
    useMocks: config.useMocks,
    apiMode: API_MODE,
    isGitHubPages,
    host,
    protocol,
    apiPort: API_PORT,
    env: {
      NODE_ENV: process.env.NODE_ENV,
      API_MODE: getEnvVariable('API_MODE', 'mocks'),
    }
  });
};

// Конфигурация API
const apiConfig: ApiConfig = {
  baseUrl: API_BASE_URL,
  useMocks: USE_MOCKS || isGitHubPages, // Всегда используем моки на GitHub Pages
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  }
};

// Логируем конфигурацию
logApiConfig(apiConfig);

// Экспортируем функцию для проверки использования моков
export const shouldUseMocks = (): boolean => {
  return USE_MOCKS || isGitHubPages;
};

export default apiConfig; 