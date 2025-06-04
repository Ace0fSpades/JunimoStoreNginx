# Этап сборки приложения
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Устанавливаем необходимые зависимости для сборки
RUN apk --no-cache add git

# Копируем только файлы, необходимые для получения зависимостей
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Копируем весь исходный код
COPY . .

# Устанавливаем Swag для генерации документации API
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g ./cmd/app/main.go -o api/docs

# Компилируем приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /build/backend ./cmd/app/main.go

# Создаем структуру директорий и копируем необходимые файлы
RUN mkdir -p /build/api
RUN cp -r /app/api/docs /build/api/
RUN mkdir -p /build/internal
RUN cp -r /app/internal/assets /build/internal/

# Создаем пустой .env файл для godotenv
RUN touch /build/.env

# Этап создания образа для запуска
FROM alpine:3.19

WORKDIR /app

# Устанавливаем зависимости времени выполнения
RUN apk --no-cache add ca-certificates tzdata

# Копируем бинарный файл и Swagger из этапа сборки
COPY --from=builder /build/backend /app/backend
COPY --from=builder /build/api /app/api
COPY --from=builder /build/internal/assets /app/internal/assets
COPY --from=builder /build/.env /app/.env

# Устанавливаем переменные окружения
ENV TZ=UTC

# Запускаем приложение
CMD ["./backend"] 