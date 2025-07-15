package main

import (
    "fmt"
    "log"
    "net/http"
    "go.opentelemetry.io/otel"
)

func main() {
  //cleanup := InitTelemetry()
  //defer cleanup()
  // tracer := otel.Tracer("user-service")

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "User Service is healthy!")
    })

    log.Println("Starting User Service on port 8000...")
    log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
