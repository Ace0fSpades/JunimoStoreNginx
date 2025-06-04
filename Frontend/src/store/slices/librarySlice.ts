import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { LibraryState } from '@/store/types';
import { Library } from '@/models/Cart';

// Начальное состояние
const initialState: LibraryState = {
  library: null,
  loading: false,
  error: null,
};

// Создаем slice для библиотеки
export const librarySlice = createSlice({
  name: 'library',
  initialState,
  reducers: {
    // Get library
    getLibraryRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    getLibrarySuccess: (state, action: PayloadAction<Library>) => {
      state.library = action.payload;
      state.loading = false;
    },
    getLibraryFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Add to library
    addToLibraryRequest: (state, _action: PayloadAction<{ userId: number, gameId: number }>) => {
      state.loading = true;
      state.error = null;
    },
    addToLibrarySuccess: (state, action: PayloadAction<Library>) => {
      state.library = action.payload;
      state.loading = false;
    },
    addToLibraryFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Remove from library
    removeFromLibraryRequest: (state, _action: PayloadAction<{ userId: number, gameId: number }>) => {
      state.loading = true;
      state.error = null;
    },
    removeFromLibrarySuccess: (state, action: PayloadAction<Library>) => {
      state.library = action.payload;
      state.loading = false;
    },
    removeFromLibraryFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Clear library (e.g., on logout)
    clearLibrary: (state) => {
      state.library = null;
      state.error = null;
    },
  },
});

// Экспортируем экшены
export const {
  getLibraryRequest,
  getLibrarySuccess,
  getLibraryFailure,
  addToLibraryRequest,
  addToLibrarySuccess,
  addToLibraryFailure,
  removeFromLibraryRequest,
  removeFromLibrarySuccess,
  removeFromLibraryFailure,
  clearLibrary,
} = librarySlice.actions;

// Экспортируем редьюсер
export default librarySlice.reducer; 