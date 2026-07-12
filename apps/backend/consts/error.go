package consts

import (
	"errors"

	"github.com/mennaverse/dotsync/apps/backend/web/handler/types"
)

var (
	ErrUserBanned                   = errors.New("User is banned")
	ErrInvalidCredentials           = errors.New("Invalid credentials")
	ErrUserNotFound                 = errors.New("User not found")
	ErrInvalidRefreshToken          = errors.New("Invalid refresh token")
	ErrVerificationTokenExpired     = errors.New("Verification token expired")
	ErrTokenInvalid                 = errors.New("Invalid token")
	ErrRefreshTokenExpired          = errors.New("Refresh token expired")
	ErrEmailAlreadyVerified         = errors.New("Email already verified")
	ErrEmailVerificationDisabled    = errors.New("Email verification is disabled")
	ErrResetPasswordTokenExpired    = errors.New("Password reset token expired")
	ErrResetPasswordTokenInvalid    = errors.New("Invalid password reset token")
	ErrHttpInvalidRefreshTokenCode  = errors.New("Invalid refresh token")
	ErrEmailOrUsernameAlreadyExists = errors.New("Email or username already exists")
	ErrHttpMissingTokenCode         = errors.New("Missing token")
	ErrPasswordTooLong              = errors.New("Password exceeds maximum length")
)

func ConvertToAppError(err error) error {
	switch err {
	case ErrUserBanned:
		return types.NewAppError(401, "user_banned", err.Error(), err)
	case ErrInvalidCredentials:
		return types.NewAppError(401, "invalid_credentials", err.Error(), err)
	case ErrUserNotFound:
		return types.NewAppError(404, "user_not_found", err.Error(), err)
	case ErrInvalidRefreshToken:
		return types.NewAppError(401, "invalid_refresh_token", err.Error(), err)
	case ErrVerificationTokenExpired:
		return types.NewAppError(401, "verification_token_expired", err.Error(), err)
	case ErrTokenInvalid:
		return types.NewAppError(401, "invalid_token", err.Error(), err)
	case ErrRefreshTokenExpired:
		return types.NewAppError(401, "refresh_token_expired", err.Error(), err)
	case ErrEmailAlreadyVerified:
		return types.NewAppError(400, "email_already_verified", err.Error(), err)
	case ErrEmailVerificationDisabled:
		return types.NewAppError(400, "email_verification_disabled", err.Error(), err)
	case ErrResetPasswordTokenExpired:
		return types.NewAppError(401, "reset_password_token_expired", err.Error(), err)
	case ErrResetPasswordTokenInvalid:
		return types.NewAppError(401, "reset_password_token_invalid", err.Error(), err)
	case ErrHttpInvalidRefreshTokenCode:
		return types.NewAppError(401, "invalid_refresh_token", err.Error(), err)
	case ErrEmailOrUsernameAlreadyExists:
		return types.NewAppError(400, "email_or_username_already_exists", err.Error(), err)
	case ErrHttpMissingTokenCode:
		return types.NewAppError(401, "missing_token", err.Error(), err)
	case ErrPasswordTooLong:
		return types.NewAppError(400, "password_too_long", err.Error(), err)
	default:
		return types.NewAppError(500, "internal_server_error", err.Error(), err)
	}
}
