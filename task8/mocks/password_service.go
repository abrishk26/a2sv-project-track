package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockPasswordService struct {
	mock.Mock
}

func (m *MockPasswordService) Hash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordService) Verify(password, hash string) error {
	args := m.Called(password, hash)
	return args.Error(0)
}
