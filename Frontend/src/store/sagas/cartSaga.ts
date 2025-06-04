import { takeLatest, put, call, all, fork } from 'redux-saga/effects';
import { PayloadAction } from '@reduxjs/toolkit';
import { CartService } from '@/services/CartService';
import { ShoppingCart } from '@/models/Cart';
import {
  getCartRequest,
  getCartSuccess,
  getCartFailure,
  addToCartRequest,
  addToCartSuccess,
  addToCartFailure,
  removeFromCartRequest,
  removeFromCartSuccess,
  removeFromCartFailure,
  updateQuantityRequest,
  updateQuantitySuccess,
  updateQuantityFailure,
  clearCartRequest,
  clearCartSuccess,
  clearCartFailure,
} from '@/store/slices/cartSlice';

// Worker Sagas
function* getCartSaga(action: PayloadAction<number>) {
  try {
    const cart: ShoppingCart = yield call(CartService.getCart, action.payload);
    yield put(getCartSuccess(cart));
  } catch (error: any) {
    yield put(getCartFailure(error.response?.data?.message || 'Failed to fetch cart'));
  }
}

function* addToCartSaga(action: PayloadAction<{ userId: number, gameId: number, quantity?: number }>) {
  try {
    const { userId, gameId, quantity } = action.payload;
    const cart: ShoppingCart = yield call(CartService.addToCart, userId, gameId, quantity);
    yield put(addToCartSuccess(cart));
  } catch (error: any) {
    yield put(addToCartFailure(error.response?.data?.message || 'Failed to add item to cart'));
  }
}

function* removeFromCartSaga(action: PayloadAction<{ userId: number, gameId: number }>) {
  try {
    const { userId, gameId } = action.payload;
    yield call(CartService.removeFromCart, userId, gameId);
    
    // After removal, fetch the updated cart
    const updatedCart: ShoppingCart = yield call(CartService.getCart, userId);
    yield put(removeFromCartSuccess(updatedCart));
  } catch (error: any) {
    yield put(removeFromCartFailure(error.response?.data?.message || 'Failed to remove item from cart'));
  }
}

function* updateQuantitySaga(action: PayloadAction<{ userId: number, gameId: number, quantity: number }>) {
  try {
    const { userId, gameId, quantity } = action.payload;
    const cart: ShoppingCart = yield call(CartService.updateQuantity, userId, gameId, quantity);
    yield put(updateQuantitySuccess(cart));
  } catch (error: any) {
    yield put(updateQuantityFailure(error.response?.data?.message || 'Failed to update quantity'));
  }
}

function* clearCartSaga(action: PayloadAction<number>) {
  try {
    yield call(CartService.clearCart, action.payload);
    yield put(clearCartSuccess());
  } catch (error: any) {
    yield put(clearCartFailure(error.response?.data?.message || 'Failed to clear cart'));
  }
}

// Watcher Sagas
function* watchGetCart() {
  yield takeLatest(getCartRequest.type, getCartSaga);
}

function* watchAddToCart() {
  yield takeLatest(addToCartRequest.type, addToCartSaga);
}

function* watchRemoveFromCart() {
  yield takeLatest(removeFromCartRequest.type, removeFromCartSaga);
}

function* watchUpdateQuantity() {
  yield takeLatest(updateQuantityRequest.type, updateQuantitySaga);
}

function* watchClearCart() {
  yield takeLatest(clearCartRequest.type, clearCartSaga);
}

// Root Cart Saga
export function* watchCart() {
  yield all([
    fork(watchGetCart),
    fork(watchAddToCart),
    fork(watchRemoveFromCart),
    fork(watchUpdateQuantity),
    fork(watchClearCart),
  ]);
} 