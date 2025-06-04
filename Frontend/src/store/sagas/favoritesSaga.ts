import { takeLatest, put, call, all, fork } from 'redux-saga/effects';
import { PayloadAction } from '@reduxjs/toolkit';
import { FavoriteService } from '@/services/FavoriteService';
import { Favorite } from '@/models/Cart';
import {
  getFavoritesRequest,
  getFavoritesSuccess,
  getFavoritesFailure,
  addToFavoritesRequest,
  addToFavoritesSuccess,
  addToFavoritesFailure,
  removeFromFavoritesRequest,
  removeFromFavoritesSuccess,
  removeFromFavoritesFailure,
  clearFavoritesRequest,
  clearFavoritesSuccess,
  clearFavoritesFailure,
} from '@/store/slices/favoritesSlice';

// Worker Sagas
function* getFavoritesSaga(action: PayloadAction<number>) {
  try {
    const favorites: Favorite = yield call(FavoriteService.getFavorites, action.payload);
    yield put(getFavoritesSuccess(favorites));
  } catch (error: any) {
    yield put(getFavoritesFailure(error.response?.data?.message || 'Failed to fetch favorites'));
  }
}

function* addToFavoritesSaga(action: PayloadAction<{ userId: number, gameId: number }>) {
  try {
    const { userId, gameId } = action.payload;
    const favorites: Favorite = yield call(FavoriteService.addToFavorites, userId, gameId);
    yield put(addToFavoritesSuccess(favorites));
  } catch (error: any) {
    yield put(addToFavoritesFailure(error.response?.data?.message || 'Failed to add to favorites'));
  }
}

function* removeFromFavoritesSaga(action: PayloadAction<{ userId: number, gameId: number }>) {
  try {
    const { userId, gameId } = action.payload;
    yield call(FavoriteService.removeFromFavorites, userId, gameId);
    
    // After removal, fetch the updated favorites
    const updatedFavorites: Favorite = yield call(FavoriteService.getFavorites, userId);
    yield put(removeFromFavoritesSuccess(updatedFavorites));
  } catch (error: any) {
    yield put(removeFromFavoritesFailure(error.response?.data?.message || 'Failed to remove from favorites'));
  }
}

function* clearFavoritesSaga(action: PayloadAction<number>) {
  try {
    yield call(FavoriteService.clearFavorites, action.payload);
    yield put(clearFavoritesSuccess());
  } catch (error: any) {
    yield put(clearFavoritesFailure(error.response?.data?.message || 'Failed to clear favorites'));
  }
}

// Watcher Sagas
function* watchGetFavorites() {
  yield takeLatest(getFavoritesRequest.type, getFavoritesSaga);
}

function* watchAddToFavorites() {
  yield takeLatest(addToFavoritesRequest.type, addToFavoritesSaga);
}

function* watchRemoveFromFavorites() {
  yield takeLatest(removeFromFavoritesRequest.type, removeFromFavoritesSaga);
}

function* watchClearFavorites() {
  yield takeLatest(clearFavoritesRequest.type, clearFavoritesSaga);
}

// Root Favorites Saga
export function* watchFavorites() {
  yield all([
    fork(watchGetFavorites),
    fork(watchAddToFavorites),
    fork(watchRemoveFromFavorites),
    fork(watchClearFavorites),
  ]);
} 