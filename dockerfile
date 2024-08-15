FROM golang:alpine3.20 AS builder

WORKDIR /app/telegram_bot

RUN apk add --no-cache make && \
    mkdir -p /app/telegram_bot

COPY telegram_bot/ .

RUN    go mod download && \
    make build

FROM alpine:3.20

WORKDIR /app/telegram_bot

COPY --from=builder /app/telegram_bot/telegram_bot .

EXPOSE 8080

CMD ["./telegram_bot"]
