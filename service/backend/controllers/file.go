package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"SecureFileshare/service/backend/auth"
)

var secretKey = []byte("mysecretkey12345") // 16 bytes for AES-128

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
	filePath := "./uploadedfiles/files/" + fileName + fileExt
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

	// Encrypt the file after it is uploaded
	encryptedFilePath := "./uploadedfiles/files/encrypted_" + fileName + fileExt

	err = auth.EncryptFile(filePath, encryptedFilePath, secretKey)
	if err != nil {
		http.Error(w, "Error encrypting file", http.StatusInternalServerError)
		return
	}

	// Remove the unencrypted file after encryption
	err = os.Remove(filePath)
	if err != nil {
		http.Error(w, "Error removing unencrypted file", http.StatusInternalServerError)
		return
	}

	username := r.Header.Get("UploadedBy")
	record := FileUploadRecord{
		Username: username,
		Filename: "encrypted_" + fileName + fileExt,
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

func (fc *FileController) Download(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	filePath := fmt.Sprintf("./uploadedfiles/files/%s", filename)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Decrypt the file before sending it to the user
	decryptedFilePath := "./uploadedfiles/files/decrypted_" + filename
	defer os.Remove(decryptedFilePath)

	err := auth.DecryptFile(filePath, decryptedFilePath, secretKey)
	if err != nil {
		http.Error(w, "Error decrypting file", http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, decryptedFilePath)
}

func (fc *FileController) DeleteFiles(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}

	// Define the file path (assuming files are in the 'uploadedfiles' folder)
	filePath := filepath.Join("./uploadedfiles/files", filename)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Delete the file from the filesystem
	err := os.Remove(filePath)
	if err != nil {
		http.Error(w, "Failed to delete file", http.StatusInternalServerError)
		return
	}

	// Optionally: Remove the file record from your uploads JSON (if necessary)
	err = deleteFileRecord(filename)
	if err != nil {
		http.Error(w, "Failed to update file records", http.StatusInternalServerError)
		return
	}

	// Return a success message
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "File deleted successfully",
	}
	_ = json.NewEncoder(w).Encode(response)
}

func deleteFileRecord(filename string) error {
	// Load existing file records
	var records []FileUploadRecord
	recordsFile := "./uploadedfiles/uploads.json"

	file, err := os.Open(recordsFile)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&records)
	if err != nil {
		return err
	}

	// Find and remove the record for the deleted file
	var updatedRecords []FileUploadRecord
	for _, record := range records {
		if record.Filename != filename {
			updatedRecords = append(updatedRecords, record)
		}
	}

	// Save the updated records back to the file
	newFile, err := os.Create(recordsFile)
	if err != nil {
		return err
	}
	defer newFile.Close()

	encoder := json.NewEncoder(newFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(updatedRecords)
	if err != nil {
		return err
	}

	return nil
}
