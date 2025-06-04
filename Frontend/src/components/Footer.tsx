import React from 'react';
import { Link } from 'react-router-dom';
import styles from '../styles/components/Footer.module.scss';
import { FaVk, FaTelegram, FaTwitter, FaInstagram } from 'react-icons/fa';

const Footer: React.FC = () => {
  const currentYear = new Date().getFullYear();

  return (
    <footer className={styles.footer}>
      <div className="container">
        <div className={styles.footerContent}>
          <div className={`${styles.footerSection} ${styles.footerLogo}`}>
            <h3>Junim</h3>
            <p>Ваш онлайн-магазин цифровых игр с лучшими ценами и широким выбором.</p>
            <div className={styles.socialLinks}>
              <a href="https://vk.com/" target="_blank" rel="noopener noreferrer" title="ВКонтакте">
                <FaVk />
              </a>
              <a href="https://t.me/" target="_blank" rel="noopener noreferrer" title="Telegram">
                <FaTelegram />
              </a>
              <a href="https://twitter.com/" target="_blank" rel="noopener noreferrer" title="Twitter">
                <FaTwitter />
              </a>
              <a href="https://www.instagram.com/" target="_blank" rel="noopener noreferrer" title="Instagram">
                <FaInstagram />
              </a>
            </div>
          </div>

          <div className={styles.footerSection}>
            <h3>Информация</h3>
            <ul className={styles.footerLinks}>
              <li><Link to="/about">О нас</Link></li>
              <li><Link to="/news">Новости</Link></li>
              <li><Link to="/support">Поддержка</Link></li>
              <li><Link to="/terms">Условия использования</Link></li>
              <li><Link to="/privacy">Политика конфиденциальности</Link></li>
            </ul>
          </div>

          <div className={styles.footerSection}>
            <h3>Контакты</h3>
            <p>Email: support@junim.com</p>
            <p>Телефон: +7 (123) 456-7890</p>
            <p>Адрес: г. Москва, ул. Игровая, 42</p>
          </div>
        </div>

        <div className={styles.footerBottom}>
          <p>&copy; {currentYear} Junim. Все права защищены.</p>
        </div>
      </div>
    </footer>
  );
};

export default Footer; 