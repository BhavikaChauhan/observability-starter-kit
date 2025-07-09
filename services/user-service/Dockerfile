# Build stage
FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o user-service

# Final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/user-service .
EXPOSE 8000
CMD ["./user-service"]
