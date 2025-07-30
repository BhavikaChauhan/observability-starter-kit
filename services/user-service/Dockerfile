# Dockerfile (multi‑stage, static binary)
# Stage 1: build
FROM golang:1.21-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Disable cgo so we get a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -o user-service main.go

# Stage 2: minimal runtime
FROM scratch
# Copy CA certs if you need HTTPS/TLS, otherwise you can omit
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/user-service /user-service
ENTRYPOINT ["/user-service"]











