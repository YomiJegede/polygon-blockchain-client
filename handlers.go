//This file provides HTTP handlers for fetching blockchain data.
package handlers

import (
	"net/http"
	"polygon-blockchain-client/client"
)

// BlockNumberHandler handles requests to fetch the latest block number.
func BlockNumberHandler(w http.ResponseWriter, r *http.Request) {
	blockNumber, err := client.FetchBlockNumber()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(blockNumber))
}

// BlockByNumberHandler handles requests to fetch a block by number.
func BlockByNumberHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	blockNumber := query.Get("block")
	if blockNumber == "" {
		http.Error(w, "Missing block number", http.StatusBadRequest)
		return
	}

	block, err := client.FetchBlockByNumber(blockNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(block))
}
