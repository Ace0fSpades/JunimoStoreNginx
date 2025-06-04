import React from 'react';
import styles from '../styles/pages/AboutPage.module.scss';
import JunimoLogo from '../components/JunimoLogo';

const AboutPage: React.FC = () => {
  return (
    <div className={styles.aboutPage}>
      <div className="container">
        <div className={styles.aboutContent}>
          <section className={styles.section}>
            <h2>О нас</h2>
            <p>
              Компания Junimo - это инновационная компания, специализирующаяся на разработке 
              интерактивных развлекательных продуктов и решений для геймеров всех возрастов. 
              Мы стремимся создавать уникальный опыт, который объединяет людей и дарит радость.
            </p>
            <p>
              Основанная в 2010 году группой энтузиастов, наша компания выросла из небольшой 
              инди-студии в крупного игрока на рынке развлечений. Мы гордимся нашими креативными 
              решениями и постоянным стремлением к инновациям.
            </p>
          </section>

          <section className={styles.section}>
            <h2>Наша миссия</h2>
            <p>
              Мы стремимся создавать игры и развлекательные продукты, которые не только увлекают, 
              но и объединяют людей, способствуют развитию творческого мышления и 
              предоставляют уникальный опыт каждому пользователю.
            </p>
          </section>

          <section className={styles.section}>
            <h2>Наши ценности</h2>
            <div className={styles.valuesList}>
              <div className={styles.valueItem}>
                <JunimoLogo color="#FF5733" size={48} />
                <h3>Креативность</h3>
                <p>Мы верим в силу креативного мышления и постоянно ищем новые идеи и подходы.</p>
              </div>
              <div className={styles.valueItem}>
                <JunimoLogo color="#33A1FF" size={48} />
                <h3>Качество</h3>
                <p>Мы стремимся к превосходству во всем, что делаем, уделяя внимание каждой детали.</p>
              </div>
              <div className={styles.valueItem}>
                <JunimoLogo color="#FFD133" size={48} />
                <h3>Сообщество</h3>
                <p>Мы ценим наше сообщество игроков и разработчиков, и всегда учитываем их мнение.</p>
              </div>
              <div className={styles.valueItem}>
                <JunimoLogo color="#33FF57" size={48} />
                <h3>Инновации</h3>
                <p>Мы постоянно исследуем новые технологии и методы для создания уникальных продуктов.</p>
              </div>
            </div>
          </section>

          <section className={styles.section}>
            <h2>Наша команда</h2>
            <p>
              В Junimo работают талантливые специалисты из разных областей: программисты, дизайнеры, 
              художники, писатели, звукорежиссеры и многие другие. Все они объединены страстью 
              к созданию удивительных игровых миров и опыта.
            </p>
            <div className={styles.teamStatistics}>
              <div className={styles.statItem}>
                <span className={styles.statNumber}>50+</span>
                <span className={styles.statLabel}>Сотрудников</span>
              </div>
              <div className={styles.statItem}>
                <span className={styles.statNumber}>15+</span>
                <span className={styles.statLabel}>Проектов</span>
              </div>
              <div className={styles.statItem}>
                <span className={styles.statNumber}>8</span>
                <span className={styles.statLabel}>Лет опыта</span>
              </div>
            </div>
          </section>

          <section className={styles.contactSection}>
            <h2>Свяжитесь с нами</h2>
            <div className={styles.contactInfo}>
              <div className={styles.contactItem}>
                <h3>Адрес</h3>
                <p>г. Москва, ул. Игровая, д. 42</p>
              </div>
              <div className={styles.contactItem}>
                <h3>Телефон</h3>
                <p>+7 (123) 456-78-90</p>
              </div>
              <div className={styles.contactItem}>
                <h3>Email</h3>
                <p>info@junimo.ru</p>
              </div>
            </div>
          </section>
        </div>
      </div>
    </div>
  );
};

export default AboutPage; 