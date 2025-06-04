export interface User {
  id: number;
  name: string;
  secondName: string;
  thirdName?: string;
  email: string;
  phoneNumber: string;
  paymentInfo?: string;
  roleID: number;
  role: Role;
  token: string;
  refreshToken: string;
}

export interface Role {
  id: number;
  type: string;
  description: string;
}

export interface AuthResponse {
  user: User;
  token: string;
  refreshToken: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  nickname: string;
  email: string;
  password: string;
} 