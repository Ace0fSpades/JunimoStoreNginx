import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import styles from '../styles/pages/AuthPages.module.scss';
import { FaUser, FaEnvelope, FaLock, FaArrowRight } from 'react-icons/fa';

// Services
import { AuthService } from '../services/AuthService';

// Context
import { useAuth } from '../context/AuthContext';

// Types
import { RegisterRequest } from '../models/User';

const RegisterPage: React.FC = () => {
  const navigate = useNavigate();
  const { login } = useAuth();
  
  // Form state
  const [formData, setFormData] = useState<RegisterRequest>({
    nickname: '',
    email: '',
    password: ''
  });
  
  // Confirm password
  const [confirmPassword, setConfirmPassword] = useState('');
  
  // Form validation
  const [errors, setErrors] = useState<Record<string, string>>({});
  
  // Loading state
  const [loading, setLoading] = useState(false);
  
  // Server error
  const [serverError, setServerError] = useState<string | null>(null);
  
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
    // Clear error when user starts typing
    if (errors[name]) {
      setErrors(prev => ({ ...prev, [name]: '' }));
    }
  };
  
  const handleConfirmPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setConfirmPassword(e.target.value);
    if (errors.confirmPassword) {
      setErrors(prev => ({ ...prev, confirmPassword: '' }));
    }
  };
  
  const validateForm = (): boolean => {
    const newErrors: Record<string, string> = {};
    
    // Name validation
    if (!formData.nickname.trim()) {
      newErrors.nickname = 'Никнейм обязателен';
    }
    
    // Email validation
    if (!formData.email.trim()) {
      newErrors.email = 'Email обязателен';
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.email)) {
      newErrors.email = 'Некорректный email';
    }
    
    // Password validation
    if (!formData.password) {
      newErrors.password = 'Пароль обязателен';
    } else if (formData.password.length < 6) {
      newErrors.password = 'Пароль должен быть не менее 6 символов';
    }
    
    // Confirm password
    if (formData.password !== confirmPassword) {
      newErrors.confirmPassword = 'Пароли не совпадают';
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
      setLoading(true);
      setServerError(null);
      
      // Call register API
      const response = await AuthService.register(formData);
      
      // Log in the user with the received token
      login(response.user);
      
      // Navigate to home page
      navigate('/');
    } catch (error: any) {
      console.error('Registration error:', error);
      setServerError(
        error.response?.data?.message || 
        'Ошибка при регистрации. Пожалуйста, попробуйте позже.'
      );
    } finally {
      setLoading(false);
    }
  };
  
  return (
    <div className={`${styles.authPage} ${styles.registerPage}`}>
      <div className="container">
        <div className={styles.authContainer}>
          <div className={styles.authContent}>
            <h1>Регистрация</h1>
            <p className={styles.authSubtitle}>Создайте аккаунт, чтобы получить доступ к покупкам и личному кабинету</p>
            
            {serverError && (
              <div className="error-message">{serverError}</div>
            )}
            
            <form className={styles.authForm} onSubmit={handleSubmit}>
              <div className={styles.formRow}>
                <div className={styles.formGroup}>
                  <label htmlFor="nickname">Никнейм:</label>
                  <div className={styles.inputWithIcon}>
                    <FaUser />
                    <input
                      type="text"
                      id="nickname"
                      name="nickname"
                      value={formData.nickname}
                      onChange={handleChange}
                      placeholder="Введите ваш никнейм"
                      autoComplete="off"
                    />
                  </div>
                  {errors.nickname && <span className={styles.error}>{errors.nickname}</span>}
                </div>
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
                      autoComplete="off"
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
                      placeholder="Введите пароль"
                    />
                  </div>
                  {errors.password && <span className={styles.error}>{errors.password}</span>}
                </div>
                <div className={styles.formGroup}>
                  <label htmlFor="confirmPassword">Подтвердите пароль:</label>
                  <div className={styles.inputWithIcon}>
                    <FaLock />
                    <input
                      type="password"
                      id="confirmPassword"
                      name="confirmPassword"
                      value={confirmPassword}
                      onChange={handleConfirmPasswordChange}
                      placeholder="Повторите пароль"
                    />
                  </div>
                  {errors.confirmPassword && <span className={styles.error}>{errors.confirmPassword}</span>}
                </div>
              </div>
              
              <div className={styles.buttonGroup}>
              <button type="button" className="btn btn-secondary" onClick={() => navigate('/login')}>
                  Уже есть аккаунт?
                </button>
                <button type="submit" className="btn btn-primary" disabled={loading}>
                  {loading ? 'Регистрация...' : 'Зарегистрироваться'} <FaArrowRight />
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

export default RegisterPage; 