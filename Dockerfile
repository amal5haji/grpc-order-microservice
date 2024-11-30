FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./orders-service ./cmd/server/main.go
RUN go build -o ./gateway-service ./cmd/client/main.go

FROM alpine:latest AS orders-service
WORKDIR /app
COPY --from=builder /app/orders-service .
EXPOSE 8080
ENTRYPOINT ["./orders-service"]

FROM alpine:latest AS gateway-service
WORKDIR /app
COPY --from=builder /app/gateway-service .
EXPOSE 3000
ENTRYPOINT ["./gateway-service"]