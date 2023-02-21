package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	tt := []struct{
		name string
		value string
		double int
		status int
		err string
	}{
		{name: "double of 2", value: "2", double: 4, status: 200},
		{name: "missing value", status: 400, err: "missing value"},
		{name: "not a number", value: "a", status: 400, err: "not a number"},
	}

	for _, tc := range tt {  // loop over each test case tc
		t.Run(tc.name, func(t *testing.T) {  // run test case in test table tt
			req, err := http.NewRequest("GET", "localhost:8080/double?v=" + tc.value, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()
			doubleHandler(rec, req)

			res := rec.Result()
			defer func() {
				_ = res.Body.Close()
			}()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}
			msg := string(body)
			msg = strings.TrimSpace(msg)

			// always check the status code
			if res.StatusCode != tc.status {
				t.Errorf("expected %v got %v", tc.status, res.StatusCode)
			}

			// check error message if expected
			if tc.err != "" {
				// do something
				if msg != tc.err {
					t.Errorf("expected '%v' got '%v'", tc.err, msg)
				}
				return
			}

			// check value
			v, err := strconv.Atoi(msg)
			if err != nil {
				t.Fatalf("expected an integer got %v", err)
			}
			if v != tc.double {
				t.Errorf("expected %v got %v", tc.double, v)
			}
		})
	}
}

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	url := fmt.Sprintf("%s/double?v=2", srv.URL)
	log.Println("NewServer created a server at:", url)

	res, err := http.Get(url)
	if err != nil {
		t.Errorf("could not send request: %v", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("expected 200 got %v", res.StatusCode)
	}
	v, err := strconv.Atoi(string(bytes.TrimSpace(body)))
	if err != nil {
		t.Fatalf("expected an integer got %v", err)
	}
	if v != 4 {
		t.Errorf("expected 4 got %d", v)
	}
}