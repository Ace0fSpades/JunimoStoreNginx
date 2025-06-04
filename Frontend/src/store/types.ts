import { Game, Category, Developer } from '@/models/Game';
import { User } from '@/models/User';
import { Order } from '@/models/Order';
import { Favorite, Library, ShoppingCart } from '@/models/Cart';
import { News } from '@/models/News';

// Состояние приложения
export interface RootState {
  auth: AuthState;
  games: GamesState;
  cart: CartState;
  orders: OrdersState;
  favorites: FavoritesState;
  library: LibraryState;
  news: NewsState;
}

// Auth state
export interface AuthState {
  user: User | null;
  loading: boolean;
  error: string | null;
}

// Games state
export interface GamesState {
  games: Game[];
  game: Game | null;
  categories: Category[];
  developers: Developer[];
  loading: boolean;
  error: string | null;
}

// Cart state
export interface CartState {
  cart: ShoppingCart | null;
  loading: boolean;
  error: string | null;
}

// Orders state
export interface OrdersState {
  orders: Order[];
  loading: boolean;
  error: string | null;
}

// Favorites state
export interface FavoritesState {
  favorites: Favorite | null;
  loading: boolean;
  error: string | null;
}

// Library state
export interface LibraryState {
  library: Library | null;
  loading: boolean;
  error: string | null;
}

// News state
export interface NewsState {
  newsList: News[];
  newsItem: News | null;
  loading: boolean;
  error: string | null;
} 