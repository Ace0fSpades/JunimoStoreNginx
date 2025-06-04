import React, { createContext, useContext, useEffect, ReactNode } from 'react';
import { useAppDispatch, useAppSelector } from '@/store/hooks';
import { loginSuccess, logout as logoutAction } from '@/store/slices/authSlice';
import { User } from '@/models/User';

interface AuthContextType {
  user: User | null;
  loading: boolean;
  login: (userData: User) => void;
  logout: () => void;
  isAuthenticated: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const dispatch = useAppDispatch();
  const { user, loading } = useAppSelector(state => state.auth);

  useEffect(() => {
    // Проверяем, залогинен ли пользователь при загрузке страницы
    const checkLoginStatus = () => {
      try {
        const storedUser = localStorage.getItem('user');
        
        if (storedUser) {
          const userData = JSON.parse(storedUser);
          dispatch(loginSuccess(userData));
        }
      } catch (error) {
        console.error('Error checking login status:', error);
        // Очищаем локальное хранилище при ошибке
        handleLogout();
      }
    };

    checkLoginStatus();
  }, [dispatch]);

  // Обработчик выхода, который можно использовать при ошибках авторизации
  const handleLogout = () => {
    localStorage.removeItem('user');
    localStorage.removeItem('token');
    localStorage.removeItem('refreshToken');
    dispatch(logoutAction());
  };

  const login = (userData: User) => {
    // Действие входа теперь обрабатывается через Redux, 
    // но оставляем этот метод для совместимости
    dispatch(loginSuccess(userData));
  };

  const logout = () => {
    // Отправляем экшен в Redux Store
    handleLogout();
  };

  const value = {
    user,
    loading,
    login,
    logout,
    isAuthenticated: !!user,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  
  return context;
}; 