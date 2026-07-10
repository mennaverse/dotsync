FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o dotsync-backend .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dotsync-backend ./

COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/seeders /app/seeders

EXPOSE 8067

CMD ["./dotsync-backend", "serve", "-b", "0.0.0.0:8067"]
