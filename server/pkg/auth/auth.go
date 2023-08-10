package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash's password
func HashPassword(password string) (string, error) {
	dat, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

// ComparePasswordHash compares hashedpassword and password. returns nil of success else error
func ComparePasswordHash(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}

	return nil
}
