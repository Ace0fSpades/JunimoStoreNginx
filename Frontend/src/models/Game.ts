export interface Game {
  id: number;
  title: string;
  description: string;
  price: number;
  releaseDate: string;
  developerID: number;
  developer: Developer;
  categoryID: number;
  category: Category;
  image_data?: string;
  image_name?: string;
  imageURL?: string;
  createdAt: string;
  updatedAt: string;
}

export interface Developer {
  id: number;
  name: string;
  country: string;
  description: string;
  websiteURL: string;
}

export interface Category {
  id: number;
  name: string;
  description: string;
}

export interface Review {
  id: number;
  gameID: number;
  userID: number;
  title: string;
  description: string;
  rating: number;
  createdAt: string;
  user: {
    name: string;
    secondName: string;
  };
} 