import { createSelector } from '@reduxjs/toolkit';
import { RootState } from '@/store/types';

// Basic selectors
const selectFavoritesState = (state: RootState) => state.favorites;

// Derived selectors
export const selectFavorites = createSelector(
  [selectFavoritesState],
  (favoritesState) => favoritesState.favorites
);

export const selectFavoriteItems = createSelector(
  [selectFavorites],
  (favorites) => favorites?.favoriteItems || []
);

export const selectFavoritesCount = createSelector(
  [selectFavoriteItems],
  (items) => items.length
);

export const selectFavoritesLoading = createSelector(
  [selectFavoritesState],
  (favoritesState) => favoritesState.loading
);

export const selectFavoritesError = createSelector(
  [selectFavoritesState],
  (favoritesState) => favoritesState.error
);

// Parameterized selectors
export const selectIsGameInFavorites = (gameId: number) =>
  createSelector(
    [selectFavoriteItems],
    (items) => items.some(item => item.gameID === gameId)
  );

export const selectFavoriteGameIds = createSelector(
  [selectFavoriteItems],
  (items) => items.map(item => item.gameID)
); 