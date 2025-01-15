package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"SecureFileshare/service/backend/auth"
	"SecureFileshare/service/backend/utils"
)

type UserController struct{}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Login attempt received")

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := utils.ExistingUsers(credentials.Username)
	if user == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if user.Password != credentials.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	//Generate the JWT token
	token, err := auth.GenerateJWT(user.Username, user.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	fmt.Println("Generated JWT : ", token)

	// Respond with the JWT token
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (uc *UserController) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateJWT(w, r)
	if err != nil {
		return
	}
	w.Write([]byte("Welcome to your profile"))
}
