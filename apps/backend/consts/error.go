package consts

import "errors"

var (
	ErrUserBanned                = errors.New("User is banned")
	ErrInvalidCredentials        = errors.New("Invalid credentials")
	ErrUserNotFound              = errors.New("User not found")
	ErrInvalidRefreshToken       = errors.New("Invalid refresh token")
	ErrVerificationTokenExpired  = errors.New("Verification token expired")
	ErrRefreshTokenExpired       = errors.New("Refresh token expired")
	ErrEmailAlreadyVerified      = errors.New("Email already verified")
	ErrEmailVerificationDisabled = errors.New("Email verification is disabled")
)
