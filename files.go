package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

func (s *Service) filesListHandler(w http.ResponseWriter, r *http.Request) {
	// Get username from context set in middleware from token claims
	claims := context.Get(r, "decoded")
	var user User
	mapstructure.Decode(claims, &user)

	// Check if user exists in storage
	if _, ok := s.Users[user.Username]; !ok {
		err := errors.New("username is not recognized")
		fmt.Println(err)
		http.Error(w, "", 403)
		return
	}

	// Get users personal folder name from data storage
	personalFolder := s.Users[user.Username].Folder
	folderPath := s.StorageDirectory + personalFolder
	output, err := s.Store.listFiles(folderPath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "", 500)
		return
	}

	// Respond
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	// Log
	fmt.Println("files listed")
}

func (s *Service) filesPutHandler(w http.ResponseWriter, r *http.Request) {
	// Get username from context set in middleware from token claims
	claims := context.Get(r, "decoded")
	var user User
	mapstructure.Decode(claims, &user)

	// Check if user exists in storage
	if _, ok := s.Users[user.Username]; !ok {
		err := errors.New("username is not recognized")
		fmt.Println(err)
		http.Error(w, "", 403)
		return
	}

	// Get users personal folder name from data storage
	vars := mux.Vars(r)
	filename := vars["filename"]
	folderpath := s.StorageDirectory + s.Users[user.Username].Folder
	filepath := folderpath + filename

	// Validate file size
	r.Body = http.MaxBytesReader(w, r.Body, s.MaxUploadSize)
	if err := r.ParseMultipartForm(s.MaxUploadSize); err != nil {
		fmt.Println(errors.New("file too big"))
		http.Error(w, "", 400)
		return
	}

	// Get uploaded file
	file, _, err := r.FormFile(s.UploadFormField)
	if err != nil {
		fmt.Println(errors.New("invalid file"))
		http.Error(w, "", 400)
		return
	}
	defer file.Close()

	// Put file in store
	if err := s.Store.addFile(filename, folderpath, file); err != nil {
		fmt.Println(errors.New("cannot write file"))
		http.Error(w, "", 500)
		return
	}

	// Respond
	w.WriteHeader(201)
	w.Header().Set("Location", filepath)

	// Log
	fmt.Println("file added")
}

func (s *Service) filesGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get username from context set in middleware from token claims
	claims := context.Get(r, "decoded")
	var user User
	mapstructure.Decode(claims, &user)

	// Check if user exists in storage
	if _, ok := s.Users[user.Username]; !ok {
		err := errors.New("username is not recognized")
		fmt.Println(err)
		http.Error(w, "", 403)
		return
	}

	// Get users personal folder name from data storage
	personalFolder := s.Users[user.Username].Folder

	// Get file name to create file path
	vars := mux.Vars(r)
	filepath := s.StorageDirectory + personalFolder + vars["filename"]

	// Check if filepath exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println(err.Error())
		http.Error(w, "", 404)
		return
	}

	// Get file and size from store
	buffer, filesize, err := s.Store.getFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "", 500)
		return
	}

	// Check file type
	filetype := http.DetectContentType(buffer)

	// Respond
	w.WriteHeader(200)
	w.Header().Set("Content-Length", strconv.FormatInt(filesize, 10))
	w.Header().Set("Content-Type", filetype)
	w.Write(buffer)

	// Log
	fmt.Println("file returned")
}

func (s *Service) filesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Get username from context set in middleware from token claims
	claims := context.Get(r, "decoded")
	var user User
	mapstructure.Decode(claims, &user)

	// Check if user exists in storage
	if _, ok := s.Users[user.Username]; !ok {
		err := errors.New("username is not recognized")
		fmt.Println(err)
		http.Error(w, "", 403)
		return
	}

	// Get users personal folder name from data storage
	personalFolder := s.Users[user.Username].Folder

	// Get file name to create file path
	vars := mux.Vars(r)
	filepath := s.StorageDirectory + personalFolder + vars["filename"]

	// Check if filepath exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println(err.Error())
		http.Error(w, "", 404)
		return
	}

	// Delete file from store
	err := s.Store.removeFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "", 500)
		return
	}

	// Respond
	w.WriteHeader(204)

	// Log
	fmt.Println("file deleted")
}
