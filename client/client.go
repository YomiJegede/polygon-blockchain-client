package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"polygon-blockchain-client/config"
)

// JSONRPCRequest for standard JSON-RPC request.
type JSONRPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
	ID      int           `json:"id"`
}

// JSONRPCResponse for standard JSON-RPC response.
type JSONRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
}

// sendRPCRequest sends a JSON-RPC request to the Polygon RPC server.
func sendRPCRequest(reqBody JSONRPCRequest) (string, error) {
	url := config.GetPolygonRPCURL()

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch data from Polygon RPC")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

// FetchBlockNumber retrieves the latest block number.
func FetchBlockNumber() (string, error) {
	req := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		ID:      2,
	}
	return sendRPCRequest(req)
}

// FetchBlockByNumber retrieves a block by its number.
func FetchBlockByNumber(blockNumber string) (string, error) {
	req := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{blockNumber, true},
		ID:      2,
	}
	return sendRPCRequest(req)
}
