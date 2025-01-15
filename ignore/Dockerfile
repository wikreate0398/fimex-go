# Используем официальный образ Go для сборки
FROM golang:1.23.4 as builder

# Устанавливаем рабочую директорию для сборки
WORKDIR /go/src/github.com/wikreate0398/fimex-go/

# Копируем файлы в рабочую директорию
COPY .. .

# Загружаем зависимости
RUN go mod download

# Собираем бинарник для Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./cmd/main.go

# Используем минимальный образ Alpine для финального контейнера
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированный бинарник из образа builder
COPY --from=builder /go/bin/app .

# Команда для запуска приложения
CMD ["./app"]