package infrastructures

import (
	"time"

	domain "github.com/abrishk26/a2sv-project-track/task8/Domain"
	"github.com/golang-jwt/jwt/v5"
)

func NewTokenService(key []byte) domain.ITokenService {
	return &JWTTokenService{
		key,
	}
}

type JWTTokenService struct {
	key []byte
}

func (ts *JWTTokenService) GenerateToken(userID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * (24 * time.Hour))),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(ts.key)
}

func (ts *JWTTokenService) VerifyToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return ts.key, nil
	})

	if err != nil {
		switch err {
		case jwt.ErrTokenMalformed:
			return "", domain.ErrInvalidToken
		case jwt.ErrTokenExpired:
			return "", domain.ErrExpiredToken
		default:
			return "", err
		}
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims.Subject, nil
	}

	return "", domain.ErrInvalidToken
}
