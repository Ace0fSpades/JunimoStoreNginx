import { Game } from '../models/Game';

// Структура для баннера, содержит Game плюс дополнительные поля для баннера
export interface BannerGame extends Game {
  discount?: number; // Процент скидки
  bannerImageURL?: string; // Специальное большое изображение для баннера
}

// Моки для баннера
export const mockBannerGames: BannerGame[] = [
  {
    id: 1,
    title: 'WITCHER 3',
    description: 'Ведьмак 3: Дикая Охота — компьютерная игра в жанре action/RPG, разработанная польской студией CD Projekt RED.',
    price: 2799,
    discount: 87,
    releaseDate: '2015-05-19',
    developerID: 1,
    developer: {
      id: 1, 
      name: 'CD Projekt RED', 
      country: 'Poland', 
      description: 'Разработчик The Witcher и Cyberpunk 2077', 
      websiteURL: 'https://www.cdprojektred.com/'
    },
    categoryID: 1,
    category: {
      id: 1,
      name: 'RPG',
      description: 'Ролевые игры'
    },
    imageURL: '/assets/images/witcher3.jpg',
    bannerImageURL: '/assets/images/banners/witcher3_banner.jpg',
    createdAt: '2022-01-01',
    updatedAt: '2022-01-01'
  },
  {
    id: 2,
    title: 'ELDEN RING',
    description: 'Elden Ring — компьютерная игра в жанре action/RPG с открытым миром, разработанная японской компанией FromSoftware.',
    price: 3999,
    discount: 20,
    releaseDate: '2022-02-25',
    developerID: 2,
    developer: {
      id: 2,
      name: 'FromSoftware',
      country: 'Japan',
      description: 'Разработчик Dark Souls и Elden Ring',
      websiteURL: 'https://www.fromsoftware.jp/'
    },
    categoryID: 1,
    category: {
      id: 1,
      name: 'RPG',
      description: 'Ролевые игры'
    },
    imageURL: '/assets/images/elden.jpg',
    bannerImageURL: '/assets/images/banners/elden_banner.jpg',
    createdAt: '2022-01-01',
    updatedAt: '2022-01-01'
  },
  {
    id: 3,
    title: 'RED DEAD REDEMPTION 2',
    description: 'Red Dead Redemption 2 — компьютерная игра в жанре action-adventure с открытым миром, разработанная Rockstar Games.',
    price: 3499,
    discount: 35,
    releaseDate: '2018-10-26',
    developerID: 3,
    developer: {
      id: 3,
      name: 'Rockstar Games',
      country: 'USA',
      description: 'Разработчик GTA и Red Dead Redemption',
      websiteURL: 'https://www.rockstargames.com/'
    },
    categoryID: 3,
    category: {
      id: 3,
      name: 'Adventure',
      description: 'Приключенческие игры'
    },
    imageURL: '/assets/images/rdr2.jpg',
    bannerImageURL: '/assets/images/banners/rdr2_banner.jpg',
    createdAt: '2022-01-01',
    updatedAt: '2022-01-01'
  }
];

// Функция для получения данных для баннера
export const getMockBannerGames = (): Promise<BannerGame[]> => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve([...mockBannerGames]);
    }, 300); // Имитация задержки сети
  });
}; 