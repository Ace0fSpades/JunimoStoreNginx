import { Routes, Route, Navigate } from 'react-router-dom';

// Pages
import {
  HomePage,
  CatalogPage,
  GameDetailsPage,
  CartPage,
  LoginPage,
  RegisterPage,
  ProfilePage,
  NewsPage,
  NewsDetailsPage,
  FavoritesPage,
  NotFoundPage,
  SearchResultsPage,
  AboutPage
} from '../pages';

// Components
import ProtectedRoute from './ProtectedRoute';

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/games" element={<CatalogPage />} />
      <Route path="/games/category/:categoryId" element={<CatalogPage />} />
      <Route path="/games/search" element={<SearchResultsPage />} />
      <Route path="/games/:id" element={<GameDetailsPage />} />
      <Route path="/news" element={<NewsPage />} />
      <Route path="/news/:id" element={<NewsDetailsPage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/cart" element={<CartPage />} />
      <Route path="/favorites" element={
        <ProtectedRoute>
          <FavoritesPage />
        </ProtectedRoute>
      } />
      <Route path="/profile" element={
        <ProtectedRoute>
          <ProfilePage />
        </ProtectedRoute>
      } />
      <Route path="/about" element={<AboutPage />} />
      <Route path="/404" element={<NotFoundPage />} />
      <Route path="*" element={<Navigate to="/404" replace />} />
    </Routes>
  );
};

export default AppRoutes; 