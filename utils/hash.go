package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ValidatePw(pwHash, pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwHash), []byte(pw))
	return err == nil
}
