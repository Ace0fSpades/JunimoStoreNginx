import { createSelector } from '@reduxjs/toolkit';
import { RootState } from '@/store/types';

// Basic selectors
const selectCartState = (state: RootState) => state.cart;

// Derived selectors
export const selectCart = createSelector(
  [selectCartState],
  (cartState) => cartState.cart
);

export const selectCartItems = createSelector(
  [selectCart],
  (cart) => cart?.cartItems || []
);

export const selectCartItemsCount = createSelector(
  [selectCartItems],
  (items) => items.reduce((total, item) => total + item.quantity, 0)
);

export const selectCartTotal = createSelector(
  [selectCartItems],
  (items) => items.reduce((total, item) => total + (item.game.price * item.quantity), 0)
);

export const selectCartLoading = createSelector(
  [selectCartState],
  (cartState) => cartState.loading
);

export const selectCartError = createSelector(
  [selectCartState],
  (cartState) => cartState.error
);

// Parameterized selectors
export const selectCartItemByGameId = (gameId: number) => 
  createSelector(
    [selectCartItems],
    (items) => items.find(item => item.gameID === gameId) || null
  );

export const selectIsGameInCart = (gameId: number) =>
  createSelector(
    [selectCartItems],
    (items) => items.some(item => item.gameID === gameId)
  ); 