import { configureStore } from '@reduxjs/toolkit';
import createSagaMiddleware from 'redux-saga';
import { rootReducer } from './rootReducer';
import { rootSaga } from './rootSaga';

// Создаем middleware для Redux-Saga
const sagaMiddleware = createSagaMiddleware();

// Конфигурируем хранилище Redux
export const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) => 
    getDefaultMiddleware({ thunk: false }).concat(sagaMiddleware),
});

// Запускаем корневую сагу
sagaMiddleware.run(rootSaga);

// Определяем тип для всего store
export type AppDispatch = typeof store.dispatch;

// Экспортируем RootState из rootReducer
export type { RootState } from './types'; 