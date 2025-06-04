import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import styles from '../styles/components/GameCard.module.scss';
import { FaShoppingCart, FaHeart } from 'react-icons/fa';
import { Game } from '../models/Game';
import { useAuth } from '../context/AuthContext';
import { CartService } from '../services/CartService';
import { FavoriteService } from '../services/FavoriteService';
import { createPlaceholderImage } from '../utils/createPlaceholderImage';

interface GameCardProps {
  game: Game;
}

const GameCard: React.FC<GameCardProps> = ({ game }) => {
  const { isAuthenticated, user } = useAuth();
  const [imageUrl, setImageUrl] = useState<string>(`data:image/png;base64,${game.image_data}`);
  const [discount, setDiscount] = useState<number>(0); // Пример скидки, в реальности должно приходить с бэкенда
  
  // Расчет скидочной цены
  const originalPrice = game.price;
  const discountedPrice = discount > 0 
    ? Math.round(originalPrice * (1 - discount / 100)) 
    : originalPrice;

  // Формируем URL изображения из base64 данных
  useEffect(() => {
    // Для демонстрации - устанавливаем случайную скидку для некоторых игр
    if (Math.random() > 0.5) {
      setDiscount(Math.floor(Math.random() * 50) + 10); // Скидка от 10% до 60%
    }
    
    if (game.image_data && game.image_data.length > 0) {
      // Используем base64 данные для отображения изображения
      console.log(`data:image/png;base64,${game.image_data}`);
      setImageUrl(`data:image/png;base64,${game.image_data}`);
    } else {
      // Создаем placeholder изображение с названием игры
      setImageUrl(createPlaceholderImage(300, 300, game.title, '#333', '#fff'));
    }
  }, [game]);
  
  const handleImageError = () => {
    // Если изображение не загрузилось, используем placeholder
    setImageUrl(createPlaceholderImage(300, 300, game.title, '#333', '#fff'));
  };

  const addToCart = async (e: React.MouseEvent) => {
    e.preventDefault();
    
    if (!isAuthenticated || !user) {
      window.location.href = '/login';
      return;
    }
    
    try {
      await CartService.addToCart(user.id, game.id);
      // You could add a toast notification here
    } catch (error) {
      console.error('Error adding to cart:', error);
    }
  };

  const addToFavorites = async (e: React.MouseEvent) => {
    e.preventDefault();
    
    if (!isAuthenticated || !user) {
      window.location.href = '/login';
      return;
    }
    
    try {
      await FavoriteService.addToFavorites(user.id, game.id);
      // You could add a toast notification here
    } catch (error) {
      console.error('Error adding to favorites:', error);
    }
  };

  return (
    <div className={styles.gameCard}>
      <Link to={`/games/${game.id}`} className={styles.cardLink}>
        <div className={styles.imageContainer}>
          <img 
            src={imageUrl} 
            alt={game.title}
            onError={handleImageError} 
            className={styles.image}
          />
          {discount > 0 && (
            <span className={styles.discount}>-{discount}%</span>
          )}
          <div className={styles.actions}>
            <button 
              className={styles.actionButton}
              onClick={addToCart}
              title="Добавить в корзину"
              aria-label="Добавить в корзину"
            >
              <FaShoppingCart />
            </button>
            <button 
              className={styles.actionButton}
              onClick={addToFavorites}
              title="Добавить в избранное"
              aria-label="Добавить в избранное"
            >
              <FaHeart />
            </button>
          </div>
        </div>
        <div className={styles.cardContent}>
          <h3 className={styles.title}>{game.title}</h3>
          <div className={styles.priceContainer}>
            {discount > 0 ? (
              <>
                <span className={styles.oldPrice}>{originalPrice}₽</span>
                <span className={styles.price}>{discountedPrice}₽</span>
              </>
            ) : (
              <span className={styles.price}>{originalPrice}₽</span>
            )}
          </div>
        </div>
      </Link>
    </div>
  );
};

export default GameCard; 