package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	// Create service with test configs
	s := &Service{
		StorageDirectory: "test_config",
		TokenSigningKey:  []byte("test_config"),
		Users:            map[string]*User{},
	}

	// Create a request
	requestJSON := `{"username": "testusername", "password": "test_password"}`
	reader := strings.NewReader(requestJSON)
	req, err := http.NewRequest("POST", "/register", reader)
	if err != nil {
		t.Fatal(err)
	}

	// Create ResponseRecorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.registerHandler)

	// Call handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != 204 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 204)
	}

	// Check the response body
	expected := ``
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
