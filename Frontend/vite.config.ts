import { defineConfig, loadEnv } from 'vite';
import react from '@vitejs/plugin-react';
import viteTsconfigPaths from 'vite-tsconfig-paths';
import svgr from 'vite-plugin-svgr';
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  // Загружаем переменные окружения в зависимости от режима
  const env = loadEnv(mode, process.cwd(), 'VITE_');
  
  // Используем базовый путь из переменной окружения, если она установлена, 
  // или из package.json для GitHub Pages
  const base = env.VITE_BASE_URL || '/JunimoWebStore/';
  
  return {
    // Базовый путь, который можно переопределить через переменную окружения
    base,
    // Определяем переменные, которые должны быть доступны в клиентском коде
    define: {
      'process.env.NODE_ENV': JSON.stringify(mode),
      // Legacy переменные для обратной совместимости
      'process.env.REACT_APP_API_URL': JSON.stringify(env.VITE_REACT_APP_API_URL),
      // Новая переменная для режима API
      'process.env.API_MODE': JSON.stringify(env.VITE_API_MODE)
    },
    plugins: [react(), viteTsconfigPaths(), svgr()],
    resolve: {
      alias: {
        '@': path.resolve(__dirname, './src')
      }
    },
    server: {
      port: 3000,
      open: true,
      proxy: {
        '/api/v1': {
          target: 'http://localhost:9090',
          changeOrigin: true,
          secure: false,
          ws: true,
          configure: (proxy, _options) => {
            proxy.on('error', (err, _req, _res) => {
              console.log('Прокси ошибка:', err);
            });
            proxy.on('proxyReq', (proxyReq, req, _res) => {
              console.log('Отправка запроса к:', req.url);
            });
            proxy.on('proxyRes', (proxyRes, req, _res) => {
              console.log('Получен ответ от:', req.url, 'статус:', proxyRes.statusCode);
            });
          }
        }
      },
    },
    css: {
      modules: {
        localsConvention: 'camelCase',
        generateScopedName: '[name]__[local]___[hash:base64:5]'
      },
      preprocessorOptions: {
        scss: {
          // Global variables are now imported using @use in each file
        }
      }
    },
    build: {
      outDir: 'build',
      sourcemap: true,
      // Убедимся, что JS файлы генерируются правильно
      rollupOptions: {
        output: {
          entryFileNames: 'assets/[name]-[hash].js',
          chunkFileNames: 'assets/[name]-[hash].js',
          assetFileNames: 'assets/[name]-[hash].[ext]'
        }
      }
    },
    // PostCSS configuration is loaded from postcss.config.js
  };
}); 