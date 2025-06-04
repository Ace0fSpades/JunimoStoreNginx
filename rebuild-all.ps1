Write-Host "=== Запуск полной пересборки проекта ===" -ForegroundColor Cyan

# 1. Сборка и запуск
Write-Host "1. Запускаем сборку и запуск всех контейнеров..." -ForegroundColor Green

# Останавливаем и удаляем старые контейнеры и образы
Write-Host "   Останавливаем и удаляем старые контейнеры..." -ForegroundColor Gray
docker-compose down
docker image rm -f game-store-app-frontend game-store-app-backend-1 game-store-app-backend-2 game-store-app-backend-3 game-store-app-nginx 2>$null

# Запускаем сборку и запуск контейнеров
docker-compose up -d --build

# 2. Проверка статуса
Write-Host "2. Проверяем статус контейнеров..." -ForegroundColor Green
docker-compose ps

Write-Host "`n=== Проект успешно собран и запущен ===" -ForegroundColor Cyan
Write-Host "Фронтенд: http://localhost:3000" -ForegroundColor Yellow
Write-Host "API (через nginx): http://localhost:9090/swagger/index.html" -ForegroundColor Yellow
Write-Host "HTTPS: https://localhost" -ForegroundColor Yellow
