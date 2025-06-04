import React, { useState, useEffect } from 'react';
//import { Link } from 'react-router-dom';
import styles from '@/styles/pages/ProfilePage.module.scss';
import { FaUser, FaEnvelope, FaPhone, FaEdit, FaShoppingBag, FaHeart, FaGamepad } from 'react-icons/fa';

// Services
import { AuthService } from '@/services/AuthService';
import { OrderService } from '@/services/OrderService';
import { FavoriteService } from '@/services/FavoriteService';
import { LibraryService } from '@/services/LibraryService';

// Context
import { useAuth } from '@/context/AuthContext';

// Types
import { Order } from '@/models/Order';
import { Favorite, Library } from '@/models/Cart';

interface ProfileData {
  name: string;
  secondName?: string;
  thirdName?: string;
  email: string;
  phoneNumber?: string;
  role?: {
    type: string;
  };
}

const ProfilePage: React.FC = () => {
  const { user, logout } = useAuth();
  
  // States for user data
  const [userData, setUserData] = useState<ProfileData | null>(null);
  const [orders, setOrders] = useState<Order[]>([]);
  const [favorites, setFavorites] = useState<Favorite | null>(null);
  const [library, setLibrary] = useState<Library | null>(null);
  
  // Loading states
  const [loadingProfile, setLoadingProfile] = useState(true);
  const [loadingOrders, setLoadingOrders] = useState(true);
  const [loadingFavorites, setLoadingFavorites] = useState(true);
  const [loadingLibrary, setLoadingLibrary] = useState(true);
  
  // Error states
  const [error, setError] = useState<string | null>(null);
  
  // Edit mode state
  const [editMode, setEditMode] = useState(false);
  const [editData, setEditData] = useState<ProfileData>({
    name: '',
    secondName: '',
    thirdName: '',
    email: '',
    phoneNumber: ''
  });
  
  // Tab state
  const [activeTab, setActiveTab] = useState('profile');
  
  useEffect(() => {
    // Set initial user data
    if (user) {
      const profile: ProfileData = {
        name: user.name,
        email: user.email,
        role: user.role
      };
      
      setUserData(profile);
      setEditData({
        name: user.name,
        email: user.email,
        secondName: '',
        thirdName: '',
        phoneNumber: ''
      });
      
      // Fetch user data
      fetchUserData();
    }
  }, [user]);
  
  const fetchUserData = async () => {
    if (!user) return;
    
    try {
      // Fetch orders
      setLoadingOrders(true);
      const ordersData = await OrderService.getUserOrders(user.id);
      setOrders(ordersData);
      setLoadingOrders(false);
      
      // Fetch favorites
      setLoadingFavorites(true);
      const favoritesData = await FavoriteService.getFavorites(user.id);
      setFavorites(favoritesData);
      setLoadingFavorites(false);
      
      // Fetch library
      setLoadingLibrary(true);
      const libraryData = await LibraryService.getLibrary(user.id);
      setLibrary(libraryData);
      setLoadingLibrary(false);
      
      setError(null);
    } catch (err) {
      console.error('Error fetching user data:', err);
      setError('Ошибка при загрузке данных. Пожалуйста, попробуйте позже.');
    } finally {
      setLoadingProfile(false);
    }
  };
  
  const handleEditChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setEditData(prev => ({ ...prev, [name]: value }));
  };
  
  const handleEditSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!user) return;
    
    try {
      setLoadingProfile(true);
      
      // Update user data - передаём только обновленные поля
      await AuthService.updateUser(user.id, editData);
      
      // Update local user data
      setUserData({
        ...userData,
        ...editData
      });
      
      // Exit edit mode
      setEditMode(false);
      
      setError(null);
    } catch (err) {
      console.error('Error updating user data:', err);
      setError('Ошибка при обновлении данных. Пожалуйста, попробуйте позже.');
    } finally {
      setLoadingProfile(false);
    }
  };
  
  const handleLogout = () => {
    logout();
    window.location.href = '/';
  };
  
  if (loadingProfile && !userData) {
    return <div className="loading">Загрузка...</div>;
  }
  
  if (!userData) {
    return <div className="error-message">Пользователь не найден</div>;
  }
  
  return (
    <div className={styles.profilePage}>
      <div className="container">
        <h1>Личный кабинет</h1>
        
        {error && (
          <div className="error-message">{error}</div>
        )}
        
        <div className={styles.profileContent}>
          <div className={styles.profileSidebar}>
            <div className={styles.profileMenu}>
              <button 
                className={`${styles.profileMenuItem} ${activeTab === 'profile' ? styles.active : ''}`}
                onClick={() => setActiveTab('profile')}
              >
                <FaUser /> Профиль
              </button>
              <button 
                className={`${styles.profileMenuItem} ${activeTab === 'orders' ? styles.active : ''}`}
                onClick={() => setActiveTab('orders')}
              >
                <FaShoppingBag /> Заказы
              </button>
              <button 
                className={`${styles.profileMenuItem} ${activeTab === 'favorites' ? styles.active : ''}`}
                onClick={() => setActiveTab('favorites')}
              >
                <FaHeart /> Избранное
              </button>
              <button 
                className={`${styles.profileMenuItem} ${activeTab === 'library' ? styles.active : ''}`}
                onClick={() => setActiveTab('library')}
              >
                <FaGamepad /> Библиотека
              </button>
            </div>
            
            <button className={`btn btn-outline ${styles.logoutBtn}`} onClick={handleLogout}>
              Выйти из аккаунта
            </button>
          </div>
          
          <div className={styles.profileDetails}>
            {/* Profile Tab */}
            {activeTab === 'profile' && (
              <div className={styles.profileTab}>
                <div className={styles.profileHeader}>
                  <h2>Информация о профиле</h2>
                  {!editMode && (
                    <button 
                      className={`btn btn-outline ${styles.editBtn}`}
                      onClick={() => setEditMode(true)}
                    >
                      <FaEdit /> Редактировать
                    </button>
                  )}
                </div>
                
                {editMode ? (
                  <form className={styles.editForm} onSubmit={handleEditSubmit}>
                    <div className={styles.formRow}>
                      <div className={styles.formGroup}>
                        <label htmlFor="name">Имя</label>
                        <input
                          type="text"
                          id="name"
                          name="name"
                          value={editData.name || ''}
                          onChange={handleEditChange}
                          required
                        />
                      </div>
                      
                      <div className={styles.formGroup}>
                        <label htmlFor="secondName">Фамилия</label>
                        <input
                          type="text"
                          id="secondName"
                          name="secondName"
                          value={editData.secondName || ''}
                          onChange={handleEditChange}
                        />
                      </div>
                    </div>
                    
                    <div className={styles.formRow}>
                      <div className={styles.formGroup}>
                        <label htmlFor="thirdName">Отчество</label>
                        <input
                          type="text"
                          id="thirdName"
                          name="thirdName"
                          value={editData.thirdName || ''}
                          onChange={handleEditChange}
                        />
                      </div>
                      
                      <div className={styles.formGroup}>
                        <label htmlFor="phoneNumber">Телефон</label>
                        <input
                          type="tel"
                          id="phoneNumber"
                          name="phoneNumber"
                          value={editData.phoneNumber || ''}
                          onChange={handleEditChange}
                        />
                      </div>
                    </div>
                    
                    <div className={styles.formActions}>
                      <button 
                        type="submit" 
                        className="btn btn-primary"
                        disabled={loadingProfile}
                      >
                        {loadingProfile ? 'Сохранение...' : 'Сохранить'}
                      </button>
                      <button 
                        type="button" 
                        className="btn btn-outline"
                        onClick={() => setEditMode(false)}
                        disabled={loadingProfile}
                      >
                        Отмена
                      </button>
                    </div>
                  </form>
                ) : (
                  <div className={styles.userInfo}>
                    <div className={styles.infoItem}>
                      <span className={styles.infoLabel}>
                        <FaUser /> Полное имя:
                      </span>
                      <span className={styles.infoValue}>
                        {userData.name} {userData.secondName || ''} {userData.thirdName || ''}
                      </span>
                    </div>
                    
                    <div className={styles.infoItem}>
                      <span className={styles.infoLabel}>
                        <FaEnvelope /> Email:
                      </span>
                      <span className={styles.infoValue}>{userData.email}</span>
                    </div>
                    
                    <div className={styles.infoItem}>
                      <span className={styles.infoLabel}>
                        <FaPhone /> Телефон:
                      </span>
                      <span className={styles.infoValue}>{userData.phoneNumber || '-'}</span>
                    </div>
                  </div>
                )}
              </div>
            )}
            
            {/* Orders Tab */}
            {activeTab === 'orders' && (
              <div className={styles.ordersTab}>
                <h2>История заказов</h2>
                
                {loadingOrders ? (
                  <div className="loading">Загрузка заказов...</div>
                ) : orders.length > 0 ? (
                  <div className={styles.ordersList}>
                    {orders.map(order => (
                      <div key={order.id} className={styles.orderItem}>
                        <div className={styles.orderHeader}>
                          <div className={styles.orderInfo}>
                            <div className={styles.orderId}>Заказ №{order.id}</div>
                            <div className={styles.orderDate}>{new Date(order.createdAt).toLocaleDateString()}</div>
                          </div>
                          <div className={`${styles.orderStatus} ${styles[order.status.toLowerCase()]}`}>
                            {order.status === 'COMPLETED' && 'Выполнен'}
                            {order.status === 'PROCESSING' && 'В обработке'}
                            {order.status === 'CANCELLED' && 'Отменен'}
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                ) : (
                  <div className={styles.noOrders}>
                    У вас пока нет заказов
                  </div>
                )}
              </div>
            )}
            
            {/* Favorites Tab */}
            {activeTab === 'favorites' && (
              <div className={styles.favoritesTab}>
                <h2>Избранные игры</h2>
                
                {loadingFavorites ? (
                  <div className="loading">Загрузка избранного...</div>
                ) : favorites && favorites.favoriteItems && favorites.favoriteItems.length > 0 ? (
                  <div className={styles.favoritesGrid}>
                    {/* Favorite games content */}
                  </div>
                ) : (
                  <div className={styles.noFavorites}>
                    У вас пока нет избранных игр
                  </div>
                )}
              </div>
            )}
            
            {/* Library Tab */}
            {activeTab === 'library' && (
              <div className={styles.libraryTab}>
                <h2>Моя библиотека</h2>
                
                {loadingLibrary ? (
                  <div className="loading">Загрузка библиотеки...</div>
                ) : library && library.libraryItems && library.libraryItems.length > 0 ? (
                  <div className={styles.libraryGrid}>
                    {/* Library games content */}
                  </div>
                ) : (
                  <div className={styles.noGames}>
                    В вашей библиотеке пока нет игр
                  </div>
                )}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProfilePage; 