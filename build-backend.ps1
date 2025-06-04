# Останавливаем все контейнеры
Write-Host "Останавливаем все контейнеры..."
docker-compose down

# Удаляем старый образ для сборки бэкенда
Write-Host "Удаляем старый образ для сборки бэкенда..."
docker image rm -f backend-builder

# Создаем директорию для артефактов, если ее нет
if (!(Test-Path -Path "./Backend/build")) {
    New-Item -ItemType Directory -Path "./Backend/build" -Force
}

# Собираем бинарный файл бэкенда
Write-Host "Собираем бинарный файл бэкенда..."
docker build -f ./Backend/Dockerfile.build -t backend-builder ./Backend

# Создаем временный контейнер для извлечения артефактов
Write-Host "Извлекаем артефакты из образа..."
docker create --name temp-backend-builder backend-builder
docker cp temp-backend-builder:/build/. ./Backend/build/
docker rm temp-backend-builder

Write-Host "Бинарный файл бэкенда успешно собран и размещен в директории ./Backend/build" 