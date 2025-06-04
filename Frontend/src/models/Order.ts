import { Game } from './Game';

export interface Order {
  id: number;
  userID: number;
  totalCost: number;
  status: string;
  createdAt: string;
  updatedAt: string;
  orderItems: OrderItem[];
}

export interface OrderItem {
  id: number;
  orderID: number;
  gameID: number;
  game: Game;
  price: number;
  quantity: number;
} 