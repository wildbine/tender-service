# Используем базовый образ для Go
FROM golang:1.19

# Создаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum и загруженные зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код в рабочую директорию
COPY . .

# Сборка приложения
RUN go build -o main ./cmd/server

# Указываем команду для запуска приложения
CMD ["./main"]