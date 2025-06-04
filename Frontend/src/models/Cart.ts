import { Game } from './Game';

export interface ShoppingCart {
  id: number;
  userID: number;
  cartItems: CartItem[];
}

export interface CartItem {
  id: number;
  shoppingCartID: number;
  gameID: number;
  game: Game;
  quantity: number;
}

export interface Favorite {
  id: number;
  userID: number;
  favoriteItems: FavoriteItem[];
}

export interface FavoriteItem {
  id: number;
  favoriteID: number;
  gameID: number;
  game: Game;
}

export interface Library {
  id: number;
  userID: number;
  libraryItems: LibraryItem[];
}

export interface LibraryItem {
  id: number;
  libraryID: number;
  gameID: number;
  game: Game;
} 