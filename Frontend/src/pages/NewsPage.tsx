import React, { useState, useEffect } from 'react';
import styles from '../styles/pages/NewsPage.module.scss';
import { NewsService } from '../services/NewsService';
import { News } from '../models/News';
import NewsCard from '../components/NewsCard';

const NewsPage: React.FC = () => {
  const [news, setNews] = useState<News[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchNews = async () => {
      try {
        setLoading(true);
        const newsData = await NewsService.getAllNews();
        setNews(newsData);
        setError(null);
      } catch (err) {
        console.error('Error fetching news:', err);
        setError('Не удалось загрузить новости. Пожалуйста, попробуйте позже.');
      } finally {
        setLoading(false);
      }
    };

    fetchNews();
  }, []);

  if (loading) {
    return <div className="loading">Загрузка...</div>;
  }

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  return (
    <div className={styles.newsPage}>
      <div className="container">
        <h1>Новости</h1>
        <div className={styles.newsGrid}>
          {news.map(item => (
            <NewsCard key={item.id} news={item} />
          ))}
        </div>
      </div>
    </div>
  );
};

export default NewsPage; 