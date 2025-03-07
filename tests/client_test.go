package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"polygon-blockchain-client/handlers"
	"io/ioutil"
)

// Test fetching the latest block number
func TestBlockNumberHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/blockNumber", nil)
	w := httptest.NewRecorder()
	handlers.BlockNumberHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(body) == 0 {
		t.Errorf("Expected non-empty response")
	}
}

// Test fetching a block by number
func TestBlockByNumberHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/blockByNumber?block=0x134e82a", nil)
	w := httptest.NewRecorder()
	handlers.BlockByNumberHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(body) == 0 {
		t.Errorf("Expected non-empty response")
	}
}
