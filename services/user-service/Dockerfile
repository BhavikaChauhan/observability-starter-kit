# Start from the official Golang image as a builder
FROM golang:1.22 as builder

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go binary statically
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-service .

# Final lightweight image
FROM alpine:latest

WORKDIR /

# Copy the statically linked binary
COPY --from=builder /app/user-service .

# Expose service port
EXPOSE 8080

# Run the binary
CMD ["./user-service"]






