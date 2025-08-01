# Stage 1: Build the Go binary
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Optional: tidy + verify if needed
# RUN go mod tidy && go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-service .

# Stage 2: Minimal base image with GLIBC 2.34+
FROM debian:bullseye-slim

# Ensure updated glibc (2.31+) is present
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates libc6 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /

COPY --from=builder /app/user-service /user-service

EXPOSE 8080

CMD ["/user-service"]






