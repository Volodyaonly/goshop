# ---------- ЭТАП 1: СБОРКА ----------
FROM golang:1.25.7 AS builder

WORKDIR /app

# Сначала копируем только зависимости
COPY go.mod go.sum ./

RUN go mod download

# Теперь копируем исходный код
COPY . .

# Собираем Linux-бинарник GoShop
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/goshop ./cmd/main.go


# ---------- ЭТАП 2: ЗАПУСК ----------
FROM debian:bookworm-slim

WORKDIR /app

# Забираем только готовый бинарник
COPY --from=builder /out/goshop ./goshop

EXPOSE 8080

CMD ["./goshop"]