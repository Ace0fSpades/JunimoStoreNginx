import React from 'react';
import { Link } from 'react-router-dom';
import styles from '../styles/pages/NotFoundPage.module.scss';
import { FaHome, FaGamepad } from 'react-icons/fa';

const NotFoundPage: React.FC = () => {
  return (
    <div className={styles.notFoundPage}>
      <div className="container">
        <div className={styles.notFoundContent}>
          <div className={styles.errorCode}>404</div>
          <h1>Страница не найдена</h1>
          <p>
            Извините, страница, которую вы ищете, не существует или была перемещена.
          </p>
          <div className={styles.notFoundActions}>
            <Link to="/" className="btn btn-primary">
              <FaHome /> На главную
            </Link>
            <Link to="/games" className="btn btn-outline">
              <FaGamepad /> В каталог
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default NotFoundPage;
