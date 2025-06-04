import { createSelector } from '@reduxjs/toolkit';
import { RootState } from '@/store/types';

// Basic selectors
const selectNewsState = (state: RootState) => state.news;

// Derived selectors
export const selectAllNews = createSelector(
  [selectNewsState],
  (newsState) => newsState.newsList
);

export const selectCurrentNewsItem = createSelector(
  [selectNewsState],
  (newsState) => newsState.newsItem
);

export const selectNewsLoading = createSelector(
  [selectNewsState],
  (newsState) => newsState.loading
);

export const selectNewsError = createSelector(
  [selectNewsState],
  (newsState) => newsState.error
);

// Parameterized selectors
export const selectNewsById = (newsId: number) => 
  createSelector(
    [selectAllNews],
    (news) => news.find(item => item.id === newsId) || null
  );

export const selectNewsByCategory = (category: string) =>
  createSelector(
    [selectAllNews],
    (news) => news.filter(item => item.category === category)
  );

export const selectLatestNews = (count: number = 5) =>
  createSelector(
    [selectAllNews],
    (news) => {
      return [...news]
        .sort((a, b) => new Date(b.publishDate).getTime() - new Date(a.publishDate).getTime())
        .slice(0, count);
    }
  ); 