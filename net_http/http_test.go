package net_http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}
	s:= &server{}

	rr:= httptest.NewRecorder()
	handler := http.HandlerFunc(s.ServeHTTP)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returnder wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"message": "hello world"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestCorrectBodyIfWrongPath(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)

	if err != nil {
		t.Fatal(err)
	}
	s:= &server{}

	rr:= httptest.NewRecorder()
	handler := http.HandlerFunc(s.ServeHTTP)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returnder wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"message": "hello world"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}


}
