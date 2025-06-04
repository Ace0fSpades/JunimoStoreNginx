import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux';
import type { RootState } from '@/store/types';
import type { AppDispatch } from '@/store/index';

// Типизированные хуки для использования в компонентах
export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector; 