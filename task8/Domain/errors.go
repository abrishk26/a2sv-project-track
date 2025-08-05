package domain

import "errors"

var (
	ErrInvalidCredential = errors.New("invalid email or password")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrInvalidToken      = errors.New("invalid token")
	ErrExpiredToken      = errors.New("token has expired")

	ErrDuplicateUsername          = errors.New("email already registered")
	ErrDuplicateTask              = errors.New("task already exists")
	ErrUserNotFound               = errors.New("user not found")
	ErrPasswordHashingFailed      = errors.New("failed to hash password")
	ErrPasswordVerificationFailed = errors.New("password does not match")

	ErrTaskNotFound = errors.New("task not found")

	ErrTokenParsingFailed = errors.New("failed to parse token")
	ErrTokenSigningFailed = errors.New("failed to sign token")
)
