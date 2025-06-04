/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_REACT_APP_API_URL: string;
  readonly VITE_API_MODE: 'api' | 'mocks';
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
} 