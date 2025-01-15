package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type FileController struct{}

type FileUploadRecord struct {
	Username string `json:"username"`
	Filename string `json:"filename"`
}

func (fc *FileController) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileName := uuid.New().String()
	fileExt := filepath.Ext(fileHeader.Filename)
	filePath := "./uploadedfiles/" + fileName + fileExt
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	username := r.Header.Get("UploadedBy")
	record := FileUploadRecord{
		Username: username,
		Filename: fileName + fileExt,
	}

	var records []FileUploadRecord
	recordsFile := "./uploadedfiles/uploads.json"
	if _, err := os.Stat(recordsFile); err == nil {
		file, err := os.Open(recordsFile)
		if err != nil {
			http.Error(w, "Unable to read the JSON file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Decode existing records
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&records); err != nil {
			http.Error(w, "Unable to decode the JSON file", http.StatusInternalServerError)
			return
		}
	}

	// Add the new record to the existing records
	records = append(records, record)

	newFile, err := os.Create(recordsFile)
	if err != nil {
		http.Error(w, "Unable to write to the JSON file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	encoder := json.NewEncoder(newFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(records); err != nil {
		http.Error(w, "Unable to save the JSON file", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Println("File uploaded successfully!")
	if err != nil {
		return
	}
}

func (fc *FileController) ShowFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	recordsFile := "./uploadedfiles/uploads.json"

	file, err := os.Open(recordsFile)
	if err != nil {
		http.Error(w, "Unable to read the JSON file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var records []FileUploadRecord
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&records); err != nil {
		http.Error(w, "Unable to decode the JSON file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(records); err != nil {
		http.Error(w, "Unable to encode the records", http.StatusInternalServerError)
		return
	}
}
