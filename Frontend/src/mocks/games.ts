import { Game, Category, Developer } from '../models/Game';

// Получаем базовый путь для GitHub Pages
const isGitHubPages = window.location.hostname.includes('github.io');
const basePath = isGitHubPages ? '/JunimoWebStore' : '';

// Функция для получения корректного пути к изображению
const getImagePath = (path: string): string => {
  if (!path) return '';
  // Если путь уже содержит http/https или data:image, то оставляем как есть
  if (path.startsWith('http') || path.startsWith('data:')) return path;
  // Иначе добавляем базовый путь для GitHub Pages
  return `${basePath}${path}`;
};

// Моки разработчиков
export const mockDevelopers = [
  { id: 1, name: 'CD Projekt RED', country: 'Poland', description: 'Разработчик The Witcher и Cyberpunk 2077', websiteURL: 'https://www.cdprojektred.com/' },
  { id: 2, name: 'FromSoftware', country: 'Japan', description: 'Разработчик Dark Souls и Elden Ring', websiteURL: 'https://www.fromsoftware.jp/' },
  { id: 3, name: 'Rockstar Games', country: 'USA', description: 'Разработчик GTA и Red Dead Redemption', websiteURL: 'https://www.rockstargames.com/' },
  { id: 4, name: 'Mojang Studios', country: 'Sweden', description: 'Разработчик Minecraft', websiteURL: 'https://www.minecraft.net/' },
  { id: 5, name: 'Deep Rock Games', country: 'Denmark', description: 'Разработчик Deep Rock Galactic', websiteURL: 'https://www.ghostship.dk/' },
  { id: 6, name: '4A Games', country: 'Ukraine', description: 'Разработчик серии Metro', websiteURL: 'https://www.4a-games.com/' },
];

// Моки категорий
export const mockCategories = [
  { id: 1, name: 'RPG', description: 'Ролевые игры' },
  { id: 2, name: 'Action', description: 'Экшен игры' },
  { id: 3, name: 'Adventure', description: 'Приключенческие игры' },
  { id: 4, name: 'Shooter', description: 'Шутеры' },
  { id: 5, name: 'Sandbox', description: 'Песочницы' },
  { id: 6, name: 'Survival', description: 'Выживание' },
  { id: 7, name: 'Indie', description: 'Инди игры' },
  { id: 8, name: 'Simulation', description: 'Симуляторы' },
];

// Моки игр
export const mockGames: Game[] = [
  {
    id: 1,
    title: 'Witcher 3: Wild Hunt',
    description: 'Ведьмак 3: Дикая Охота — компьютерная игра в жанре action/RPG, разработанная польской студией CD Projekt RED.',
    price: 2799,
    releaseDate: '2015-05-19',
    developerID: 1,
    developer: mockDevelopers[0],
    categoryID: 1,
    category: mockCategories[0],
    imageURL: getImagePath('/assets/images/witcher3.jpg'),
    createdAt: '2022-01-01',
    updatedAt: '2022-01-01'
  },
  {
    id: 2,
    title: 'Elden Ring',
    description: 'Elden Ring — компьютерная игра в жанре action/RPG с открытым миром, разработанная японской компанией FromSoftware.',
    price: 3999,
    releaseDate: '2022-02-25',
    developerID: 2,
    developer: mockDevelopers[1],
    categoryID: 1,
    category: mockCategories[0],
    imageURL: getImagePath('/assets/images/elden.jpg'),
    createdAt: '2022-01-01',
    updatedAt: '2022-01-01'
  },
  {
    id: 3,
    title: 'Red Dead Redemption 2',
    description: 'Red Dead Redemption 2 — компьютерная игра в жанре action-adventure с открытым миром, разработанная Rockstar Games.',
    price: 3499,
    releaseDate: '2018-10-26',
    developerID: 3,
    developer: mockDevelopers[2],
    categoryID: 3,
    category: mockCategories[2],
    imageURL: getImagePath('/assets/images/rdr2.jpg'),
    createdAt: '2022-01-01',
    updatedAt: '2022-01-01'
  },
  {
    id: 4,
    title: 'Minecraft',
    description: 'Minecraft — компьютерная инди-игра в жанре песочницы, разработанная шведским программистом Маркусом Перссоном.',
    price: 1499,
    releaseDate: '2011-11-18',
    developerID: 4,
    developer: mockDevelopers[3],
    categoryID: 5,
    category: mockCategories[4],
    imageURL: getImagePath('/assets/images/minecraft.jpg'),
    createdAt: '2022-01-01',
    updatedAt: '2022-01-01'
  },
  {
    id: 5,
    title: 'Deep Rock Galactic',
    description: 'Deep Rock Galactic — кооперативная компьютерная игра в жанре шутера от первого лица, разработанная датской студией Ghost Ship Games.',
    price: 1500,
    releaseDate: '2020-05-13',
    developerID: 5,
    developer: mockDevelopers[4],
    categoryID: 4,
    category: mockCategories[3],
    imageURL: getImagePath('/assets/images/deeprock.jpg'),
    createdAt: '2022-01-01',
    updatedAt: '2022-01-01'
  },
  {
    id: 6,
    title: 'Metro Redux',
    description: 'Metro Redux — сборник ремастированных версий компьютерных игр Metro 2033 и Metro: Last Light, разработанный украинской студией 4A Games.',
    price: 1999,
    releaseDate: '2014-08-26',
    developerID: 6,
    developer: mockDevelopers[5],
    categoryID: 4,
    category: mockCategories[3],
    imageURL: getImagePath('/assets/images/metro.jpg'),
    createdAt: '2022-01-01',
    updatedAt: '2022-01-01'
  },
  {
    id: 7,
    title: 'Cyberpunk 2077',
    description: 'Cyberpunk 2077 — компьютерная игра в жанре action/RPG, разработанная польской студией CD Projekt RED.',
    price: 3999,
    releaseDate: '2020-12-10',
    developerID: 1,
    developer: mockDevelopers[0],
    categoryID: 1,
    category: mockCategories[0],
    imageURL: getImagePath('/assets/images/cyberpunk.jpg'),
    createdAt: '2022-01-01',
    updatedAt: '2022-01-01'
  },
  {
    id: 8,
    title: 'Grand Theft Auto V',
    description: 'Grand Theft Auto V — компьютерная игра в жанре action-adventure, разработанная Rockstar North и изданная Rockstar Games.',
    price: 2499,
    releaseDate: '2013-09-17',
    developerID: 3,
    developer: mockDevelopers[2],
    categoryID: 2,
    category: mockCategories[1],
    imageURL: getImagePath('/assets/images/gtav.jpg'),
    createdAt: '2022-01-01',
    updatedAt: '2022-01-01'
  },
  {
    id: 9,
    title: 'Balatro',
    price: 1999,
    description: 'Balatro — карточная игра жанра рогалик, в которой вы создаете комбинации карт в покере, чтобы побеждать противников.',
    releaseDate: '2023-02-20',
    developerID: 7,
    developer: { id: 7, name: 'LocalThunk', country: 'USA', description: 'Инди-разработчик Balatro', websiteURL: 'https://www.playbalatro.com/' },
    categoryID: 7,
    category: mockCategories[6],
    imageURL: getImagePath('/assets/images/balatro.jpg'),
    createdAt: '2023-02-20',
    updatedAt: '2023-02-20'
  }
];

// Функции для получения данных, имитирующие API запросы
export const getMockGames = (): Promise<Game[]> => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve([...mockGames]);
    }, 300); // Имитация задержки сети
  });
};

export const getMockGameById = (id: number): Promise<Game | null> => {
  return new Promise((resolve) => {
    setTimeout(() => {
      const game = mockGames.find(g => g.id === id) || null;
      resolve(game);
    }, 300);
  });
};

export const getMockTopSellingGames = (limit = 4): Promise<Game[]> => {
  return new Promise((resolve) => {
    setTimeout(() => {
      // Случайная сортировка для имитации популярности
      const shuffled = [...mockGames].sort(() => 0.5 - Math.random());
      resolve(shuffled.slice(0, limit));
    }, 300);
  });
};

export const getMockDiscountedGames = (limit = 4): Promise<Game[]> => {
  return new Promise((resolve) => {
    setTimeout(() => {
      // Возвращаем случайные игры, чтобы имитировать скидки
      const shuffled = [...mockGames].sort(() => 0.5 - Math.random());
      resolve(shuffled.slice(0, limit));
    }, 300);
  });
};

// Экспортируем функции для получения моков
export const getMockCategories = async (): Promise<Category[]> => {
  return [...mockCategories];
};

export const getMockDevelopers = async (): Promise<Developer[]> => {
  return [...mockDevelopers];
}; 