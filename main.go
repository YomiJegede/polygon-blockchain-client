package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"bytes"
	"io/ioutil"
)

const polygonRPCURL = "https://polygon-rpc.com/"

// JSON-RPC Request Struct
type JSONRPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
	ID      int           `json:"id"`
}

// JSON-RPC Response Struct
type JSONRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
}

// FetchBlockNumber gets the latest block number
func FetchBlockNumber() (string, error) {
	reqBody := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		ID:      2,
	}
	return sendRPCRequest(reqBody)
}

// FetchBlockByNumber gets a block by number
func FetchBlockByNumber(blockNumber string) (string, error) {
	reqBody := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{blockNumber, true},
		ID:      2,
	}
	return sendRPCRequest(reqBody)
}

// sendRPCRequest sends an RPC request to Polygon
func sendRPCRequest(reqBody JSONRPCRequest) (string, error) {
	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(polygonRPCURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

// API Handlers
func blockNumberHandler(w http.ResponseWriter, r *http.Request) {
	blockNumber, err := FetchBlockNumber()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(blockNumber))
}

func blockByNumberHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	blockNumber := query.Get("block")
	if blockNumber == "" {
		http.Error(w, "Missing block number", http.StatusBadRequest)
		return
	}

	block, err := FetchBlockByNumber(blockNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(block))
}

func main() {
	http.HandleFunc("/blockNumber", blockNumberHandler)
	http.HandleFunc("/blockByNumber", blockByNumberHandler)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
