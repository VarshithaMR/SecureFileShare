package controllers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type FileController struct{}

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

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	dst, err := os.Create(filepath.Join("uploads", "uploaded_file"))
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {

		}
	}(dst)

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Println("File uploaded successfully!")
	if err != nil {
		return
	}
}
