package utils

import "golang.org/x/crypto/bcrypt"

func PasswordControl(hash, pass string) bool {
	passwordControl := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if passwordControl != nil {
		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hasPassword), err
}
