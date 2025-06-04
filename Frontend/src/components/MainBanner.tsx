import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import styles from '../styles/components/MainBanner.module.scss';
import { FaChevronUp, FaChevronDown } from 'react-icons/fa';
import { shouldUseMocks } from '../config/apiConfig';
import { getMockBannerGames, BannerGame } from '../mocks/mainBanner';
import { GameService } from '../services/GameService';

// Интерфейс для баннера
interface MainBannerProps {}

const MainBanner: React.FC<MainBannerProps> = () => {
  const [slides, setSlides] = useState<BannerGame[]>([]);
  const [currentSlide, setCurrentSlide] = useState(0);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  
  // Загрузка данных для баннера
  useEffect(() => {
    const loadBannerGames = async () => {
      try {
        setLoading(true);
        
        // Проверяем, нужно ли использовать моки
        if (shouldUseMocks()) {
          console.log('MainBanner: Загружаем моки баннера');
          const mockGames = await getMockBannerGames();
          if (mockGames && mockGames.length > 0) {
            console.log('MainBanner: Моки загружены успешно', mockGames.length);
            setSlides(mockGames);
          } else {
            console.error('MainBanner: Моки загружены, но данные пусты');
            setError('Не удалось загрузить данные для баннера');
          }
        } else {
          console.log('MainBanner: Загружаем данные из API');
          // В реальном API мы бы использовали отдельный эндпоинт для баннера
          // Но пока можем использовать популярные игры и добавить к ним поле discount
          const featuredGames = await GameService.getTopSellingGames(3);
          
          if (featuredGames && featuredGames.length > 0) {
            console.log('MainBanner: Данные из API загружены успешно', featuredGames.length);
            // Добавляем скидки и преобразуем в формат BannerGame
            const bannerGames: BannerGame[] = featuredGames.map(game => ({
              ...game,
              discount: Math.floor(Math.random() * 50) + 10, // Временно: случайная скидка
              bannerImageURL: game.imageURL
            }));
            
            setSlides(bannerGames);
          } else {
            console.error('MainBanner: Данные из API загружены, но пусты');
            setError('Не удалось загрузить данные для баннера');
          }
        }
      } catch (err) {
        console.error('MainBanner: Ошибка загрузки данных для баннера:', err);
        setError('Не удалось загрузить данные для баннера');
      } finally {
        setLoading(false);
      }
    };
    
    loadBannerGames();
  }, []);
  
  // Автоматическая прокрутка слайдов каждые 5 секунд
  useEffect(() => {
    if (slides.length === 0) return;
    
    const timer = setInterval(() => {
      setCurrentSlide((prevSlide) => (prevSlide + 1) % slides.length);
    }, 5000);
    
    return () => clearInterval(timer);
  }, [slides.length]);
  
  const nextSlide = () => {
    if (slides.length === 0) return;
    setCurrentSlide((prevSlide) => (prevSlide + 1) % slides.length);
  };
  
  const prevSlide = () => {
    if (slides.length === 0) return;
    setCurrentSlide((prevSlide) => (prevSlide === 0 ? slides.length - 1 : prevSlide - 1));
  };
  
  // Если данные загружаются, показываем заглушку
  if (loading) {
    return (
      <div className={styles.mainBanner} style={{ background: '#1a2634' }}>
        <div className={styles.overlay} />
        <div className="container">
          <div className={styles.content}>
            <h1 className={styles.title}>Загрузка...</h1>
          </div>
        </div>
      </div>
    );
  }
  
  // Если произошла ошибка, показываем сообщение
  if (error || slides.length === 0) {
    return (
      <div className={styles.mainBanner}>
        <div className={styles.overlay} />
        <div className="container">
          <div className={styles.content}>
            <h1 className={styles.title}>{error || 'Нет игр для отображения'}</h1>
          </div>
        </div>
      </div>
    );
  }
  
  const activeSlide = slides[currentSlide];
  const discountedPrice = activeSlide.discount && activeSlide.discount > 0 
    ? Math.round(activeSlide.price * (1 - activeSlide.discount / 100)) 
    : activeSlide.price;
  
  // Выбираем изображение для баннера (предпочитаем специальное изображение для баннера, если оно есть)
  const bannerImage = activeSlide.bannerImageURL || activeSlide.imageURL;

  return (
    <div className={styles.mainBanner} style={{ backgroundImage: `url(${bannerImage})` }}>
      <div className={styles.overlay} />
      
      {/* Вертикальный слайдер */}
      <div className={styles.verticalSlider}>
        <button className={styles.sliderButton} onClick={prevSlide}>
          <FaChevronUp />
        </button>
        <div className={styles.sliderTrack}>
          <div 
            className={styles.sliderThumb} 
            style={{ 
              top: `${(currentSlide / (slides.length - 1)) * 100}%` 
            }}
          ></div>
        </div>
        <button className={styles.sliderButton} onClick={nextSlide}>
          <FaChevronDown />
        </button>
      </div>
      
      <div className="container">
        <div className={styles.content}>
          <h1 className={styles.title}>{activeSlide.title}</h1>
          
          <div className={styles.actions}>
            <Link to={`/cart?add=${activeSlide.id}`} className={styles.cartButton}>
              В КОРЗИНУ
            </Link>
            
            <div className={styles.priceBox}>
              <span className={styles.price}>{discountedPrice}₽</span>
              {activeSlide.discount && activeSlide.discount > 0 && (
                <span className={styles.discount}>-{activeSlide.discount}%</span>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default MainBanner; 