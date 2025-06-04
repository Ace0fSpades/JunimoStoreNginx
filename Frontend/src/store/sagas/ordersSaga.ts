import { takeLatest, put, call, all, fork } from 'redux-saga/effects';
import { PayloadAction } from '@reduxjs/toolkit';
import { OrderService } from '@/services/OrderService';
import { Order } from '@/models/Order';
import {
  createOrderRequest,
  createOrderSuccess,
  createOrderFailure,
  getOrderByIdRequest,
  getOrderByIdSuccess,
  getOrderByIdFailure,
  getUserOrdersRequest,
  getUserOrdersSuccess,
  getUserOrdersFailure,
} from '@/store/slices/ordersSlice';
import { clearCartSuccess } from '@/store/slices/cartSlice';

// Worker Sagas
function* createOrderSaga(action: PayloadAction<number>) {
  try {
    const order: Order = yield call(OrderService.createOrder, action.payload);
    yield put(createOrderSuccess(order));
    
    // Clear the cart after successful order creation
    yield put(clearCartSuccess());
  } catch (error: any) {
    yield put(createOrderFailure(error.response?.data?.message || 'Failed to create order'));
  }
}

function* getOrderByIdSaga(action: PayloadAction<number>) {
  try {
    const order: Order = yield call(OrderService.getOrderById, action.payload);
    yield put(getOrderByIdSuccess(order));
  } catch (error: any) {
    yield put(getOrderByIdFailure(error.response?.data?.message || 'Failed to fetch order details'));
  }
}

function* getUserOrdersSaga(action: PayloadAction<number>) {
  try {
    const orders: Order[] = yield call(OrderService.getUserOrders, action.payload);
    yield put(getUserOrdersSuccess(orders));
  } catch (error: any) {
    yield put(getUserOrdersFailure(error.response?.data?.message || 'Failed to fetch user orders'));
  }
}

// Watcher Sagas
function* watchCreateOrder() {
  yield takeLatest(createOrderRequest.type, createOrderSaga);
}

function* watchGetOrderById() {
  yield takeLatest(getOrderByIdRequest.type, getOrderByIdSaga);
}

function* watchGetUserOrders() {
  yield takeLatest(getUserOrdersRequest.type, getUserOrdersSaga);
}

// Root Orders Saga
export function* watchOrders() {
  yield all([
    fork(watchCreateOrder),
    fork(watchGetOrderById),
    fork(watchGetUserOrders),
  ]);
} 