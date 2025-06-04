import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import styles from '@/styles/pages/GameDetailsPage.module.scss';
import { FaShoppingCart, FaHeart, FaArrowLeft } from 'react-icons/fa';

// Services
import { CartService } from '@/services/CartService';
import { FavoriteService } from '@/services/FavoriteService';

// Context
import { useAuth } from '@/context/AuthContext';

// Types
import { Review } from '@/models/Game';

// Redux
import { useAppDispatch, useAppSelector } from '@/store/hooks';
import { getGameByIdRequest, clearCurrentGame } from '@/store/slices/gamesSlice';

// Utils
import { createPlaceholderImage } from '@/utils/createPlaceholderImage';

const GameDetailsPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const { isAuthenticated, user } = useAuth();
  const dispatch = useAppDispatch();
  
  // Redux state
  const { game, loading, error } = useAppSelector(state => state.games);
  
  const [reviews] = useState<Review[]>([]);
  const [addedToCart, setAddedToCart] = useState(false);
  const [addedToFavorites, setAddedToFavorites] = useState(false);
  const [imageUrl, setImageUrl] = useState<string>('');

  // Fetch game details on mount
  useEffect(() => {
    if (id) {
      dispatch(getGameByIdRequest(parseInt(id)));
    }
    
    return () => {
      dispatch(clearCurrentGame());
    };
  }, [id, dispatch]);
  
  // Set image URL when game data changes
  useEffect(() => {
    if (game) {
      if (game.image_data && game.image_data.length > 0) {
        // Используем base64 данные для отображения изображения
        setImageUrl(`data:image/png;base64,${game.image_data}`);
      } else {
        setImageUrl(createPlaceholderImage(600, 400, game.title, '#333', '#fff'));
      }
    }
  }, [game]);

  const handleAddToCart = async () => {
    if (!isAuthenticated || !user || !game) {
      window.location.href = '/login';
      return;
    }
    
    try {
      await CartService.addToCart(user.id, game.id);
      setAddedToCart(true);
      setTimeout(() => setAddedToCart(false), 3000);
    } catch (error) {
      console.error('Error adding to cart:', error);
    }
  };

  const handleAddToFavorites = async () => {
    if (!isAuthenticated || !user || !game) {
      window.location.href = '/login';
      return;
    }
    
    try {
      await FavoriteService.addToFavorites(user.id, game.id);
      setAddedToFavorites(true);
      setTimeout(() => setAddedToFavorites(false), 3000);
    } catch (error) {
      console.error('Error adding to favorites:', error);
    }
  };

  const handleImageError = () => {
    // Если изображение не загрузилось, используем placeholder
    if (game) {
      setImageUrl(createPlaceholderImage(600, 400, game.title, '#333', '#fff'));
    }
  };

  if (loading) {
    return <div className="loading">Загрузка...</div>;
  }

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  if (!game) {
    return <div className="error-message">Игра не найдена</div>;
  }

  return (
    <div className={styles.gameDetailsPage}>
      <div className="container">
        <div className={styles.backNav}>
          <Link to="/games" className={styles.backLink}>
            <FaArrowLeft /> Назад к каталогу
          </Link>
        </div>
        
        <div className={styles.gameDetails}>
          <div className={styles.gameMedia}>
            <div className={styles.gameImage}>
              <img 
                src={imageUrl} 
                alt={game.title} 
                onError={handleImageError}
              />
            </div>
          </div>
          
          <div className={styles.gameInfo}>
            <h1 className={styles.gameTitle}>{game.title}</h1>
            
            <div className={styles.gameMeta}>
              <span>{game.category.name}</span>
              <span>{game.developer.name}</span>
              <span>
                {new Date(game.releaseDate).toLocaleDateString('ru-RU')}
              </span>
            </div>
            
            <div className={styles.gamePriceSection}>
              <div className={styles.gamePrice}>{game.price.toFixed(2)}₽</div>
              
              <div className={styles.gameActions}>
                <button
                  className={`btn btn-primary ${styles.addToCartBtn} ${addedToCart ? styles.success : ''}`}
                  onClick={handleAddToCart}
                >
                  <FaShoppingCart /> {addedToCart ? 'Добавлено в корзину' : 'В корзину'}
                </button>
                
                <button
                  className={`btn btn-outline ${styles.addToFavoritesBtn} ${addedToFavorites ? styles.success : ''}`}
                  onClick={handleAddToFavorites}
                >
                  <FaHeart /> {addedToFavorites ? 'Добавлено в избранное' : 'В избранное'}
                </button>
              </div>
            </div>
            
            <div className={styles.gameDescription}>
              <h2>Описание</h2>
              <p>{game.description}</p>
            </div>
            
            <div className={styles.developerInfo}>
              <h2>О разработчике</h2>
              <div className={styles.developerContent}>
                <h3>{game.developer.name}</h3>
                <p>{game.developer.description}</p>
                <p><strong>Страна:</strong> {game.developer.country}</p>
                {game.developer.websiteURL && (
                  <p>
                    <strong>Веб-сайт:</strong>{' '}
                    <a href={game.developer.websiteURL} target="_blank" rel="noopener noreferrer">
                      {game.developer.websiteURL}
                    </a>
                  </p>
                )}
              </div>
            </div>
          </div>
        </div>
        
        <div className={styles.gameReviews}>
          <h2>Отзывы покупателей</h2>
          
          {reviews.length === 0 ? (
            <div className={styles.noReviews}>
              Пока нет отзывов для этой игры. Будьте первым, кто оставит отзыв!
            </div>
          ) : (
            <div className={styles.reviewsList}>
              {reviews.map((review) => (
                <div key={review.id} className={styles.reviewItem}>
                  <div className={styles.reviewHeader}>
                    <div className={styles.reviewAuthor}>{review.user.name}</div>
                    <div className={styles.reviewRating}>Оценка: {review.rating}/5</div>
                    <div className={styles.reviewDate}>
                      {new Date(review.createdAt).toLocaleDateString('ru-RU')}
                    </div>
                  </div>
                  <div className={styles.reviewTitle}>{review.title}</div>
                  <div className={styles.reviewText}>{review.description}</div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default GameDetailsPage; 