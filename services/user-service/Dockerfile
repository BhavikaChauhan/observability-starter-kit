# Stage 1 - Build
FROM golang:1.22 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o user-service .

# Stage 2 - Run
FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=builder /app/user-service .

EXPOSE 8001
ENTRYPOINT ["/user-service"]























