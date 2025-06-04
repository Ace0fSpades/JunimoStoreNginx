import React, { useState, useEffect } from 'react';
import { useParams, useLocation } from 'react-router-dom';
import styles from '@/styles/pages/CatalogPage.module.scss';
import { FaTimes, FaChevronDown, FaChevronUp } from 'react-icons/fa';
import RangeSlider from 'react-range-slider-input';
import 'react-range-slider-input/dist/style.css';

// Services
import { Game } from '@/models/Game';

// Redux
import { useAppDispatch, useAppSelector } from '@/store/hooks';
import { 
  selectAllGames,
  selectCategories, 
  selectDevelopers,
  selectGamesLoading
} from '@/store/selectors/gamesSelectors';
import { 
  getAllGamesRequest, 
  getCategoriesRequest, 
  getDevelopersRequest 
} from '@/store/slices/gamesSlice';

// Components
import GameCard from '@/components/GameCard';

const CatalogPage: React.FC = () => {
  const { categoryId } = useParams<{ categoryId: string }>();
  const location = useLocation();
  const dispatch = useAppDispatch();
  
  // Получаем данные из Redux
  const allGames = useAppSelector(selectAllGames);
  const categories = useAppSelector(selectCategories);
  const developers = useAppSelector(selectDevelopers);
  const loading = useAppSelector(selectGamesLoading);
  
  const [filteredGames, setFilteredGames] = useState<Game[]>([]);
  const [error] = useState<string | null>(null);
  
  // Filter states
  const [priceRange, setPriceRange] = useState<{ min: number; max: number }>({
    min: 0,
    max: 5000
  });
  const [minMaxPrices, setMinMaxPrices] = useState<[number, number]>([0, 5000]);
  const [priceRangeDisplay, setPriceRangeDisplay] = useState<[string, string]>(['0', '5000']);
  const [selectedCategories, setSelectedCategories] = useState<number[]>([]);
  const [selectedDevelopers, setSelectedDevelopers] = useState<number[]>([]);
  const [sortBy, setSortBy] = useState<string>('newest');
  const [filtersOpen, setFiltersOpen] = useState(false);
  const [expandedSections, setExpandedSections] = useState({
    categories: true,
    developers: true,
    price: true
  });

  // Загружаем данные при монтировании компонента
  useEffect(() => {
    dispatch(getAllGamesRequest({ limit: 100, offset: 0 }));
    dispatch(getCategoriesRequest());
    dispatch(getDevelopersRequest());
  }, [dispatch]);

  // Обновляем выбранные категории при изменении categoryId из URL
  useEffect(() => {
    if (categoryId) {
      const catId = parseInt(categoryId);
      setSelectedCategories([catId]);
    } else {
      // Если мы на странице всех игр, сбрасываем выбранные категории
      setSelectedCategories([]);
    }
  }, [categoryId, location.pathname]);

  // Обновляем диапазон цен при изменении списка игр
  useEffect(() => {
    if (allGames.length > 0) {
      const prices = allGames.map(game => game.price);
      const minPrice = Math.min(...prices);
      const maxPrice = Math.max(...prices);
      setMinMaxPrices([minPrice, maxPrice]);
      setPriceRange({ min: minPrice, max: maxPrice });
      setPriceRangeDisplay([minPrice.toString(), maxPrice.toString()]);
    }
    
    // При изменении URL или обновлении страницы сбрасываем остальные фильтры
    if (categoryId) {
      setSelectedDevelopers([]);
      setSortBy('newest');
    } else {
      resetFilters();
    }
  }, [allGames, categoryId, location.pathname]);
  
  // Apply filters when filter states change
  useEffect(() => {
    applyFilters();
  }, [allGames, selectedCategories, selectedDevelopers, priceRange, sortBy]);
  
  const applyFilters = () => {
    let result = [...allGames];
    
    // Filter by categories
    if (selectedCategories.length > 0) {
      result = result.filter(game => selectedCategories.includes(game.category.id));
    }
    
    // Filter by developers
    if (selectedDevelopers.length > 0) {
      result = result.filter(game => selectedDevelopers.includes(game.developer.id));
    }
    
    // Filter by price range
    result = result.filter(
      game => game.price >= priceRange.min && game.price <= priceRange.max
    );
    
    // Sort results
    switch (sortBy) {
      case 'newest':
        result.sort((a, b) => new Date(b.releaseDate).getTime() - new Date(a.releaseDate).getTime());
        break;
      case 'oldest':
        result.sort((a, b) => new Date(a.releaseDate).getTime() - new Date(b.releaseDate).getTime());
        break;
      case 'price_low':
        result.sort((a, b) => a.price - b.price);
        break;
      case 'price_high':
        result.sort((a, b) => b.price - a.price);
        break;
      case 'alphabetical':
        result.sort((a, b) => a.title.localeCompare(b.title));
        break;
      default:
        break;
    }
    
    setFilteredGames(result);
  };
  
  const handleCategoryChange = (categoryId: number) => {
    if (selectedCategories.includes(categoryId)) {
      setSelectedCategories(selectedCategories.filter(id => id !== categoryId));
    } else {
      setSelectedCategories([...selectedCategories, categoryId]);
    }
  };
  
  const handleDeveloperChange = (developerId: number) => {
    if (selectedDevelopers.includes(developerId)) {
      setSelectedDevelopers(selectedDevelopers.filter(id => id !== developerId));
    } else {
      setSelectedDevelopers([...selectedDevelopers, developerId]);
    }
  };
  
  // Обработчик ввода цены вручную
  const handlePriceChange = (e: React.ChangeEvent<HTMLInputElement>, index: number) => {
    try {
      // Конвертируем введенное значение в число
      let value = e.target.value;
      
      // Обновляем отображаемое значение
      const newPriceRangeDisplay = [...priceRangeDisplay] as [string, string];
      newPriceRangeDisplay[index] = value;
      setPriceRangeDisplay(newPriceRangeDisplay);
      
      // Проверяем, что введено корректное число и оно не отрицательное
      let newValue = parseFloat(value);
      if (isNaN(newValue) || newValue < 0) return;
      
      // Обеспечиваем, чтобы значение не было отрицательным
      newValue = Math.max(0, newValue);
      
      // Обновляем фактический ценовой диапазон
      const newPriceRange = { ...priceRange };
      
      if (index === 0) {
        // Нижний предел не может быть больше верхнего
        newPriceRange.min = Math.min(newValue, priceRange.max);
      } else {
        // Верхний предел не может быть меньше нижнего
        newPriceRange.max = Math.max(newValue, priceRange.min);
      }
      
      setPriceRange(newPriceRange);
    } catch (err) {
      console.error('Ошибка при вводе цены:', err);
    }
  };
  
  const resetFilters = () => {
    setSelectedCategories(categoryId ? [parseInt(categoryId)] : []);
    setSelectedDevelopers([]);
    setPriceRange({ min: minMaxPrices[0], max: minMaxPrices[1] });
    setPriceRangeDisplay([minMaxPrices[0].toString(), minMaxPrices[1].toString()]);
    setSortBy('newest');
  };
  
  const toggleFilterSection = (section: keyof typeof expandedSections) => {
    setExpandedSections({
      ...expandedSections,
      [section]: !expandedSections[section]
    });
  };
  
  // Обработчик изменения значений слайдера
  const handleRangeSliderChange = (values: number[]) => {
    // Округляем значения до целых чисел
    const roundedValues: [number, number] = [
      Math.round(values[0]),
      Math.round(values[1])
    ];
    
    // Обновляем состояние диапазона цен
    setPriceRange({ min: roundedValues[0], max: roundedValues[1] });
    setPriceRangeDisplay([
      roundedValues[0].toString(),
      roundedValues[1].toString()
    ]);
  };
  
  if (loading) {
    return <div className="loading">Загрузка...</div>;
  }
  
  if (error) {
    return <div className="error-message">{error}</div>;
  }

  return (
    <div className={styles.catalogPage}>
      <div className="container">
        <div className={styles.catalogHeader}>
          <h1>Каталог игр</h1>
          <div className={styles.catalogControls}>
            <div className={styles.sortControl}>
              <label htmlFor="sort-select">Сортировать по:</label>
              <select 
                id="sort-select" 
                value={sortBy} 
                onChange={(e) => setSortBy(e.target.value)}
              >
                <option value="newest">Новые</option>
                <option value="oldest">Старые</option>
                <option value="price_low">Цена (по возрастанию)</option>
                <option value="price_high">Цена (по убыванию)</option>
                <option value="alphabetical">По алфавиту</option>
              </select>
            </div>
          </div>
        </div>
        
        <div className={styles.catalogContent}>
          <aside className={`${styles.catalogFilters} ${filtersOpen ? styles.open : ''}`}>
            <div className={styles.filtersHeader}>
              <h2>Фильтры</h2>
              <button 
                className={styles.closeFiltersButton}
                onClick={() => setFiltersOpen(false)}
              >
                <FaTimes />
              </button>
            </div>
            
            <div className="filters-content">
              {/* Categories Filter */}
              <div className={styles.filterSection}>
                <div 
                  className={styles.filterSectionHeader}
                  onClick={() => toggleFilterSection('categories')}
                >
                  <h3>Категории</h3>
                  {expandedSections.categories ? <FaChevronUp /> : <FaChevronDown />}
                </div>
                {expandedSections.categories && (
                  <div className={styles.filterSectionContent}>
                    <div className={styles.filterItems}>
                      {categories.map(category => (
                        <div key={category.id} className={styles.filterItem}>
                          <label>
                            <input 
                              type="checkbox" 
                              checked={selectedCategories.includes(category.id)} 
                              onChange={() => handleCategoryChange(category.id)}
                            />
                            <span>{category.name}</span>
                          </label>
                        </div>
                      ))}
                    </div>
                  </div>
                )}
              </div>
              
              {/* Developers Filter */}
              <div className={styles.filterSection}>
                <div 
                  className={styles.filterSectionHeader}
                  onClick={() => toggleFilterSection('developers')}
                >
                  <h3>Разработчики</h3>
                  {expandedSections.developers ? <FaChevronUp /> : <FaChevronDown />}
                </div>
                {expandedSections.developers && (
                  <div className={styles.filterSectionContent}>
                    <div className={styles.filterItems}>
                      {developers.map(developer => (
                        <div key={developer.id} className={styles.filterItem}>
                          <label>
                            <input 
                              type="checkbox" 
                              checked={selectedDevelopers.includes(developer.id)} 
                              onChange={() => handleDeveloperChange(developer.id)}
                            />
                            <span>{developer.name}</span>
                          </label>
                        </div>
                      ))}
                    </div>
                  </div>
                )}
              </div>
              
              {/* Price Range Filter */}
              <div className={styles.filterSection}>
                <div 
                  className={styles.filterSectionHeader}
                  onClick={() => toggleFilterSection('price')}
                >
                  <h3>Цена</h3>
                  {expandedSections.price ? <FaChevronUp /> : <FaChevronDown />}
                </div>
                {expandedSections.price && (
                  <div className={styles.filterSectionContent}>
                    <div className={styles.priceInputs}>
                      <div className={styles.priceInputGroup}>
                        <label>От</label>
                        <input 
                          type="text" 
                          value={priceRangeDisplay[0]} 
                          onChange={(e) => handlePriceChange(e, 0)}
                        />
                      </div>
                      <div className={styles.priceInputGroup}>
                        <label>До</label>
                        <input 
                          type="text" 
                          value={priceRangeDisplay[1]} 
                          onChange={(e) => handlePriceChange(e, 1)}
                        />
                      </div>
                    </div>
                    
                    <div className={styles.priceSlider}>
                      <RangeSlider
                        min={minMaxPrices[0]}
                        max={minMaxPrices[1]}
                        step={1}
                        value={[priceRange.min, priceRange.max]}
                        onInput={handleRangeSliderChange}
                        className={styles.rangeSlider}
                      />
                    </div>
                  </div>
                )}
              </div>
            </div>
            
            <div className={styles.filtersActions}>
              <button 
                className={styles.resetFiltersBtn}
                onClick={resetFilters}
              >
                Сбросить
              </button>
              <button 
                className={styles.applyFiltersBtn}
                onClick={() => setFiltersOpen(false)}
              >
                Применить
              </button>
            </div>
          </aside>
          
          <section className={styles.catalogGames}>
            {filteredGames.length === 0 ? (
              <div className={styles.noGames}>
                <h2>Игры не найдены</h2>
                <p>Попробуйте изменить параметры фильтрации</p>
                <button 
                  className="btn btn-primary"
                  onClick={resetFilters}
                >
                  Сбросить фильтры
                </button>
              </div>
            ) : (
              <>
                <div className={styles.gamesGrid}>
                  {filteredGames.map(game => (
                    <GameCard key={game.id} game={game} />
                  ))}
                </div>
                <div className={styles.gamesCount}>
                  Всего игр: {filteredGames.length} из {allGames.length}
                </div>
              </>
            )}
          </section>
        </div>
      </div>
    </div>
  );
};

export default CatalogPage; 