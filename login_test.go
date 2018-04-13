package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestLoginHandler(t *testing.T) {
	// Create service with test configs
	s := &Service{
		StorageDirectory: "test_config",
		TokenSigningKey:  []byte("test_config"),
		Users:            map[string]*User{},
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte("test_password"), bcrypt.DefaultCost)
	s.Users["testusername"] = &User{
		Username: "testusername",
		Password: string(hash),
		Folder:   "test_files/",
	}

	// Create a request
	requestJSON := `{"username": "testusername", "password": "test_password"}`
	reader := strings.NewReader(requestJSON)
	req, err := http.NewRequest("POST", "/login", reader)
	if err != nil {
		t.Fatal(err)
	}

	// Create ResponseRecorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.loginHandler)

	// Call handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 200)
	}

	// Unmarshal
	var responseMsg LoginResponse
	if err := json.Unmarshal([]byte(rr.Body.String()), &responseMsg); err != nil {
		t.Fatal(err)
	}

	// Check the response body
	if responseMsg.Token == "" {
		t.Errorf("handler returned unexpected Token: got %v want %v",
			responseMsg.Token, "valid jwt string")
	}
}
