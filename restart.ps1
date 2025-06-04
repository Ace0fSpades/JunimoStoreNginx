# Создаем файл .env с правильным портом 9090
@"
APP_ENV=development
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gamestore
PORT=9090
"@ | Out-File -FilePath ".env" -Encoding utf8

# Запускаем скрипт сборки бэкенда
Write-Host "Запускаем сборку бэкенда..."
./build-backend.ps1

# Останавливаем все контейнеры (если не остановлены)
Write-Host "Останавливаем все контейнеры..."
docker-compose down

# Удаляем образы фронтенда и бэкенда, чтобы пересобрать их
Write-Host "Удаляем старые образы..."
docker image rm -f game-store-app-frontend game-store-app-backend-1 game-store-app-backend-2 game-store-app-backend-3 game-store-app-nginx

# Запускаем контейнеры с флагом --build
Write-Host "Запускаем контейнеры..."
docker-compose up -d --build

# Выводим логи
Write-Host "Вывод логов для проверки..."
docker-compose logs 