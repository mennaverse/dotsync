package auth

import (
	"errors"
)

var (
	ErrUserBanned         = errors.New("User is banned")
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrUserNotFound       = errors.New("User not found")
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
