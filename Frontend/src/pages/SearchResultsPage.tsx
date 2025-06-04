import React, { useState, useEffect } from 'react';
import { useSearchParams, Link } from 'react-router-dom';
import styles from '../styles/pages/SearchResultsPage.module.scss';
import { FaSearch, FaArrowLeft } from 'react-icons/fa';

// Services
import { GameService } from '../services/GameService';

// Types
import { Game } from '../models/Game';

// Components
import GameCard from '../components/GameCard';

const SearchResultsPage: React.FC = () => {
  const [searchParams] = useSearchParams();
  const query = searchParams.get('query') || '';
  
  const [results, setResults] = useState<Game[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const searchGames = async () => {
      try {
        setLoading(true);
        
        if (!query) {
          setResults([]);
          return;
        }
        
        const searchResults = await GameService.searchGames(query);
        setResults(searchResults);
        setError(null);
      } catch (err) {
        console.error('Error searching games:', err);
        setError('Не удалось выполнить поиск. Пожалуйста, попробуйте позже.');
      } finally {
        setLoading(false);
      }
    };

    searchGames();
  }, [query]);

  if (loading) {
    return <div className="loading">Загрузка...</div>;
  }

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  return (
    <div className={styles.searchResultsPage}>
      <div className="container">
        <div className={styles.searchHeader}>
          <h1>Результаты поиска</h1>
          <div className={styles.searchInfo}>
            <span className={styles.searchQuery}>«{query}»</span>
            <span className={styles.resultsCount}>Найдено игр: {results.length}</span>
          </div>
          <Link to="/games" className={styles.backLink}>
            <FaArrowLeft /> Вернуться в каталог
          </Link>
        </div>

        {results.length === 0 ? (
          <div className={styles.noResults}>
            <div className={styles.noResultsIcon}>
              <FaSearch />
            </div>
            <h2>Ничего не найдено</h2>
            <p>К сожалению, по вашему запросу ничего не найдено. Попробуйте изменить запрос или просмотреть каталог.</p>
            <Link to="/games" className="btn btn-primary">
              Перейти в каталог
            </Link>
          </div>
        ) : (
          <div className={styles.searchResults}>
            <div className={styles.gamesGrid}>
              {results.map(game => (
                <GameCard key={game.id} game={game} />
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default SearchResultsPage;