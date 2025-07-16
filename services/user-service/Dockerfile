# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o user-service

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/user-service .
EXPOSE 8080
CMD ["./user-service"]