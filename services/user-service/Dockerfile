FROM golang:1.22 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o user-service

# Final stage
FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=builder /app/user-service /user-service

EXPOSE 3001
ENTRYPOINT ["/user-service"]
























