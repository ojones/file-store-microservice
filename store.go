package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// Storer decouple and to test without writing to disk
type Storer interface {
	listFiles(folderpath string) ([]byte, error)
    putFile(filename string, folderpath string, file io.Reader) error
    getFile(filepath string) ([]byte, int64, error)
    deleteFile(filepath string) error
}

type store struct {}

func (s *store) listFiles(folderPath string) ([]byte, error) {
	// Check if folder path actually exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return nil, err
	}
	
	// Read files
	files, err := ioutil.ReadDir(folderPath)
    if err != nil {
		return nil, err
	}
	
	// Create output
	filenames := []string{}
    for _, f := range files {
		filenames = append(filenames, f.Name())
	}
	output, err := json.Marshal(filenames)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *store) putFile(filename string, folderpath string, file io.Reader) error {
	// Read file
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	
	// Create folder if not there
	if _, err := os.Stat(folderpath); os.IsNotExist(err) {
		os.Mkdir(folderpath, os.ModePerm)
	}
	
	// Write file
	newPath := folderpath + filename
	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer newFile.Close()
	if _, err := newFile.Write(fileBytes); err != nil {
		return err
	}
	return nil
}

func (s *store) getFile(filepath string) ([]byte, int64, error) {
	// Open
	file, err := os.Open(filepath)
	if err != nil {
	  	return nil, 0, err
	}
	defer file.Close()
	
	// Get info for file size
	fileinfo, err := file.Stat()
	if err != nil {
	  	return nil, 0, err
	}
	filesize := fileinfo.Size()
	
	// Create output
	buffer := make([]byte, filesize)
	_, err = file.Read(buffer)
	if err != nil {
	  	return nil, 0, err
	}
	return buffer, filesize, nil
}

func (s *store) deleteFile(filepath string) error {
	// Delete file
	err := os.Remove(filepath)
	if err != nil {
		return err
	}
	return nil
}
