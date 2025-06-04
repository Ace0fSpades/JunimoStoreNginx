import React, { useState, useRef, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '@/context/AuthContext';
import styles from '@/styles/components/Header.module.scss';
import { FaSearch, FaShoppingCart, FaHeart, FaUser, FaTimes } from 'react-icons/fa';
import JunimoLogo from './JunimoLogo';
import { GameService } from '@/services/GameService';
import { Game } from '@/models/Game';

// Redux
import { useAppDispatch, useAppSelector } from '@/store/hooks';
import { selectCategories, selectGamesLoading } from '@/store/selectors/gamesSelectors';
import { getCategoriesRequest } from '@/store/slices/gamesSlice';

// Убираем статичный массив категорий
// const gameCategories = [
//   { id: "action", name: "Экшен" },
//   { id: "adventure", name: "Приключения" },
//   { id: "rpg", name: "Ролевые" },
//   { id: "strategy", name: "Стратегии" },
//   { id: "simulation", name: "Симуляторы" },
//   { id: "sports", name: "Спортивные" },
//   { id: "racing", name: "Гонки" },
//   { id: "indie", name: "Инди" }
// ];

const Header: React.FC = () => {
  const { isAuthenticated, logout } = useAuth();
  const dispatch = useAppDispatch();
  const categories = useAppSelector(selectCategories);
  const loading = useAppSelector(selectGamesLoading);
  
  const navigate = useNavigate();
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
  const [searchQuery, setSearchQuery] = useState('');
  const [catalogMenuOpen, setCatalogMenuOpen] = useState(false);
  const catalogMenuRef = useRef<HTMLDivElement>(null);
  const [searchResults, setSearchResults] = useState<Game[]>([]);
  const [showSearchResults, setShowSearchResults] = useState(false);
  const [isSearching, setIsSearching] = useState(false);
  const searchRef = useRef<HTMLDivElement>(null);

  // Загрузка категорий при монтировании компонента
  useEffect(() => {
    dispatch(getCategoriesRequest());
  }, [dispatch]);

  // Поиск игр при вводе
  useEffect(() => {
    const searchTimeout = setTimeout(async () => {
      if (searchQuery.trim().length >= 2) {
        setIsSearching(true);
        try {
          const results = await GameService.searchGames(searchQuery.trim(), 5);
          setSearchResults(results);
          setShowSearchResults(true);
        } catch (error) {
          console.error('Ошибка при поиске:', error);
        } finally {
          setIsSearching(false);
        }
      } else {
        setSearchResults([]);
        setShowSearchResults(false);
      }
    }, 300);

    return () => clearTimeout(searchTimeout);
  }, [searchQuery]);

  // Закрываем выпадающее меню при клике вне его
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (catalogMenuRef.current && !catalogMenuRef.current.contains(event.target as Node)) {
        setCatalogMenuOpen(false);
      }
      if (searchRef.current && !searchRef.current.contains(event.target as Node)) {
        setShowSearchResults(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  const handleLogout = () => {
    logout();
    navigate('/');
    setMobileMenuOpen(false);
  };

  const toggleMobileMenu = () => {
    setMobileMenuOpen(!mobileMenuOpen);
  };

  const toggleCatalogMenu = () => {
    setCatalogMenuOpen(!catalogMenuOpen);
  };

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    if (searchQuery.trim()) {
      navigate(`/games/search?query=${encodeURIComponent(searchQuery.trim())}`);
      setShowSearchResults(false);
    }
  };

  const handleSearchInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchQuery(e.target.value);
  };

  const clearSearch = () => {
    setSearchQuery('');
    setSearchResults([]);
    setShowSearchResults(false);
  };

  const handleSearchResultClick = (gameId: number) => {
    navigate(`/games/${gameId}`);
    setShowSearchResults(false);
    setSearchQuery('');
  };

  return (
    <header className={styles.header}>
      <div className="container">
        <div className={styles.headerContent}>
          <div className={styles.logo}>
            <Link to="/">
              <span className={styles.logoText}>Junim</span>
              <span className={styles.logoIcon}><JunimoLogo strokeColor="white" size={30} /></span>
            </Link>
          </div>

          <div className={styles.catalogLink} ref={catalogMenuRef}>
            <button 
              className={styles.catalogButton} 
              onClick={toggleCatalogMenu}
            >
              КАТАЛОГ ИГР
            </button>
            {catalogMenuOpen && (
              <div className={styles.catalogDropdown}>
                <Link to="/games" onClick={() => setCatalogMenuOpen(false)}>
                  Все игры
                </Link>
                <div className={styles.divider}></div>
                {loading ? (
                  <div className={styles.loading}>Загрузка...</div>
                ) : (
                  categories.map(category => (
                    <Link 
                      key={category.id} 
                      to={`/games/category/${category.id}`} 
                      onClick={() => setCatalogMenuOpen(false)}
                    >
                      {category.name}
                    </Link>
                  ))
                )}
              </div>
            )}
          </div>

          <div className={styles.searchBar} ref={searchRef}>
            <form onSubmit={handleSearch}>
              <div className={styles.searchInputWrapper}>
                <input
                  type="text"
                  placeholder="Поиск игр..."
                  value={searchQuery}
                  onChange={handleSearchInputChange}
                />
                {searchQuery && (
                  <button
                    type="button"
                    className={styles.clearSearchButton}
                    onClick={clearSearch}
                  >
                    <FaTimes />
                  </button>
                )}
                <button type="submit" className={styles.searchButton}>
                  <FaSearch />
                </button>
              </div>
            </form>
            
            {showSearchResults && (
              <div className={styles.searchResultsDropdown}>
                {isSearching ? (
                  <div className={styles.searchingIndicator}>Поиск...</div>
                ) : searchResults?.length > 0 ? (
                  <>
                    <div className={styles.searchResultsList}>
                      {searchResults.map(game => (
                        <div 
                          key={game.id} 
                          className={styles.searchResultItem}
                          onClick={() => handleSearchResultClick(game.id)}
                        >
                          {game.image_data && (
                            <div className={styles.searchResultThumb}>
                              <img src={`data:image/png;base64,${game.image_data}`} alt={game.title} />
                            </div>
                          )}
                          <div className={styles.searchResultInfo}>
                            <div className={styles.searchResultTitle}>{game.title}</div>
                            <div className={styles.searchResultPrice}>{game.price} ₽</div>
                          </div>
                        </div>
                      ))}
                    </div>
                    <div className={styles.searchResultsFooter}>
                      <Link 
                        to={`/games/search?query=${encodeURIComponent(searchQuery.trim())}`}
                        onClick={() => setShowSearchResults(false)}
                      >
                        Все результаты ({searchResults.length}+)
                      </Link>
                    </div>
                  </>
                ) : (
                  <div className={styles.noSearchResults}>
                    По запросу "{searchQuery}" ничего не найдено
                  </div>
                )}
              </div>
            )}
          </div>

          <nav className={`${styles.nav} ${mobileMenuOpen ? styles.open : ''}`}>
            <ul className={styles.navLinks}>
              {isAuthenticated ? (
                <>
                  <li>
                    <Link to="/cart" className={styles.iconLink} title="Корзина">
                      <FaShoppingCart />
                    </Link>
                  </li>
                  <li>
                    <Link to="/favorites" className={styles.iconLink} title="Избранное">
                      <FaHeart />
                    </Link>
                  </li>
                  <li>
                    <Link to="/profile" className={styles.iconLink} title="Личный кабинет">
                      <FaUser />
                    </Link>
                  </li>
                  <li>
                    <button className={styles.loginButton} onClick={handleLogout}>
                      ВЫЙТИ
                    </button>
                  </li>
                </>
              ) : (
                <>
                  <li>
                    <Link to="/cart" className={styles.iconLink} title="Корзина">
                      <FaShoppingCart />
                    </Link>
                  </li>
                  <li>
                    <Link to="/login" className={styles.loginButton}>
                      ВОЙТИ
                    </Link>
                  </li>
                  <li>
                    <Link to="/register" className={styles.loginButton}>
                      РЕГИСТРАЦИЯ
                    </Link>
                  </li>
                </>
              )}
            </ul>
          </nav>

          <button 
            className={styles.mobileMenuButton} 
            onClick={toggleMobileMenu}
            aria-label="Toggle menu"
          >
            <span className={`${styles.hamburger} ${mobileMenuOpen ? styles.open : ''}`}></span>
          </button>
        </div>
      </div>
    </header>
  );
};

export default Header; 