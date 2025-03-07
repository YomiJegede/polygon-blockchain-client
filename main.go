// Entry Point for the Application
package main

import (
	"fmt"
	"log"
	"net/http"
	"blockchain-client/handlers"
)

func main() {
	http.HandleFunc("/blockNumber", handlers.BlockNumberHandler)
	http.HandleFunc("/blockByNumber", handlers.BlockByNumberHandler)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
