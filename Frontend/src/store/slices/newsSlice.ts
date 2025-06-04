import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { NewsState } from '@/store/types';
import { News } from '@/models/News';

// Начальное состояние
const initialState: NewsState = {
  newsList: [],
  newsItem: null,
  loading: false,
  error: null,
};

// Создаем slice для новостей
export const newsSlice = createSlice({
  name: 'news',
  initialState,
  reducers: {
    // Get all news
    getAllNewsRequest: (state) => {
      state.loading = true;
      state.error = null;
    },
    getAllNewsSuccess: (state, action: PayloadAction<News[]>) => {
      state.newsList = action.payload;
      state.loading = false;
    },
    getAllNewsFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Get news by ID
    getNewsByIdRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    getNewsByIdSuccess: (state, action: PayloadAction<News>) => {
      state.newsItem = action.payload;
      state.loading = false;
    },
    getNewsByIdFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Get news by category
    getNewsByCategoryRequest: (state, _action: PayloadAction<string>) => {
      state.loading = true;
      state.error = null;
    },
    getNewsByCategorySuccess: (state, action: PayloadAction<News[]>) => {
      state.newsList = action.payload;
      state.loading = false;
    },
    getNewsByCategoryFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Clear current news item
    clearCurrentNewsItem: (state) => {
      state.newsItem = null;
    },
  },
});

// Экспортируем экшены
export const {
  getAllNewsRequest,
  getAllNewsSuccess,
  getAllNewsFailure,
  getNewsByIdRequest,
  getNewsByIdSuccess,
  getNewsByIdFailure,
  getNewsByCategoryRequest,
  getNewsByCategorySuccess,
  getNewsByCategoryFailure,
  clearCurrentNewsItem,
} = newsSlice.actions;

// Экспортируем редьюсер
export default newsSlice.reducer; 