import { createSelector } from '@reduxjs/toolkit';
import { RootState } from '@/store/types';

// Basic selectors
const selectOrdersState = (state: RootState) => state.orders;

// Derived selectors
export const selectAllOrders = createSelector(
  [selectOrdersState],
  (ordersState) => ordersState.orders
);

export const selectOrdersLoading = createSelector(
  [selectOrdersState],
  (ordersState) => ordersState.loading
);

export const selectOrdersError = createSelector(
  [selectOrdersState],
  (ordersState) => ordersState.error
);

// Parameterized selectors
export const selectOrderById = (orderId: number) => 
  createSelector(
    [selectAllOrders],
    (orders) => orders.find(order => order.id === orderId) || null
  );

export const selectOrdersByStatus = (status: string) =>
  createSelector(
    [selectAllOrders],
    (orders) => orders.filter(order => order.status === status)
  );

// Utility selectors
export const selectCompletedOrders = createSelector(
  [selectAllOrders],
  (orders) => orders.filter(order => order.status === 'completed')
);

export const selectPendingOrders = createSelector(
  [selectAllOrders],
  (orders) => orders.filter(order => order.status === 'pending')
);

export const selectTotalSpent = createSelector(
  [selectAllOrders],
  (orders) => orders.reduce((total, order) => total + order.totalCost, 0)
);

export const selectOrdersCount = createSelector(
  [selectAllOrders],
  (orders) => orders.length
); 