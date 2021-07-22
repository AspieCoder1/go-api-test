package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWorks(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(get)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Invalid status code: got %v want %v", rr.Code, http.StatusOK)
	}

	expected := `{"message": "get method okay"}`

	if rr.Body.String() != expected {
		t.Errorf("Invalid body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPostWorks(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(post)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Invalid status code: got %v want %v", rr.Code, http.StatusOK)
	}

	expected := `{"message": "post method okay"}`

	if rr.Body.String() != expected {
		t.Errorf("Invalid body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteWorks(t *testing.T) {
	req, err := http.NewRequest("PATCH", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(delete)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Invalid status code: got %v want %v", rr.Code, http.StatusOK)
	}

	expected := `{"message": "delete method okay"}`

	if rr.Body.String() != expected {
		t.Errorf("Invalid body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPatchWorks(t *testing.T) {
	req, err := http.NewRequest("PATCH", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(patch)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Invalid status code: got %v want %v", rr.Code, http.StatusOK)
	}

	expected := `{"message": "patch method okay"}`

	if rr.Body.String() != expected {
		t.Errorf("Invalid body: got %v want %v", rr.Body.String(), expected)
	}
}
