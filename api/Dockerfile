FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
COPY .env /.env
RUN CGO_ENABLED=0 GOOS=linux go build -o api-server

FROM alpine:latest AS production
COPY --from=builder /app .
EXPOSE 8888
ENTRYPOINT ["./api-server"]