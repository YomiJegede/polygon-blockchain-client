// Entry Point for the Application
package main

import (
	"fmt"
	"log"
	"net/http"
	"polygon-blockchain-client/handlers"
)

func main() {
	mux := http.NewServeMux() // Use ServeMux for better routing

	// Register handlers
	mux.HandleFunc("/healthz", handlers.HealthCheckHandler)
	mux.HandleFunc("/blockNumber", handlers.BlockNumberHandler)
	mux.HandleFunc("/blockByNumber", handlers.BlockByNumberHandler)

	port := "8080"
	fmt.Println("✅ Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux)) // Use mux for cleaner routing
}
