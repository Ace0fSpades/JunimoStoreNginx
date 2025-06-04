import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Game, Category, Developer } from '@/models/Game';
import { GamesState } from '@/store/types';

// Начальное состояние
const initialState: GamesState = {
  games: [],
  game: null,
  categories: [],
  developers: [],
  loading: false,
  error: null,
};

// Создаем slice для игр
export const gamesSlice = createSlice({
  name: 'games',
  initialState,
  reducers: {
    // Get all games
    getAllGamesRequest: (state, _action: PayloadAction<{ limit?: number, offset?: number }>) => {
      state.loading = true;
      state.error = null;
    },
    getAllGamesSuccess: (state, action: PayloadAction<Game[]>) => {
      state.games = action.payload;
      state.loading = false;
    },
    getAllGamesFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Get game by ID
    getGameByIdRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    getGameByIdSuccess: (state, action: PayloadAction<Game>) => {
      state.game = action.payload;
      state.loading = false;
    },
    getGameByIdFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Search games
    searchGamesRequest: (state, _action: PayloadAction<string>) => {
      state.loading = true;
      state.error = null;
    },
    searchGamesSuccess: (state, action: PayloadAction<Game[]>) => {
      state.games = action.payload;
      state.loading = false;
    },
    searchGamesFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Get games by category
    getGamesByCategoryRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    getGamesByCategorySuccess: (state, action: PayloadAction<Game[]>) => {
      state.games = action.payload;
      state.loading = false;
    },
    getGamesByCategoryFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Get latest games
    getLatestGamesRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    getLatestGamesSuccess: (state, action: PayloadAction<Game[]>) => {
      state.games = action.payload;
      state.loading = false;
    },
    getLatestGamesFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Get top selling games
    getTopSellingGamesRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    getTopSellingGamesSuccess: (state, action: PayloadAction<Game[]>) => {
      state.games = action.payload;
      state.loading = false;
    },
    getTopSellingGamesFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Get discounted games
    getDiscountedGamesRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    getDiscountedGamesSuccess: (state, action: PayloadAction<Game[]>) => {
      state.games = action.payload;
      state.loading = false;
    },
    getDiscountedGamesFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Get categories
    getCategoriesRequest: (state) => {
      state.loading = true;
      state.error = null;
    },
    getCategoriesSuccess: (state, action: PayloadAction<Category[]>) => {
      state.categories = action.payload;
      state.loading = false;
    },
    getCategoriesFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Get developers
    getDevelopersRequest: (state) => {
      state.loading = true;
      state.error = null;
    },
    getDevelopersSuccess: (state, action: PayloadAction<Developer[]>) => {
      state.developers = action.payload;
      state.loading = false;
    },
    getDevelopersFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Clear current game
    clearCurrentGame: (state) => {
      state.game = null;
    },
  },
});

// Экспортируем экшены
export const {
  getAllGamesRequest,
  getAllGamesSuccess,
  getAllGamesFailure,
  getGameByIdRequest,
  getGameByIdSuccess,
  getGameByIdFailure,
  searchGamesRequest,
  searchGamesSuccess,
  searchGamesFailure,
  getGamesByCategoryRequest,
  getGamesByCategorySuccess,
  getGamesByCategoryFailure,
  getLatestGamesRequest,
  getLatestGamesSuccess,
  getLatestGamesFailure,
  getTopSellingGamesRequest,
  getTopSellingGamesSuccess,
  getTopSellingGamesFailure,
  getDiscountedGamesRequest,
  getDiscountedGamesSuccess,
  getDiscountedGamesFailure,
  getCategoriesRequest,
  getCategoriesSuccess,
  getCategoriesFailure,
  getDevelopersRequest,
  getDevelopersSuccess,
  getDevelopersFailure,
  clearCurrentGame,
} = gamesSlice.actions;

// Экспортируем редьюсер
export default gamesSlice.reducer; 