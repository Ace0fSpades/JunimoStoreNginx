import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import styles from '@/styles/pages/AuthPages.module.scss';
import { FaEnvelope, FaLock, FaArrowRight } from 'react-icons/fa';

// Types
import { LoginRequest } from '@/models/User';

// Redux
import { useAppDispatch, useAppSelector } from '@/store/hooks';
import { loginRequest } from '@/store/slices/authSlice';

// Auth context
import { useAuth } from '@/context/AuthContext';

const LoginPage: React.FC = () => {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();
  
  // Состояние из Redux
  const { loading, error: authError } = useAppSelector(state => ({
    loading: state.auth.loading,
    error: state.auth.error
  }));
  
  // Form state
  const [formData, setFormData] = useState<LoginRequest>({
    email: '',
    password: ''
  });
  
  // Form validation
  const [errors, setErrors] = useState<Record<string, string>>({});
  
  // Server error
  const [serverError, setServerError] = useState<string | null>(null);
  
  useEffect(() => {
    // Устанавливаем серверную ошибку из Redux
    if (authError) {
      setServerError(authError);
    }
  }, [authError]);
  
  // Редирект после успешной авторизации
  useEffect(() => {
    if (isAuthenticated) {
      navigate('/');
    }
  }, [isAuthenticated, navigate]);
  
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
    // Clear error when user starts typing
    if (errors[name]) {
      setErrors(prev => ({ ...prev, [name]: '' }));
    }
  };
  
  const validateForm = (): boolean => {
    const newErrors: Record<string, string> = {};
    
    // Email validation
    if (!formData.email.trim()) {
      newErrors.email = 'Email обязателен';
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.email)) {
      newErrors.email = 'Некорректный email';
    }
    
    // Password validation
    if (!formData.password) {
      newErrors.password = 'Пароль обязателен';
    }
    
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };
  
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    // Validate form
    if (!validateForm()) {
      return;
    }
    
    try {
      // Диспатчим экшен для логина
      dispatch(loginRequest(formData));
      // Редирект будет выполнен в useEffect при изменении isAuthenticated
    } catch (error) {
      console.error('Ошибка авторизации:', error);
      setServerError('Произошла ошибка при авторизации. Пожалуйста, попробуйте снова.');
    }
  };
  
  return (
    <div className={styles.authPage}>
      <div className="container">
        <div className={styles.authContainer}>
          <div className={styles.authContent}>
            <h1>Вход в аккаунт</h1>
            <p className={styles.authSubtitle}>Войдите, чтобы получить доступ к покупкам и личному кабинету</p>
            
            {serverError && (
              <div className="error-message">{serverError}</div>
            )}
            
            <form className={styles.authForm} onSubmit={handleSubmit}>
              <div className={styles.formGroup}>
                <label htmlFor="email">Email:</label>
                <div className={styles.inputWithIcon}>
                  <FaEnvelope />
                  <input
                    type="email"
                    id="email"
                    name="email"
                    value={formData.email}
                    onChange={handleChange}
                    placeholder="Введите ваш email"
                  />
                </div>
                {errors.email && <span className={styles.error}>{errors.email}</span>}
              </div>
              
              <div className={styles.formGroup}>
                <label htmlFor="password">Пароль:</label>
                <div className={styles.inputWithIcon}>
                  <FaLock />
                  <input
                    type="password"
                    id="password"
                    name="password"
                    value={formData.password}
                    onChange={handleChange}
                    placeholder="Введите ваш пароль"
                  />
                </div>
                {errors.password && <span className={styles.error}>{errors.password}</span>}
              </div>
              
              <div className={styles.buttonGroup}>
              <button type="button" className="btn btn-secondary" onClick={() => navigate('/register')}>
                  Ещё нет аккаунта?
                </button>
                <button type="submit" className="btn btn-primary" disabled={loading}>
                  {loading ? 'Вход...' : 'Войти'} <FaArrowRight />
                </button>
              </div>
              
              <div className={styles.authLinks}>
                <Link to="/" className={styles.underlineLink}>Вернуться на главную</Link>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage; 