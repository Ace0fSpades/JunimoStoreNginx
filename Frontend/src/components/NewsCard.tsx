import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import styles from '../styles/components/NewsCard.module.scss';
import { News } from '../models/News';
import { createNewsPlaceholderImage } from '../utils/createPlaceholderImage';

interface NewsCardProps {
  news: News;
}

const NewsCard: React.FC<NewsCardProps> = ({ news }) => {
  const [imageError, setImageError] = useState(false);

  // Format date
  const formattedDate = new Date(news.publishDate).toLocaleDateString('ru-RU', {
    day: 'numeric',
    month: 'long',
    year: 'numeric'
  });

  const handleImageError = () => {
    setImageError(true);
  };

  const imageSource = imageError || !news.imageURL ? 
    createNewsPlaceholderImage(400, 225, 'Новости') : 
    news.imageURL;

  return (
    <div className={styles.newsCard}>
      <Link to={`/news/${news.id}`} className={styles.cardLink}>
        <div className={styles.imageContainer}>
          <img 
            src={imageSource} 
            alt={news.title} 
            onError={handleImageError}
            className={styles.image}
          />
          <div className={styles.date}>{formattedDate}</div>
        </div>
        <div className={styles.cardContent}>
          <h3 className={styles.title}>{news.title}</h3>
          <p className={styles.summary}>
            {news.content.length > 120 
              ? `${news.content.substring(0, 120)}...` 
              : news.content
            }
          </p>
        </div>
      </Link>
    </div>
  );
};

export default NewsCard; 