package main

import (
    "net/http"
	"net/http/httptest"
    "testing"
    
    "github.com/gorilla/context"
)

func TestFilesHandler(t *testing.T) {
    mockStore := &MockStorer{}
	// Create service with test configs
	s := &Service{
		StorageDirectory: "./test_files/",
        TokenSigningKey: []byte("test_config"),
        Store: mockStore,
		Users: map[string]*User{
            "testusername": {
                Username: "testusername", 
                Password: "test_password",
                Folder: "",
            },
        },
    }
    testFolderPath := s.StorageDirectory + s.Users["testusername"].Folder
    mockStore.On("listFiles", testFolderPath).Return(nil, nil)
    // Create a request
    req, err := http.NewRequest("POST", "/files", nil)
    if err != nil {
        t.Fatal(err)
    }
    // Set context
    testclaims := map[string]interface{}{"username": "testusername"}
    context.Set(req, "decoded", testclaims)
    // Create ResponseRecorder
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(s.filesListHandler)
    // Call handler
    handler.ServeHTTP(rr, req)
	// Check the status code
    if status := rr.Code; status != 200 {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, 200)
    }
}