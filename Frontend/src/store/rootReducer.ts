import { combineReducers } from '@reduxjs/toolkit';
import authReducer from '@/store/slices/authSlice';
import gamesReducer from '@/store/slices/gamesSlice';
import cartReducer from '@/store/slices/cartSlice';
import ordersReducer from '@/store/slices/ordersSlice';
import favoritesReducer from '@/store/slices/favoritesSlice';
import libraryReducer from '@/store/slices/librarySlice';
import newsReducer from '@/store/slices/newsSlice';

export const rootReducer = combineReducers({
  auth: authReducer,
  games: gamesReducer,
  cart: cartReducer,
  orders: ordersReducer,
  favorites: favoritesReducer,
  library: libraryReducer,
  news: newsReducer,
}); 