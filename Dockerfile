# Stage 1: Build
FROM golang:1.22.0-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin .

RUN chmod +x /app/bin

ENTRYPOINT ["/app/bin"]
