FROM node:18-alpine AS build

WORKDIR /app

# Включаем Corepack для управления версиями Yarn
RUN corepack enable && corepack prepare yarn@4.1.1 --activate

# Копируем только необходимые файлы для установки зависимостей
COPY package.json yarn.lock .yarnrc.yml ./
COPY .yarn ./.yarn

# Устанавливаем зависимости
RUN yarn install

# Устанавливаем недостающие зависимости
RUN yarn add react-range-slider-input

# Устанавливаем переменные окружения
ENV VITE_BASE_URL=/
ENV VITE_BACKEND_PORT=9090
ENV VITE_API_MODE=api

# Копируем исходный код, но исключаем node_modules
COPY src ./src
COPY public ./public
COPY index.html ./
COPY tsconfig.json ./
COPY tsconfig.node.json ./
COPY vite.config.ts ./
COPY postcss.config.js ./

# Собираем приложение
RUN yarn build

# Запускаем в production режиме
FROM node:18-alpine

WORKDIR /app

# Устанавливаем serve для раздачи статических файлов
RUN yarn global add serve

# Копируем собранные файлы из предыдущего этапа
COPY --from=build /app/build ./build

# Запускаем сервер
CMD ["serve", "-s", "build", "-l", "3000"] 