package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/telemetry"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Service is live!")
}

func main() {
	shutdown := telemetry.InitProvider()
	defer shutdown()

	http.HandleFunc("/", handler)
	log.Println("Starting User Service on port 3001...")
	log.Fatal(http.ListenAndServe(":3001", nil))
}

