FROM golang:1.21 as builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o user-service .

FROM gcr.io/distroless/static
COPY --from=builder /app/user-service /
CMD ["/user-service"]








