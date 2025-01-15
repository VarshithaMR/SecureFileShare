package routes

import (
	"net/http"

	"SecureFileshare/service/backend/controllers"
)

func Mux(controllers *controllers.Controllers, req *http.Request, theURI string) http.HandlerFunc {
	switch req.Method {
	case http.MethodPost:
		if theURI == "/login" {
			return controllers.UserController.Login
		} else if theURI == "/upload" {
			return controllers.FileController.Upload
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

	http.HandleFunc("/upload", func(w http.ResponseWriter, req *http.Request) {
		handler := Mux(controllers, req, req.URL.Path)
		handler(w, req)
	})
}
