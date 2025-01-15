package auth

import (
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
