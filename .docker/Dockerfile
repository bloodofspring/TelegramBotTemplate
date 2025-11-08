# syntax=docker/dockerfile:1.4

# --- Build stage ---
FROM golang:latest AS builder
WORKDIR /build
COPY app/go.mod app/go.sum ./
RUN go mod download
COPY app/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/bot ./cmd/main.go

# --- Production image ---
FROM alpine:3.19 AS prod
WORKDIR /app
COPY --from=builder /build/bot /app/

# Добавляем ca-certificates для HTTPS запросов
RUN apk --no-cache add ca-certificates

ENV GIN_MODE=release
# HEALTHCHECK --interval=30s --timeout=5s --start-period=30s CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1
ENTRYPOINT ["/app/bot"]

# --- Development image ---
FROM golang:latest AS dev
WORKDIR /app

# Установка Air для hot reload
RUN go install github.com/air-verse/air@latest

# Установка ca-certificates для HTTPS запросов
RUN apk --no-cache add ca-certificates

COPY app/go.mod app/go.sum ./
RUN go mod download

# Копируем остальные файлы
COPY app/ ./

# Air будет запущен из docker-compose с примонтированным volume
CMD ["air", "-c", "../.air.toml"]