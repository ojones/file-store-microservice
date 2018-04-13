package main

import (
	"encoding/json"
	"errors"
    "fmt"
	"io/ioutil"
	"net/http"

	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"golang.org/x/crypto/bcrypt"
)

// RegisterRequest for register requests
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
	Password string `json:"password" validate:"required,min=8"`
}

// Validate request returns nil or ValidationErrors ( []FieldError )
func (r *RegisterRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) registerHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	defer r.Body.Close()
	// Unmarshal
	var requestMsg RegisterRequest
	err = json.Unmarshal(b, &requestMsg)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	// Validate
	if err := requestMsg.Validate(); err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	// Check if user already exists
	if _, ok := s.Users[requestMsg.Username]; ok {
		err := errors.New("given username has already registerd")
		fmt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
    // Generate hash to store from user password
    hash, err := bcrypt.GenerateFromPassword([]byte(requestMsg.Password), bcrypt.DefaultCost)
    if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	// Store user with hash
	s.Users[requestMsg.Username] = &User{
		Username: requestMsg.Username, 
		Password: string(hash),
		Folder: uuid.NewV4().String() + "/",
	}
	// Respond
	w.WriteHeader(204)
	// Log
	fmt.Printf("user registered: %s\n", requestMsg.Username)
}