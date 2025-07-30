# Stage 1: Build a static Go binary
FROM golang:1.21-alpine AS builder

# Ensure we have git and libc-dev to compile statically
RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 👇 Very important: disable cgo to make binary fully static
ENV CGO_ENABLED=0

# Build the binary statically
RUN go build -a -installsuffix cgo -o user-service main.go

# Stage 2: Create a minimal final image
FROM scratch

# Optional: add CA certs if HTTPS needed
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/user-service /user-service

ENTRYPOINT ["/user-service"]







