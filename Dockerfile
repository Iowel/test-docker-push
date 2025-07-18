# Этап сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копируем модули и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники из папки app
COPY app ./app

# Сборка бинарника из каталога app
RUN cd app && go build -o ../server .

# Финальный минимальный образ
FROM debian:bookworm-slim

COPY --from=builder /server /server

EXPOSE 8080

ENTRYPOINT ["/server"]
