import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import styles from '../styles/pages/FavoritesPage.module.scss';
import { FaTrash } from 'react-icons/fa';

// Services
import { FavoriteService } from '../services/FavoriteService';

// Context
import { useAuth } from '../context/AuthContext';

// Types
import { Favorite } from '../models/Cart';

const FavoritesPage: React.FC = () => {
  const { user } = useAuth();
  const [favorites, setFavorites] = useState<Favorite | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (user) {
      fetchFavorites();
    } else {
      setLoading(false);
    }
  }, [user]);

  const fetchFavorites = async () => {
    try {
      setLoading(true);
      const favoritesData = await FavoriteService.getFavorites(user!.id);
      setFavorites(favoritesData);
      setError(null);
    } catch (err) {
      console.error('Error fetching favorites:', err);
      setError('Не удалось загрузить избранное. Пожалуйста, попробуйте позже.');
    } finally {
      setLoading(false);
    }
  };

  const handleRemoveFromFavorites = async (gameId: number) => {
    try {
      await FavoriteService.removeFromFavorites(user!.id, gameId);
      
      // Update local state
      setFavorites(prev => {
        if (!prev) return null;
        
        return {
          ...prev,
          favoriteItems: prev.favoriteItems.filter(item => item.game.id !== gameId)
        };
      });
    } catch (err) {
      console.error('Error removing from favorites:', err);
      setError('Не удалось удалить игру из избранного. Пожалуйста, попробуйте позже.');
    }
  };

  const handleClearFavorites = async () => {
    try {
      await FavoriteService.clearFavorites(user!.id);
      setFavorites(prev => prev ? { ...prev, favoriteItems: [] } : null);
    } catch (err) {
      console.error('Error clearing favorites:', err);
      setError('Не удалось очистить избранное. Пожалуйста, попробуйте позже.');
    }
  };

  if (loading) {
    return <div className="loading">Загрузка...</div>;
  }

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  return (
    <div className={styles.favoritesPage}>
      <div className="container">
        <div className={styles.favoritesHeader}>
          <h1>Избранное</h1>
          {favorites && favorites.favoriteItems.length > 0 && (
            <button className="btn btn-outline" onClick={handleClearFavorites}>
              Очистить избранное
            </button>
          )}
        </div>

        {!favorites || favorites.favoriteItems.length === 0 ? (
          <div className={styles.emptyFavorites}>
            <h2>В избранном пока нет игр</h2>
            <p>Добавьте игры из каталога в избранное, чтобы они отображались здесь.</p>
            <Link to="/games" className="btn btn-primary">Перейти в каталог</Link>
          </div>
        ) : (
          <div className={styles.favoritesList}>
            {favorites.favoriteItems.map(item => (
              <div key={item.id} className={styles.favoriteItem}>
                <div className={styles.favoriteImage}>
                  <img src={item.game.imageURL} alt={item.game.title} />
                </div>
                
                <div className={styles.favoriteInfo}>
                  <h3>
                    <Link to={`/games/${item.game.id}`}>{item.game.title}</Link>
                  </h3>
                  
                  <div className={styles.favoriteMeta}>
                    <span>{item.game.category.name}</span>
                    <span>{item.game.developer.name}</span>
                  </div>
                  
                  <div className={styles.favoriteDescription}>
                    {item.game.description.length > 150
                      ? `${item.game.description.substring(0, 150)}...`
                      : item.game.description
                    }
                  </div>
                </div>
                
                <div className={styles.favoriteActions}>
                  <div className={styles.favoritePrice}>{item.game.price.toFixed(2)}₽</div>
                  
                  <div className={styles.favoriteButtons}>
                    <Link to={`/games/${item.game.id}`} className="btn btn-primary">
                      Подробнее
                    </Link>
                    
                    <button 
                      className={`btn btn-outline ${styles.removeBtn}`}
                      onClick={() => handleRemoveFromFavorites(item.game.id)}
                      title="Удалить из избранного"
                    >
                      <FaTrash />
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default FavoritesPage; 