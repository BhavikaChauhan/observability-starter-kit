# Stage 1: Build the Go binary
FROM golang:1.21 as builder
WORKDIR /app

# Copy go files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-service main.go

# Stage 2: Create a minimal image
FROM alpine:latest
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/user-service .

# Expose port and run
EXPOSE 8080
ENTRYPOINT ["./user-service"]
