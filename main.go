package main

import (
	"log"
	"net/http"
	"rhythm/server"
)

func main() {
    r := server.NewRouter()

    log.Println("Starting Rhythm server on :8080...")

    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}