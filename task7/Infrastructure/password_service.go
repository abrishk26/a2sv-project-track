package infrastructures

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CompareHashAndPassword(password, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
