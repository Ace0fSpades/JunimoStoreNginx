import { createSelector } from '@reduxjs/toolkit';
import { RootState } from '@/store/types';

// Basic selectors
const selectAuthState = (state: RootState) => state.auth;

// Derived selectors
export const selectUser = createSelector(
  [selectAuthState],
  (authState) => authState.user
);

export const selectIsAuthenticated = createSelector(
  [selectUser],
  (user) => !!user
);

export const selectAuthLoading = createSelector(
  [selectAuthState],
  (authState) => authState.loading
);

export const selectAuthError = createSelector(
  [selectAuthState],
  (authState) => authState.error
);

export const selectUserId = createSelector(
  [selectUser],
  (user) => user?.id || null
);

export const selectUserRole = createSelector(
  [selectUser],
  (user) => user?.role || null
);

export const selectIsAdmin = createSelector(
  [selectUserRole],
  (role) => role?.type === 'admin'
); 