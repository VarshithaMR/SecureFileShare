package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"time"
)

// GenerateMFACode generates a time-based OTP (MFA code)
func GenerateMFACode(secret string) string {
	timeWindow := time.Now().Unix() / 30

	// Decode the secret key (base32 encoded)
	hmacKey, _ := base32.StdEncoding.DecodeString(secret)

	// Create the HMAC using the key and the time window
	hash := hmac.New(sha1.New, hmacKey)
	hash.Write([]byte(fmt.Sprintf("%d", timeWindow)))
	otp := hash.Sum(nil)

	return fmt.Sprintf("%06d", otp[0]%10) // Return 6-digit OTP
}

func ValidateMFA(code, secret string) bool {
	expectedCode := GenerateMFACode(secret)
	if code == expectedCode {
		return true
	}
	return false
}
