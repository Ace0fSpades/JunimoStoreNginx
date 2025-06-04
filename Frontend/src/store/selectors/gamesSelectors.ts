import { createSelector } from '@reduxjs/toolkit';
import { RootState } from '@/store/types';

// Basic selectors
const selectGamesState = (state: RootState) => state.games;

// Derived selectors
export const selectAllGames = createSelector(
  [selectGamesState],
  (gamesState) => gamesState.games
);

export const selectCurrentGame = createSelector(
  [selectGamesState],
  (gamesState) => gamesState.game
);

export const selectCategories = createSelector(
  [selectGamesState],
  (gamesState) => gamesState.categories
);

export const selectDevelopers = createSelector(
  [selectGamesState],
  (gamesState) => gamesState.developers
);

export const selectGamesLoading = createSelector(
  [selectGamesState],
  (gamesState) => gamesState.loading
);

export const selectGamesError = createSelector(
  [selectGamesState],
  (gamesState) => gamesState.error
);

// Parameterized selectors
export const selectGameById = (gameId: number) => 
  createSelector(
    [selectAllGames],
    (games) => games.find(game => game.id === gameId) || null
  );

export const selectGamesByCategory = (categoryId: number) =>
  createSelector(
    [selectAllGames],
    (games) => games.filter(game => game.categoryID === categoryId)
  );

export const selectGamesByDeveloper = (developerId: number) =>
  createSelector(
    [selectAllGames],
    (games) => games.filter(game => game.developerID === developerId)
  );

export const selectCategoryById = (categoryId: number) =>
  createSelector(
    [selectCategories],
    (categories) => categories.find(category => category.id === categoryId) || null
  );

export const selectDeveloperById = (developerId: number) =>
  createSelector(
    [selectDevelopers],
    (developers) => developers.find(developer => developer.id === developerId) || null
  ); 