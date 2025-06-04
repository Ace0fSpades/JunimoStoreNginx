import { takeLatest, put, call, all, fork } from 'redux-saga/effects';
import { PayloadAction } from '@reduxjs/toolkit';
import { 
  loginRequest, 
  loginSuccess, 
  loginFailure, 
  registerRequest, 
  registerSuccess, 
  registerFailure,
  logout,
  getCurrentUserRequest,
  getCurrentUserSuccess,
  getCurrentUserFailure,
  updateUserRequest,
  updateUserSuccess,
  updateUserFailure
} from '@/store/slices/authSlice';
import { AuthService } from '@/services/AuthService';
import { LoginRequest, RegisterRequest, User, AuthResponse } from '@/models/User';

// Worker Sagas
function* loginSaga(action: PayloadAction<LoginRequest>) {
  try {
    const response: AuthResponse = yield call(AuthService.login, action.payload);
    yield put(loginSuccess(response.user));
    // Save user data to localStorage
    localStorage.setItem('user', JSON.stringify(response.user));
  } catch (error: any) {
    yield put(loginFailure(error.response?.data?.message || 'Failed to login'));
  }
}

function* registerSaga(action: PayloadAction<RegisterRequest>) {
  try {
    const response: AuthResponse = yield call(AuthService.register, action.payload);
    yield put(registerSuccess(response.user));
    // Save user data to localStorage
    localStorage.setItem('user', JSON.stringify(response.user));
  } catch (error: any) {
    yield put(registerFailure(error.response?.data?.message || 'Failed to register'));
  }
}

function* logoutSaga() {
  try {
    yield call(AuthService.logout);
    localStorage.removeItem('user');
  } catch (error: any) {
    console.error('Logout error:', error);
  }
}

function* getCurrentUserSaga() {
  try {
    const user: User | null = yield call(AuthService.getCurrentUser);
    yield put(getCurrentUserSuccess(user));
  } catch (error: any) {
    yield put(getCurrentUserFailure(error.message || 'Failed to get current user'));
  }
}

function* updateUserSaga(action: PayloadAction<{userId: number, userData: Partial<User>}>) {
  try {
    const { userId, userData } = action.payload;
    const updatedUser: User = yield call(AuthService.updateUser, userId, userData);
    yield put(updateUserSuccess(updatedUser));
    
    // Update user data in localStorage
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      const user = JSON.parse(storedUser);
      localStorage.setItem('user', JSON.stringify({ ...user, ...updatedUser }));
    }
  } catch (error: any) {
    yield put(updateUserFailure(error.response?.data?.message || 'Failed to update user'));
  }
}

// Watcher Sagas
function* watchLogin() {
  yield takeLatest(loginRequest.type, loginSaga);
}

function* watchRegister() {
  yield takeLatest(registerRequest.type, registerSaga);
}

function* watchLogout() {
  yield takeLatest(logout.type, logoutSaga);
}

function* watchGetCurrentUser() {
  yield takeLatest(getCurrentUserRequest.type, getCurrentUserSaga);
}

function* watchUpdateUser() {
  yield takeLatest(updateUserRequest.type, updateUserSaga);
}

// Root Auth Saga
export function* watchAuth() {
  yield all([
    fork(watchLogin),
    fork(watchRegister),
    fork(watchLogout),
    fork(watchGetCurrentUser),
    fork(watchUpdateUser)
  ]);
} 