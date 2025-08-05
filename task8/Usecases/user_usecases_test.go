package usecases_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/abrishk26/a2sv-project-track/task8/Domain"
	"github.com/abrishk26/a2sv-project-track/task8/Usecases"
	"github.com/abrishk26/a2sv-project-track/task8/mocks"
)

func TestUserUsecases_Login_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockPass := new(mocks.MockPasswordService)
	mockToken := new(mocks.MockTokenService)

	us := usecases.NewUserUsecases(mockRepo, mockPass, mockToken)
	ctx := context.Background()

	user := &domain.User{ID: "u1", Username: "john", PasswordHash: "hash"}

	mockRepo.On("GetByUsername", ctx, "john").Return(user, nil)
	mockPass.On("Verify", "pass", "hash").Return(nil)
	mockToken.On("GenerateToken", "u1").Return("token123", nil)

	token, err := us.Login(ctx, "john", "pass")

	assert.NoError(t, err)
	assert.Equal(t, "token123", token)

	mockRepo.AssertExpectations(t)
	mockPass.AssertExpectations(t)
	mockToken.AssertExpectations(t)
}

func TestUserUsecases_Login_InvalidUsername(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockPass := new(mocks.MockPasswordService)
	mockToken := new(mocks.MockTokenService)

	us := usecases.NewUserUsecases(mockRepo, mockPass, mockToken)
	ctx := context.Background()

	mockRepo.On("GetByUsername", ctx, "unknown").Return(nil, domain.ErrUserNotFound)

	token, err := us.Login(ctx, "unknown", "pass")

	assert.ErrorIs(t, err, domain.ErrInvalidCredential)
	assert.Empty(t, token)
}

func TestUserUsecases_Login_WrongPassword(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockPass := new(mocks.MockPasswordService)
	mockToken := new(mocks.MockTokenService)

	us := usecases.NewUserUsecases(mockRepo, mockPass, mockToken)
	ctx := context.Background()

	user := &domain.User{ID: "u1", Username: "john", PasswordHash: "hash"}

	mockRepo.On("GetByUsername", ctx, "john").Return(user, nil)
	mockPass.On("Verify", "wrongpass", "hash").Return(domain.ErrPasswordVerificationFailed)

	token, err := us.Login(ctx, "john", "wrongpass")

	assert.ErrorIs(t, err, domain.ErrInvalidCredential)
	assert.Empty(t, token)
}

func TestUserUsecases_Register_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockPass := new(mocks.MockPasswordService)
	mockToken := new(mocks.MockTokenService)

	us := usecases.NewUserUsecases(mockRepo, mockPass, mockToken)
	ctx := context.Background()

	user := domain.User{Username: "john", PasswordHash: "pass"}

	mockPass.On("Hash", "pass").Return("hashedpass", nil)
	mockRepo.On("Add", ctx, mock.MatchedBy(func(u domain.User) bool {
		return u.PasswordHash == "hashedpass" && u.Username == "john"
	})).Return(nil)

	err := us.Register(ctx, user)

	assert.NoError(t, err)

	mockPass.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
