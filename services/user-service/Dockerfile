# Step 1: Build binary
FROM golang:1.21 as builder

WORKDIR /app
COPY . .

# Build the Go binary (non-static is okay since Debian has glibc)
RUN go build -o user-service .

# Step 2: Use minimal Debian base image
FROM debian:bullseye-slim

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/user-service .

# If you need TLS or DNS resolution, install a few tools
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

EXPOSE 8080

ENTRYPOINT ["./user-service"]
