package infrastructures

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/golang-jwt/jwt/v5"

	"github.com/abrishk26/a2sv-project-track/task8/Domain"
)

func TestJWTTokenService_GenerateAndVerify(t *testing.T) {
	secretKey := []byte("mysecretkey1234567890")

	service := NewTokenService(secretKey)

	userID := "user123"

	// Generate token
	tokenString, err := service.GenerateToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
	assert.True(t, strings.Count(tokenString, ".") == 2, "token should have 3 parts separated by dots")

	// Verify token (valid)
	subject, err := service.VerifyToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, userID, subject)
}

func TestJWTTokenService_VerifyToken_InvalidToken(t *testing.T) {
	service := NewTokenService([]byte("somekey"))

	// Pass invalid token string
	_, err := service.VerifyToken("invalid.token.string")
	assert.ErrorIs(t, err, domain.ErrInvalidToken)
}

func TestJWTTokenService_VerifyToken_ExpiredToken(t *testing.T) {
	// To test expiration, generate a token with expiration in the past
	secretKey := []byte("testkey123")

	service := &JWTTokenService{key: secretKey}

	// Create expired claims
	expiredClaims := jwt.RegisteredClaims{
		Subject:   "user123",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)), // expired 1 hour ago
		IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
	tokenString, err := token.SignedString(secretKey)
	assert.NoError(t, err)

	subject, err := service.VerifyToken(tokenString)
	assert.ErrorIs(t, err, domain.ErrExpiredToken)
	assert.Empty(t, subject)
}
