package main

import (
	"errors"
	"fmt"
	"net/http"
    "strings"
    
	"github.com/dgrijalva/jwt-go"
    "github.com/gorilla/context"
)

// JwtToken a session token
type JwtToken struct {
    Token string `json:"token"`
}

// ValidateMiddleware middleware to validate session from authorization header
func (s *Service) ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        authorizationHeader := req.Header.Get("authorization")
        if authorizationHeader != "" {
            bearerToken := strings.Split(authorizationHeader, " ")
            if len(bearerToken) == 2 {
                token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
                    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                        return nil, fmt.Errorf("There was an error with token")
                    }
                    return s.TokenSigningKey, nil
                })
                if err != nil {
					fmt.Println(err)
					http.Error(w, "", 403)
                    return
                }
                if token.Valid {
                    context.Set(req, "decoded", token.Claims)
                    next(w, req)
                } else {
					fmt.Println(errors.New("invalid authorization token"))
					http.Error(w, "", 403)
					return
                }
            }
        } else {
			fmt.Println(errors.New("an authorization header is required"))
			http.Error(w, "", 403)
			return
        }
    })
}