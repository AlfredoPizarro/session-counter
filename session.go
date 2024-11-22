package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

var connectionCount uint64

func main() {
	// Get the hostname of the pod
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting hostname: %v", err)
	}

	// Handler for incoming requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Increment the connection count
		newCount := atomic.AddUint64(&connectionCount, 1)

		// Log the current connection count
		log.Printf("Current connections to this pod are: %d", newCount)

		// Write the response
		response := fmt.Sprintf("Connection to %s", hostname)
		_, _ = w.Write([]byte(response))
	})

	// Start the server on port 8080
	port := "8080"
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
