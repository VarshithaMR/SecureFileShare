package auth

import (
	"fmt"
	"net/http"
	"time"

	twj "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secretkey")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	twj.RegisteredClaims
}

func GenerateJWT(username string, role string) (string, error) {
	claims := &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: twj.RegisteredClaims{
			IssuedAt:  twj.NewNumericDate(time.Now()),
			ExpiresAt: twj.NewNumericDate(time.Now().Add(time.Minute * 15)), //expires for 15min
		},
	}
	token := twj.NewWithClaims(twj.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(w http.ResponseWriter, r *http.Request) (*twj.Token, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return nil, fmt.Errorf("missing token")
	}

	// Remove the "Bearer " prefix if it exists
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := twj.ParseWithClaims(tokenString, &Claims{}, func(token *twj.Token) (interface{}, error) {
		_, ok := token.Method.(*twj.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		fmt.Println("Valid token for user:", claims.Username, "with role:", claims.Role)
		return token, nil
	} else {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return nil, fmt.Errorf("invalid token")
	}
}

func ExtractRoleFromJWT(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", fmt.Errorf("missing token")
	}

	// Parse the token to get the claims
	token, err := twj.Parse(tokenString, func(token *twj.Token) (interface{}, error) {
		if _, ok := token.Method.(*twj.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(twj.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token claims")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", fmt.Errorf("role not found in token claims")
	}

	return role, nil
}
