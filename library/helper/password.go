package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// ComparePasswords compares a hashed password with a plain text password.
// Returns true if they match, false otherwise.
func ComparePasswords(hashed string, plain string) bool {
	byteHash := []byte(hashed)
	bytePlain := []byte(plain)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// HashPassword generates a bcrypt hash from the given password.
// Returns the hashed password byte slice and any error encountered.
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
