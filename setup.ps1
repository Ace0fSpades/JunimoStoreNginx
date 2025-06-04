# Копирование переменных окружения из Backend/env в .env
Copy-Item -Path "Backend/.env" -Destination ".env"

# Создание директории для конфигурации nginx
if (-not (Test-Path -Path "nginx")) {
    New-Item -ItemType Directory -Path "nginx"
}

# Создание директории для SSL сертификатов
$sslPath = "nginx/ssl"
if (-not (Test-Path -Path $sslPath)) {
    New-Item -ItemType Directory -Path $sslPath
}

# Генерация самоподписанных SSL сертификатов, если их еще нет
$certPath = "$sslPath/nginx.crt"
$keyPath = "$sslPath/nginx.key"

if (-not (Test-Path -Path $certPath) -or -not (Test-Path -Path $keyPath)) {
    Write-Host "Генерация самоподписанных SSL сертификатов..."
    
    # Использование OpenSSL для генерации сертификатов
    # Убедитесь, что OpenSSL установлен в вашей системе
    
    # Создание приватного ключа и самоподписанного сертификата
    & openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout $keyPath -out $certPath -subj "/CN=localhost"
    
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Ошибка при генерации SSL сертификатов. Убедитесь, что OpenSSL установлен."
        Write-Host "Вы можете установить OpenSSL с помощью команды: choco install openssl"
        Write-Host "Или вручную создать сертификаты и поместить их в папку $sslPath"
    } else {
        Write-Host "SSL сертификаты успешно созданы."
    }
}

Write-Host "Подготовка окружения завершена. Запустите docker-compose up -d для запуска контейнеров." 