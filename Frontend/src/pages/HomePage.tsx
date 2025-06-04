import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import styles from '../styles/pages/HomePage.module.scss';
import { FaArrowRight } from 'react-icons/fa';

// Components
import MainBanner from '../components/MainBanner';
import GameCard from '../components/GameCard';
import NewsCard from '../components/NewsCard';

// Services
import { GameService } from '../services/GameService';
import { NewsService } from '../services/NewsService';
import { Game } from '../models/Game';
import { News } from '../models/News';

const HomePage: React.FC = () => {
  const [topSellingGames, setTopSellingGames] = useState<Game[]>([]);
  const [discountedGames, setDiscountedGames] = useState<Game[]>([]);
  const [latestNews, setLatestNews] = useState<News[]>([]);
  const [loading, setLoading] = useState({
    topSelling: true,
    discounted: true,
    news: true
  });
  const [errors, setErrors] = useState({
    topSelling: null as string | null,
    discounted: null as string | null,
    news: null as string | null
  });

  // Загрузка данных при монтировании компонента
  useEffect(() => {
    console.log('HomePage: Начало загрузки данных');
    const loadData = async () => {
      try {
        // Загрузка хитов продаж
        setLoading(prev => ({ ...prev, topSelling: true }));
        try {
          console.log('HomePage: Загрузка хитов продаж');
          const topGames = await GameService.getAllGames();
          if (topGames && topGames.length > 0) {
            console.log('HomePage: Хиты продаж загружены успешно', topGames.length);
            setTopSellingGames(topGames);
            setErrors(prev => ({ ...prev, topSelling: null }));
          } else {
            console.error('HomePage: Хиты продаж загружены, но данные пусты');
            setErrors(prev => ({ ...prev, topSelling: 'Не удалось загрузить хиты продаж' }));
          }
        } catch (err) {
          console.error('HomePage: Ошибка загрузки хитов продаж:', err);
          setErrors(prev => ({ ...prev, topSelling: 'Ошибка загрузки хитов продаж' }));
        } finally {
          setLoading(prev => ({ ...prev, topSelling: false }));
        }
        
        // Загрузка игр со скидками
        setLoading(prev => ({ ...prev, discounted: true }));
        try {
          console.log('HomePage: Загрузка игр со скидками');
          const discGames = await GameService.getDiscountedGames(4);
          if (discGames && discGames.length > 0) {
            console.log('HomePage: Игры со скидками загружены успешно', discGames.length);
            setDiscountedGames(discGames);
            setErrors(prev => ({ ...prev, discounted: null }));
          } else {
            console.error('HomePage: Игры со скидками загружены, но данные пусты');
            setErrors(prev => ({ ...prev, discounted: 'Не удалось загрузить игры со скидками' }));
          }
        } catch (err) {
          console.error('HomePage: Ошибка загрузки игр со скидками:', err);
          setErrors(prev => ({ ...prev, discounted: 'Ошибка загрузки игр со скидками' }));
        } finally {
          setLoading(prev => ({ ...prev, discounted: false }));
        }
        
        // Загрузка новостей
        setLoading(prev => ({ ...prev, news: true }));
        try {
          console.log('HomePage: Загрузка новостей');
          const newsData = await NewsService.getLatestNews(3);
          
          if (newsData && newsData.length > 0) {
            console.log('HomePage: Новости загружены успешно', newsData.length);
            setLatestNews(newsData);
            setErrors(prev => ({ ...prev, news: null }));
          } else {
            console.error('HomePage: Новости загружены, но данные пусты');
            setErrors(prev => ({ ...prev, news: 'Не удалось загрузить новости' }));
          }
        } catch (err) {
          console.error('HomePage: Ошибка загрузки новостей:', err);
          setErrors(prev => ({ ...prev, news: 'Ошибка загрузки новостей' }));
        } finally {
          setLoading(prev => ({ ...prev, news: false }));
        }
      } catch (error) {
        console.error('HomePage: Общая ошибка загрузки данных для главной страницы:', error);
      }
    };
    
    loadData();
  }, []);

  return (
    <div className={styles.homePage}>
      {/* Главный баннер со слайдером */}
      <MainBanner />

      <div className="container">
        {/* Секция хитов продаж */}
        <section className={styles.gamesSection}>
          <div className={styles.sectionHeader}>
            <h2>ХИТЫ ПРОДАЖ</h2>
            <Link to="/games" className={styles.viewAll}>Все игры <FaArrowRight /></Link>
          </div>
          
          <div className={styles.gamesGrid}>
            {loading.topSelling ? (
              <div className={styles.loading}>Загрузка...</div>
            ) : errors.topSelling ? (
              <div className={styles.error}>{errors.topSelling}</div>
            ) : topSellingGames.length === 0 ? (
              <div className={styles.empty}>Нет данных для отображения</div>
            ) : (
              topSellingGames.map(game => (
                <GameCard
                  key={game.id}
                  game={game}
                />
              ))
            )}
          </div>
        </section>

        {/* Секция скидок */}
        <section className={styles.gamesSection}>
          <div className={styles.sectionHeader}>
            <h2>СКИДКИ</h2>
            <Link to="/games?discounted=true" className={styles.viewAll}>Все скидки <FaArrowRight /></Link>
          </div>
          
          <div className={styles.gamesGrid}>
            {loading.discounted ? (
              <div className={styles.loading}>Загрузка...</div>
            ) : errors.discounted ? (
              <div className={styles.error}>{errors.discounted}</div>
            ) : discountedGames.length === 0 ? (
              <div className={styles.empty}>Нет данных для отображения</div>
            ) : (
              discountedGames.map(game => (
                <GameCard
                  key={game.id}
                  game={game}
                />
              ))
            )}
          </div>
        </section>

        {/* Секция новостей */}
        <section className={styles.newsSection}>
          <div className={styles.sectionHeader}>
            <h2>НОВОСТИ</h2>
            <Link to="/news" className={styles.viewAll}>Все новости <FaArrowRight /></Link>
          </div>
          
          <div className={styles.newsGrid}>
            {loading.news ? (
              <div className={styles.loading}>Загрузка...</div>
            ) : errors.news ? (
              <div className={styles.error}>{errors.news}</div>
            ) : latestNews.length === 0 ? (
              <div className={styles.empty}>Нет данных для отображения</div>
            ) : (
              latestNews.map(news => (
                <NewsCard
                  key={news.id}
                  news={news}
                />
              ))
            )}
          </div>
        </section>
      </div>
    </div>
  );
};

export default HomePage; 