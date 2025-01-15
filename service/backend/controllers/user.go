package controllers

import (
	"fmt"
	"net/http"
)

type UserController struct{}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("Login attempt received")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
