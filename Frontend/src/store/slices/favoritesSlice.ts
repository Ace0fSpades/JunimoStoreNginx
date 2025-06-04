import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { FavoritesState } from '@/store/types';
import { Favorite } from '@/models/Cart';

// Начальное состояние
const initialState: FavoritesState = {
  favorites: null,
  loading: false,
  error: null,
};

// Создаем slice для избранного
export const favoritesSlice = createSlice({
  name: 'favorites',
  initialState,
  reducers: {
    // Get favorites
    getFavoritesRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    getFavoritesSuccess: (state, action: PayloadAction<Favorite>) => {
      state.favorites = action.payload;
      state.loading = false;
    },
    getFavoritesFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Add to favorites
    addToFavoritesRequest: (state, _action: PayloadAction<{ userId: number, gameId: number }>) => {
      state.loading = true;
      state.error = null;
    },
    addToFavoritesSuccess: (state, action: PayloadAction<Favorite>) => {
      state.favorites = action.payload;
      state.loading = false;
    },
    addToFavoritesFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Remove from favorites
    removeFromFavoritesRequest: (state, _action: PayloadAction<{ userId: number, gameId: number }>) => {
      state.loading = true;
      state.error = null;
    },
    removeFromFavoritesSuccess: (state, action: PayloadAction<Favorite>) => {
      state.favorites = action.payload;
      state.loading = false;
    },
    removeFromFavoritesFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Clear favorites
    clearFavoritesRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    clearFavoritesSuccess: (state) => {
      state.favorites = null;
      state.loading = false;
    },
    clearFavoritesFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
  },
});

// Экспортируем экшены
export const {
  getFavoritesRequest,
  getFavoritesSuccess,
  getFavoritesFailure,
  addToFavoritesRequest,
  addToFavoritesSuccess,
  addToFavoritesFailure,
  removeFromFavoritesRequest,
  removeFromFavoritesSuccess,
  removeFromFavoritesFailure,
  clearFavoritesRequest,
  clearFavoritesSuccess,
  clearFavoritesFailure,
} = favoritesSlice.actions;

// Экспортируем редьюсер
export default favoritesSlice.reducer; 