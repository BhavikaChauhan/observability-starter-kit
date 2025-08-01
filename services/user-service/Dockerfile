# Build Stage
FROM golang:1.21 as builder

WORKDIR /app

# Copy source code
COPY . .

# Disable CGO for static build (avoids glibc dependency)
ENV CGO_ENABLED=0

# Build the Go binary
RUN go build -o user-service .

# Final lightweight image
FROM gcr.io/distroless/static

# Copy only the binary into the distroless image
COPY --from=builder /app/user-service /

# Run the binary
CMD ["/user-service"]

