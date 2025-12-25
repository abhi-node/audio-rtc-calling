package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	pass_hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(pass_hash), err
}

func CompareHash(password string, password_hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
	return err == nil
}
