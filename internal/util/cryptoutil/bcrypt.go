package cryptoutil

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// CreateBcryptHash creates a bcrypt hash of the given password
func CreateBcryptHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// VerifyBcryptHash verifies the given password against the bcrypt hash
func VerifyBcryptHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password")
	}
	return nil
}
