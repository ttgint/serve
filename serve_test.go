package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaults(t *testing.T) {
	handler := http.FileServer(FileSystem{
		fs:   http.Dir("."),
		root: "index.html",
	})

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() != "<p>hello world</p>" {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), "<p>hello world</p>")
	}
}

func TestRewrite(t *testing.T) {
	handler := http.FileServer(FileSystem{
		fs:   http.Dir("."),
		root: "index.html",
	})

	req, err := http.NewRequest("GET", "/virtual/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() != "<p>hello world</p>" {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), "<p>hello world</p>")
	}
}
