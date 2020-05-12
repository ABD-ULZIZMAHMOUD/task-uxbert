package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

/**
* hash passwords
 */
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	return string(bytes), err
}

/**
* check if password is valid
 */
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
