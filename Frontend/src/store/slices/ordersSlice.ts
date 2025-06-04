import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { OrdersState } from '@/store/types';
import { Order } from '@/models/Order';

// Начальное состояние
const initialState: OrdersState = {
  orders: [],
  loading: false,
  error: null,
};

// Создаем slice для заказов
export const ordersSlice = createSlice({
  name: 'orders',
  initialState,
  reducers: {
    // Create order
    createOrderRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    createOrderSuccess: (state, action: PayloadAction<Order>) => {
      state.orders = [...state.orders, action.payload];
      state.loading = false;
    },
    createOrderFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Get order by ID
    getOrderByIdRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    getOrderByIdSuccess: (state, action: PayloadAction<Order>) => {
      // If order already exists in state, update it, otherwise add it
      const index = state.orders.findIndex(order => order.id === action.payload.id);
      if (index !== -1) {
        state.orders[index] = action.payload;
      } else {
        state.orders = [...state.orders, action.payload];
      }
      state.loading = false;
    },
    getOrderByIdFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Get user orders
    getUserOrdersRequest: (state, _action: PayloadAction<number>) => {
      state.loading = true;
      state.error = null;
    },
    getUserOrdersSuccess: (state, action: PayloadAction<Order[]>) => {
      state.orders = action.payload;
      state.loading = false;
    },
    getUserOrdersFailure: (state, action: PayloadAction<string>) => {
      state.error = action.payload;
      state.loading = false;
    },
    
    // Clear orders (e.g., on logout)
    clearOrders: (state) => {
      state.orders = [];
      state.error = null;
    },
  },
});

// Экспортируем экшены
export const {
  createOrderRequest,
  createOrderSuccess,
  createOrderFailure,
  getOrderByIdRequest,
  getOrderByIdSuccess,
  getOrderByIdFailure,
  getUserOrdersRequest,
  getUserOrdersSuccess,
  getUserOrdersFailure,
  clearOrders,
} = ordersSlice.actions;

// Экспортируем редьюсер
export default ordersSlice.reducer; 