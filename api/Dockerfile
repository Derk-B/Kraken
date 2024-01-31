FROM golang:1.21-alpine AS builder
ARG GIT_COMMIT
ENV GIT_COMMIT=$GIT_COMMIT
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o api-server -ldflags="-X 'kraken/api-server/controllers.ApiVersion=$GIT_COMMIT'"

FROM alpine:latest AS production
COPY --from=builder /app .
EXPOSE 8888
ENTRYPOINT ["./api-server"]