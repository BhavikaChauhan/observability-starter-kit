package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "User Service is healthy!")
    })

    log.Println("Starting User Service on port 8000...")
    http.ListenAndServe("0.0.0.0:8000", nil) // This is the important change
}
