package utils

import "golang.org/x/crypto/bcrypt"

func PasswordControl(hash, pass string) bool {
	passwordControl := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return passwordControl == nil
}

func HashPassword(password string) (string, error) {
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hasPassword), err
}
