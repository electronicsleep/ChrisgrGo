package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: (%v : %v)", status, http.StatusOK)
	} else {
		fmt.Printf("handler returned correct status code: (%v : %v)\n", status, http.StatusOK)
	}

	expected := "chrisgr: health ok"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: (%v, %v)", rr.Body.String(), expected)
	} else {
		fmt.Printf("handler returned expected body: (%v, %v)\n", rr.Body.String(), expected)
	}
}

func TestTimeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/time", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(timeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: ( %v : %v )", status, http.StatusOK)
	} else {
		fmt.Printf("handler returned correct status code: ( %v : %v )\n", status, http.StatusOK)
	}

	expected := "time:"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: ( %v, %v )", rr.Body.String(), expected)
	} else {
		fmt.Printf("handler contained expected string: ( %v, %v )\n", rr.Body.String(), expected)
	}
}
