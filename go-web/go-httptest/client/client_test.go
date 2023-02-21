package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func mockUpperCaseHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("dummy data"))
}

func TestClient_UpperCase(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(mockUpperCaseHandler))
	defer mockServer.Close()

	c := NewClient(mockServer.URL)
	res, err := c.UpperCase("anything")
	if err != nil {
		t.Errorf("expect error to be nil got %v", err)
	}

	res = strings.TrimSpace(res) // remove new line characters and spaces
	if res != "dummy data" {
		t.Errorf("expected dummy data got %v", res)
	}
}
