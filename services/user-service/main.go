package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User service is healthy!"))
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the User service"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	fmt.Printf("User service running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
