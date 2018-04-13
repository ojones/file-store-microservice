package main

// Service for handling requests
type Service struct {
	MaxUploadSize int64
	StorageDirectory string
	Store Storer
	TokenSigningKey []byte
	UploadFormField string
	Users map[string]*User
}