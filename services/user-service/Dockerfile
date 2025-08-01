# STEP 1: Build the Go binary
FROM golang:1.21 as builder

WORKDIR /app
COPY . .

# Build without static linking, use default dynamic link
RUN go build -o user-service .

# STEP 2: Use newer Debian with GLIBC >= 2.34
FROM debian:bookworm-slim

WORKDIR /app

# Optional: install CA certs for HTTPS clients
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/user-service .

EXPOSE 8080
ENTRYPOINT ["./user-service"]
