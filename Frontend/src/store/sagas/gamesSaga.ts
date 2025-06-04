import { takeLatest, put, call, all, fork } from 'redux-saga/effects';
import { PayloadAction } from '@reduxjs/toolkit';
import { GameService } from '@/services/GameService';
import { Game, Category, Developer } from '@/models/Game';
import {
  getAllGamesRequest,
  getAllGamesSuccess,
  getAllGamesFailure,
  getGameByIdRequest,
  getGameByIdSuccess,
  getGameByIdFailure,
  searchGamesRequest,
  searchGamesSuccess,
  searchGamesFailure,
  getGamesByCategoryRequest,
  getGamesByCategorySuccess,
  getGamesByCategoryFailure,
  getLatestGamesRequest,
  getLatestGamesSuccess,
  getLatestGamesFailure,
  getTopSellingGamesRequest,
  getTopSellingGamesSuccess,
  getTopSellingGamesFailure,
  getDiscountedGamesRequest,
  getDiscountedGamesSuccess,
  getDiscountedGamesFailure,
  getCategoriesRequest,
  getCategoriesSuccess,
  getCategoriesFailure,
  getDevelopersRequest,
  getDevelopersSuccess,
  getDevelopersFailure,
} from '@/store/slices/gamesSlice';

// Worker Sagas
function* getAllGamesSaga(action: PayloadAction<{ limit?: number, offset?: number }>) {
  try {
    const { limit, offset } = action.payload;
    const games: Game[] = yield call(GameService.getAllGames, limit, offset);
    yield put(getAllGamesSuccess(games));
  } catch (error: any) {
    yield put(getAllGamesFailure(error.response?.data?.message || 'Failed to fetch games'));
  }
}

function* getGameByIdSaga(action: PayloadAction<number>) {
  try {
    const game: Game = yield call(GameService.getGameById, action.payload);
    yield put(getGameByIdSuccess(game));
  } catch (error: any) {
    yield put(getGameByIdFailure(error.response?.data?.message || 'Failed to fetch game details'));
  }
}

function* searchGamesSaga(action: PayloadAction<string>) {
  try {
    const games: Game[] = yield call(GameService.searchGames, action.payload);
    yield put(searchGamesSuccess(games));
  } catch (error: any) {
    yield put(searchGamesFailure(error.response?.data?.message || 'Failed to search games'));
  }
}

function* getGamesByCategorySaga(action: PayloadAction<number>) {
  try {
    const games: Game[] = yield call(GameService.getGamesByCategory, action.payload);
    yield put(getGamesByCategorySuccess(games));
  } catch (error: any) {
    yield put(getGamesByCategoryFailure(error.response?.data?.message || 'Failed to fetch games by category'));
  }
}

function* getLatestGamesSaga(action: PayloadAction<number>) {
  try {
    const games: Game[] = yield call(GameService.getLatestGames, action.payload);
    yield put(getLatestGamesSuccess(games));
  } catch (error: any) {
    yield put(getLatestGamesFailure(error.response?.data?.message || 'Failed to fetch latest games'));
  }
}

function* getTopSellingGamesSaga(action: PayloadAction<number>) {
  try {
    const games: Game[] = yield call(GameService.getTopSellingGames, action.payload);
    yield put(getTopSellingGamesSuccess(games));
  } catch (error: any) {
    yield put(getTopSellingGamesFailure(error.response?.data?.message || 'Failed to fetch top selling games'));
  }
}

function* getDiscountedGamesSaga(action: PayloadAction<number>) {
  try {
    const games: Game[] = yield call(GameService.getDiscountedGames, action.payload);
    yield put(getDiscountedGamesSuccess(games));
  } catch (error: any) {
    yield put(getDiscountedGamesFailure(error.response?.data?.message || 'Failed to fetch discounted games'));
  }
}

function* getCategoriesSaga() {
  try {
    const categories: Category[] = yield call(GameService.getCategories);
    yield put(getCategoriesSuccess(categories));
  } catch (error: any) {
    yield put(getCategoriesFailure(error.response?.data?.message || 'Failed to fetch categories'));
  }
}

function* getDevelopersSaga() {
  try {
    const developers: Developer[] = yield call(GameService.getDevelopers);
    yield put(getDevelopersSuccess(developers));
  } catch (error: any) {
    yield put(getDevelopersFailure(error.response?.data?.message || 'Failed to fetch developers'));
  }
}

// Watcher Sagas
function* watchGetAllGames() {
  yield takeLatest(getAllGamesRequest.type, getAllGamesSaga);
}

function* watchGetGameById() {
  yield takeLatest(getGameByIdRequest.type, getGameByIdSaga);
}

function* watchSearchGames() {
  yield takeLatest(searchGamesRequest.type, searchGamesSaga);
}

function* watchGetGamesByCategory() {
  yield takeLatest(getGamesByCategoryRequest.type, getGamesByCategorySaga);
}

function* watchGetLatestGames() {
  yield takeLatest(getLatestGamesRequest.type, getLatestGamesSaga);
}

function* watchGetTopSellingGames() {
  yield takeLatest(getTopSellingGamesRequest.type, getTopSellingGamesSaga);
}

function* watchGetDiscountedGames() {
  yield takeLatest(getDiscountedGamesRequest.type, getDiscountedGamesSaga);
}

function* watchGetCategories() {
  yield takeLatest(getCategoriesRequest.type, getCategoriesSaga);
}

function* watchGetDevelopers() {
  yield takeLatest(getDevelopersRequest.type, getDevelopersSaga);
}

// Root Games Saga
export function* watchGames() {
  yield all([
    fork(watchGetAllGames),
    fork(watchGetGameById),
    fork(watchSearchGames),
    fork(watchGetGamesByCategory),
    fork(watchGetLatestGames),
    fork(watchGetTopSellingGames),
    fork(watchGetDiscountedGames),
    fork(watchGetCategories),
    fork(watchGetDevelopers),
  ]);
} 