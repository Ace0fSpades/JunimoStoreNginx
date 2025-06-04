import { all, fork } from 'redux-saga/effects';
import { watchAuth } from '@/store/sagas/authSaga';
import { watchGames } from '@/store/sagas/gamesSaga';
import { watchCart } from '@/store/sagas/cartSaga';
import { watchOrders } from '@/store/sagas/ordersSaga';
import { watchFavorites } from '@/store/sagas/favoritesSaga';
import { watchLibrary } from '@/store/sagas/librarySaga';
import { watchNews } from '@/store/sagas/newsSaga';

// Корневая сага, объединяющая все саги приложения
export function* rootSaga(): Generator<any, void, any> {
  yield all([
    fork(watchAuth),
    fork(watchGames),
    fork(watchCart),
    fork(watchOrders),
    fork(watchFavorites),
    fork(watchLibrary),
    fork(watchNews),
  ]);
} 