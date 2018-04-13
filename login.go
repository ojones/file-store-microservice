package main

import (
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	
	"github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

// LoginRequest for login requests
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Validate request returns nil or ValidationErrors ( []FieldError )
func (l *LoginRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(l)
	if err != nil {
		return err
	}
	return nil
}

// LoginResponse for login requests
type LoginResponse struct {
    Token string `json:"token"`
}

func (s *Service) loginHandler(w http.ResponseWriter, r *http.Request) {
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
	var requestMsg LoginRequest
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
	// Check if user exists in storage
	if _, ok := s.Users[requestMsg.Username]; !ok {
		err := errors.New("username is not recognized please register")
		fmt.Println(err)
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
    // Comparing the password with the hash
    if err := bcrypt.CompareHashAndPassword([]byte(s.Users[requestMsg.Username].Password), []byte(requestMsg.Password)); err != nil {
		http.Error(w, err.Error(), 403)
		return
    }
    // Create the token
    token := jwt.New(jwt.SigningMethodHS256)
    // Create a map to store our claims
    claims := token.Claims.(jwt.MapClaims)
	claims["username"] = requestMsg.Username
    claims["expires"] = time.Now().Add(time.Hour * 24).Unix()
    // Sign the token
	tokenString, _ := token.SignedString(s.TokenSigningKey)
	// Create response
	responseMsg := &LoginResponse{Token: tokenString}
	output, err := json.Marshal(responseMsg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Respond
	w.Header().Set("content-type", "application/json")
	w.Write(output)
	// Log
	fmt.Printf("user logged in: %s\n", requestMsg.Username)
}