package routes

import (
	"net/http"
	"strings"

	"SecureFileshare/service/backend/auth"
	"SecureFileshare/service/backend/controllers"
)

func Mux(controllers *controllers.Controllers, req *http.Request, uri string) http.HandlerFunc {
	switch req.Method {
	case http.MethodPost:
		if uri == "/login" {
			return controllers.UserController.Login
		} else if uri == "/verify" {
			return controllers.UserController.Verify
		} else if uri == "/upload" {
			return validateJWTAndHandle(controllers.FileController.Upload)
		}
		return http.NotFound
	case http.MethodGet:
		if uri == "/showfiles" {
			return controllers.FileController.ShowFiles
		} else if strings.HasPrefix(uri, "/download") {
			return controllers.FileController.Download
		}
		return http.NotFound
	default:
		return http.NotFound
	}
}

func RegisterRoutes(controllers *controllers.Controllers) {
	http.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		handler := Mux(controllers, req, req.URL.Path)
		handler(w, req)
	})

	http.HandleFunc("/verify", func(w http.ResponseWriter, req *http.Request) {
		handler := Mux(controllers, req, req.URL.Path)
		handler(w, req)
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, req *http.Request) {
		handler := Mux(controllers, req, req.URL.Path)
		handler(w, req)
	})

	http.HandleFunc("/showfiles", func(w http.ResponseWriter, req *http.Request) {
		handler := Mux(controllers, req, req.URL.Path)
		handler(w, req)
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, req *http.Request) {
		handler := Mux(controllers, req, req.URL.Path)
		handler(w, req)
	})
}

func validateJWTAndHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := auth.ValidateJWT(w, r)
		if err != nil {
			return
		}
		next(w, r)
	}
}
