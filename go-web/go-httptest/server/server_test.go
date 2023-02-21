package main

import (
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"
)

func TestUpperCaseHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/upper?word=abc", nil)
	w := httptest.NewRecorder()

	// pass the reqeust and response to handler
	upperCaseHandler(w, req)
	res := w.Result()
	defer func(){
		err := res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// read body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	// test
	if string(data) != "ABC" {
		t.Errorf("expected ABC got %v", string(data))
	}
}