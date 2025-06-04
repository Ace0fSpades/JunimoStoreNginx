import { takeLatest, put, call, all, fork } from 'redux-saga/effects';
import { PayloadAction } from '@reduxjs/toolkit';
import { LibraryService } from '@/services/LibraryService';
import { Library } from '@/models/Cart';
import {
  getLibraryRequest,
  getLibrarySuccess,
  getLibraryFailure,
  addToLibraryRequest,
  addToLibrarySuccess,
  addToLibraryFailure,
  removeFromLibraryRequest,
  removeFromLibrarySuccess,
  removeFromLibraryFailure,
} from '@/store/slices/librarySlice';

// Worker Sagas
function* getLibrarySaga(action: PayloadAction<number>) {
  try {
    const library: Library = yield call(LibraryService.getLibrary, action.payload);
    yield put(getLibrarySuccess(library));
  } catch (error: any) {
    yield put(getLibraryFailure(error.response?.data?.message || 'Failed to fetch library'));
  }
}

function* addToLibrarySaga(action: PayloadAction<{ userId: number, gameId: number }>) {
  try {
    const { userId, gameId } = action.payload;
    const library: Library = yield call(LibraryService.addToLibrary, userId, gameId);
    yield put(addToLibrarySuccess(library));
  } catch (error: any) {
    yield put(addToLibraryFailure(error.response?.data?.message || 'Failed to add to library'));
  }
}

function* removeFromLibrarySaga(action: PayloadAction<{ userId: number, gameId: number }>) {
  try {
    const { userId, gameId } = action.payload;
    yield call(LibraryService.removeFromLibrary, userId, gameId);
    
    // After removal, fetch the updated library
    const updatedLibrary: Library = yield call(LibraryService.getLibrary, userId);
    yield put(removeFromLibrarySuccess(updatedLibrary));
  } catch (error: any) {
    yield put(removeFromLibraryFailure(error.response?.data?.message || 'Failed to remove from library'));
  }
}

// Watcher Sagas
function* watchGetLibrary() {
  yield takeLatest(getLibraryRequest.type, getLibrarySaga);
}

function* watchAddToLibrary() {
  yield takeLatest(addToLibraryRequest.type, addToLibrarySaga);
}

function* watchRemoveFromLibrary() {
  yield takeLatest(removeFromLibraryRequest.type, removeFromLibrarySaga);
}

// Root Library Saga
export function* watchLibrary() {
  yield all([
    fork(watchGetLibrary),
    fork(watchAddToLibrary),
    fork(watchRemoveFromLibrary),
  ]);
} 