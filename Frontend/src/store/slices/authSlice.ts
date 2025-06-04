import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { AuthState } from '@/store/types';
import { LoginRequest, RegisterRequest, User } from '@/models/User';

// Начальное состояние
const initialState: AuthState = {
  user: null,
  loading: false,
  error: null,
};

// Создаем slice для аутентификации
export const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    // Login actions
    loginRequest: (state, _action: PayloadAction<LoginRequest>) => {
      state.loading = true;
      state.error = null;
    },
    loginSuccess: (state, action: PayloadAction<User>) => {
      state.user = action.payload;
      state.loading = false;
    },
    loginFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Register actions
    registerRequest: (state, _action: PayloadAction<RegisterRequest>) => {
      state.loading = true;
      state.error = null;
    },
    registerSuccess: (state, action: PayloadAction<User>) => {
      state.user = action.payload;
      state.loading = false;
    },
    registerFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Logout action
    logout: (state) => {
      state.user = null;
    },
    
    // Get current user action
    getCurrentUserRequest: (state) => {
      state.loading = true;
    },
    getCurrentUserSuccess: (state, action: PayloadAction<User | null>) => {
      state.user = action.payload;
      state.loading = false;
    },
    getCurrentUserFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Update user actions
    updateUserRequest: (state, _action: PayloadAction<{userId: number, userData: Partial<User>}>) => {
      state.loading = true;
      state.error = null;
    },
    updateUserSuccess: (state, action: PayloadAction<User>) => {
      state.user = action.payload;
      state.loading = false;
    },
    updateUserFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
  },
});

// Экспортируем экшены
export const {
  loginRequest,
  loginSuccess,
  loginFailure,
  registerRequest,
  registerSuccess,
  registerFailure,
  logout,
  getCurrentUserRequest,
  getCurrentUserSuccess,
  getCurrentUserFailure,
  updateUserRequest,
  updateUserSuccess,
  updateUserFailure,
} = authSlice.actions;

// Экспортируем редьюсер
export default authSlice.reducer; 