# Используйте базовый образ Golang
FROM golang:latest

# Устанавливаем зависимости
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -tags netgo -ldflags '-s -w' -o app

# Запускаем приложение
CMD ["./app"]