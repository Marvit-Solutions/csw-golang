package password

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

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

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
