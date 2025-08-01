FROM golang:1.21 as builder

WORKDIR /app

COPY . .

# Force static linking (very important!)
ENV CGO_ENABLED=0
RUN go build -ldflags="-linkmode external -extldflags '-static'" -o user-service .

FROM gcr.io/distroless/static
COPY --from=builder /app/user-service /
CMD ["/user-service"]

