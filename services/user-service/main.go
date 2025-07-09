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
    log.Fatal(http.ListenAndServe(":8000", nil))
}
