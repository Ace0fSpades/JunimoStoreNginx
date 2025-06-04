import { createSelector } from '@reduxjs/toolkit';
import { RootState } from '@/store/types';

// Basic selectors
const selectLibraryState = (state: RootState) => state.library;

// Derived selectors
export const selectLibrary = createSelector(
  [selectLibraryState],
  (libraryState) => libraryState.library
);

export const selectLibraryItems = createSelector(
  [selectLibrary],
  (library) => library?.libraryItems || []
);

export const selectLibraryCount = createSelector(
  [selectLibraryItems],
  (items) => items.length
);

export const selectLibraryLoading = createSelector(
  [selectLibraryState],
  (libraryState) => libraryState.loading
);

export const selectLibraryError = createSelector(
  [selectLibraryState],
  (libraryState) => libraryState.error
);

// Parameterized selectors
export const selectIsGameInLibrary = (gameId: number) =>
  createSelector(
    [selectLibraryItems],
    (items) => items.some(item => item.gameID === gameId)
  );

export const selectLibraryGameIds = createSelector(
  [selectLibraryItems],
  (items) => items.map(item => item.gameID)
); 