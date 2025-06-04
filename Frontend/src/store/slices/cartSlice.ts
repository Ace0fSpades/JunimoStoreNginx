import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { CartState } from '@/store/types';
import { ShoppingCart } from '@/models/Cart';

// Начальное состояние
const initialState: CartState = {
  cart: null,
  loading: false,
  error: null,
};

// Создаем slice для корзины
export const cartSlice = createSlice({
  name: 'cart',
  initialState,
  reducers: {
    // Get cart
    getCartRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    getCartSuccess: (state, action: PayloadAction<ShoppingCart>) => {
      state.cart = action.payload;
      state.loading = false;
    },
    getCartFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Add to cart
    addToCartRequest: (state, _action: PayloadAction<{ userId: number, gameId: number, quantity?: number }>) => {
      state.loading = true;
      state.error = null;
    },
    addToCartSuccess: (state, action: PayloadAction<ShoppingCart>) => {
      state.cart = action.payload;
      state.loading = false;
    },
    addToCartFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Remove from cart
    removeFromCartRequest: (state, _action: PayloadAction<{ userId: number, gameId: number }>) => {
      state.loading = true;
      state.error = null;
    },
    removeFromCartSuccess: (state, action: PayloadAction<ShoppingCart>) => {
      state.cart = action.payload;
      state.loading = false;
    },
    removeFromCartFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Update quantity
    updateQuantityRequest: (state, _action: PayloadAction<{ userId: number, gameId: number, quantity: number }>) => {
      state.loading = true;
      state.error = null;
    },
    updateQuantitySuccess: (state, action: PayloadAction<ShoppingCart>) => {
      state.cart = action.payload;
      state.loading = false;
    },
    updateQuantityFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Clear cart
    clearCartRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    clearCartSuccess: (state) => {
      state.cart = null;
      state.loading = false;
    },
    clearCartFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
  },
});

// Экспортируем экшены
export const {
  getCartRequest,
  getCartSuccess,
  getCartFailure,
  addToCartRequest,
  addToCartSuccess,
  addToCartFailure,
  removeFromCartRequest,
  removeFromCartSuccess,
  removeFromCartFailure,
  updateQuantityRequest,
  updateQuantitySuccess,
  updateQuantityFailure,
  clearCartRequest,
  clearCartSuccess,
  clearCartFailure,
} = cartSlice.actions;

// Экспортируем редьюсер
export default cartSlice.reducer; 