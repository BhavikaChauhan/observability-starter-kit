# 👷‍♂️ Builder stage: Build a statically linked Go binary
FROM golang:1.22 AS builder

WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build a statically-linked binary (no glibc needed)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o user-service

# 🏃‍♀️ Final image: Alpine for small size
FROM alpine:latest

WORKDIR /

# Copy the compiled binary from builder
COPY --from=builder /app/user-service .

# Make it executable
RUN chmod +x ./user-service

# Expose service port (adjust if needed)
EXPOSE 8080

# Start the service
CMD ["./user-service"]








