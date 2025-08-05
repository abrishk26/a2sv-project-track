package infrastructures

import (
	"errors"

	"github.com/abrishk26/a2sv-project-track/task8/Domain"
	"golang.org/x/crypto/bcrypt"
)

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

type PasswordService struct{}

func (ps *PasswordService) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.ErrPasswordHashingFailed
	}

	return string(hash), nil
}

func (ps *PasswordService) Verify(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return domain.ErrPasswordVerificationFailed
		}
		return err
	}

	return nil
}
