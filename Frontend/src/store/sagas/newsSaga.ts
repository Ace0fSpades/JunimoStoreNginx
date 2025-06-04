import { takeLatest, put, call, all, fork } from 'redux-saga/effects';
import { PayloadAction } from '@reduxjs/toolkit';
import { NewsService } from '@/services/NewsService';
import { News } from '@/models/News';
import {
  getAllNewsRequest,
  getAllNewsSuccess,
  getAllNewsFailure,
  getNewsByIdRequest,
  getNewsByIdSuccess,
  getNewsByIdFailure,
  getNewsByCategoryRequest,
  getNewsByCategorySuccess,
  getNewsByCategoryFailure,
} from '@/store/slices/newsSlice';

// Worker Sagas
function* getAllNewsSaga() {
  try {
    const news: News[] = yield call(NewsService.getAllNews);
    yield put(getAllNewsSuccess(news));
  } catch (error: any) {
    yield put(getAllNewsFailure(error.message || 'Failed to fetch news'));
  }
}

function* getNewsByIdSaga(action: PayloadAction<number>) {
  try {
    const newsItem: News | undefined = yield call(NewsService.getNewsById, action.payload);
    if (newsItem) {
      yield put(getNewsByIdSuccess(newsItem));
    } else {
      yield put(getNewsByIdFailure('News item not found'));
    }
  } catch (error: any) {
    yield put(getNewsByIdFailure(error.message || 'Failed to fetch news item'));
  }
}

function* getNewsByCategorySaga(action: PayloadAction<string>) {
  try {
    const news: News[] = yield call(NewsService.getNewsByCategory, action.payload);
    yield put(getNewsByCategorySuccess(news));
  } catch (error: any) {
    yield put(getNewsByCategoryFailure(error.message || 'Failed to fetch news by category'));
  }
}

// Watcher Sagas
function* watchGetAllNews() {
  yield takeLatest(getAllNewsRequest.type, getAllNewsSaga);
}

function* watchGetNewsById() {
  yield takeLatest(getNewsByIdRequest.type, getNewsByIdSaga);
}

function* watchGetNewsByCategory() {
  yield takeLatest(getNewsByCategoryRequest.type, getNewsByCategorySaga);
}

// Root News Saga
export function* watchNews() {
  yield all([
    fork(watchGetAllNews),
    fork(watchGetNewsById),
    fork(watchGetNewsByCategory),
  ]);
} 