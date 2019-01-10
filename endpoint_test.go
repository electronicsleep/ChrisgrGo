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

func TestEndpointsCheckHandler(t *testing.T) {
	endpoint_list := []string{"/about", "/", "/linux", "/apple", "/projects"}
	for _, endpoint := range endpoint_list {
		fmt.Println("endpoint: ", endpoint)
		req, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(templatePageHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: (%v : %v)", status, http.StatusOK)
		} else {
			fmt.Printf("handler returned correct status code: (%v : %v)\n", status, http.StatusOK)
		}

		expected := "linux"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("handler did not find expected string body: (expected: %v, endpoint: %v)", expected, endpoint)
		} else {
			fmt.Printf("handler contained expected string: (expected: %v, endpoint: %v)\n", expected, endpoint)
		}
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
