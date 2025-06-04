import { News } from '../models/News';

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

// Моки новостей
export const mockNews: News[] = [
  {
    id: 1,
    title: 'Анонс новой части The Elder Scrolls',
    content: 'Компания Bethesda анонсировала новую часть легендарной серии The Elder Scrolls. Ожидается, что игра выйдет в 2025 году. В новой части игроки отправятся в провинцию Хаммерфелл, и смогут исследовать обширный открытый мир с новыми подземельями, заданиями и возможностями.',
    imageURL: getImagePath('/assets/images/news/tes.jpg'),
    publishDate: '2023-06-15',
    author: 'Иван Игровой',
    category: 'Анонсы',
    tags: ['The Elder Scrolls', 'Bethesda', 'RPG', 'Анонс']
  },
  {
    id: 2,
    title: 'Обновление Cyberpunk 2077 добавит новые квесты',
    content: 'CD Projekt RED объявила о выходе крупного обновления для Cyberpunk 2077, которое добавит новые сюжетные задания и улучшит производительность. Обновление также включает множество исправлений багов и новую систему транспорта.',
    imageURL: getImagePath('/assets/images/news/cyberpunk.jpg'),
    publishDate: '2023-06-10',
    author: 'Мария Геймерова',
    category: 'Обновления',
    tags: ['Cyberpunk 2077', 'CD Projekt RED', 'Патч', 'Обновление']
  },
  {
    id: 3,
    title: 'Sony представила PlayStation 6',
    content: 'Sony официально анонсировала PlayStation 6. Консоль следующего поколения обещает революционные возможности и выйдет уже в следующем году. Новая консоль будет поддерживать разрешение 8K и частоту кадров до 240 FPS, а также будет оснащена новой системой трассировки лучей.',
    imageURL: getImagePath('/assets/images/news/ps6.jpg'),
    publishDate: '2023-06-05',
    author: 'Петр Консольный',
    category: 'Железо',
    tags: ['PlayStation', 'Sony', 'Консоли', 'Новое поколение']
  },
  {
    id: 4,
    title: 'Новая часть GTA находится в разработке',
    content: 'Rockstar Games подтвердила, что новая часть серии Grand Theft Auto находится в активной разработке. Хотя подробности пока не раскрываются, компания обещает, что это будет самая амбициозная игра в истории студии.',
    imageURL: getImagePath('/assets/images/news/gta.jpg'),
    publishDate: '2023-05-25',
    author: 'Алексей Рокстарович',
    category: 'Анонсы',
    tags: ['GTA', 'Rockstar Games', 'Анонс', 'Разработка']
  },
  {
    id: 5,
    title: 'Valve представила новую версию Steam Deck',
    content: 'Компания Valve представила обновленную версию портативной игровой консоли Steam Deck с улучшенным дисплеем, более мощным процессором и увеличенным временем автономной работы. Продажи начнутся в следующем месяце.',
    imageURL: getImagePath('/assets/images/news/steamdeck.jpg'),
    publishDate: '2023-05-20',
    author: 'Елена Вальвова',
    category: 'Железо',
    tags: ['Steam Deck', 'Valve', 'Портативные консоли', 'Железо']
  }
];

// Функции для получения данных, имитирующие API запросы
export const getMockNews = (): Promise<News[]> => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve([...mockNews]);
    }, 300); // Имитация задержки сети
  });
};

export const getMockNewsById = (id: number): Promise<News | null> => {
  return new Promise((resolve) => {
    setTimeout(() => {
      const news = mockNews.find(n => n.id === id) || null;
      resolve(news);
    }, 300);
  });
};

export const getMockLatestNews = (limit = 3): Promise<News[]> => {
  return new Promise((resolve) => {
    setTimeout(() => {
      // Сортировка по дате публикации (от новых к старым)
      const sorted = [...mockNews].sort((a, b) => {
        return new Date(b.publishDate).getTime() - new Date(a.publishDate).getTime();
      });
      resolve(sorted.slice(0, limit));
    }, 300);
  });
}; 