package infrastructures

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(claims map[string]any, secret string) (string, error) {
	c := jwt.MapClaims{}
	for claim, value := range claims {
		c[claim] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(token, secret string, claims jwt.Claims) (*jwt.Token, error) {
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}
