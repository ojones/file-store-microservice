package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	// Get configs
	c := getConfigs()
	// Create service with configs
	s := &Service{
		MaxUploadSize: c.MaxUploadSize,
		StorageDirectory: c.StorageDirectory,
		Store: &store{},
		TokenSigningKey: []byte(c.TokenSigningSecret),
		UploadFormField: c.UploadFormField,
		Users: map[string]*User{},
	}
	// Create routes
    r := mux.NewRouter()
	r.HandleFunc("/register", s.registerHandler).Methods("POST")
	r.HandleFunc("/login", s.loginHandler).Methods("POST")
	r.HandleFunc("/files", s.filesListHandler).Methods("GET")
	r.HandleFunc("/files/{filename}", s.ValidateMiddleware(s.filesGetHandler)).Methods("GET")
	r.HandleFunc("/files/{filename}", s.ValidateMiddleware(s.filesPutHandler)).Methods("PUT")
	r.HandleFunc("/files/{filename}", s.ValidateMiddleware(s.filesDeleteHandler)).Methods("DELETE")
	// Create server
    srv := &http.Server{
        Handler:      handlers.LoggingHandler(os.Stdout, r),
        Addr:         fmt.Sprintf("%s:%s", c.IP, c.Port),
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
	// Run server
    log.Fatal(srv.ListenAndServe())
}