# STEP 1: Build a static binary
FROM golang:1.21 as builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0
RUN go build -tags netgo -ldflags '-w -extldflags "-static"' -o user-service .

# STEP 2: Minimal image
FROM gcr.io/distroless/static
COPY --from=builder /app/user-service /
CMD ["/user-service"]
