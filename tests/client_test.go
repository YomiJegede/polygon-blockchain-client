package tests

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"polygon-blockchain-client"
)

func TestBlockNumberHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/blockNumber", nil)
	w := httptest.NewRecorder()
	polygon-blockchain-client.blockNumberHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(body) == 0 {
		t.Errorf("Expected non-empty response")
	}
}

func TestBlockByNumberHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/blockByNumber?block=0x134e82a", nil)
	w := httptest.NewRecorder()
	polygon-blockchain-client.blockByNumberHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(body) == 0 {
		t.Errorf("Expected non-empty response")
	}
}
