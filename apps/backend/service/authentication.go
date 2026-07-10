package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mennaverse/dotsync/apps/backend/consts"
	"github.com/mennaverse/dotsync/apps/backend/db"
	"github.com/mennaverse/dotsync/apps/backend/dto"
	"github.com/mennaverse/dotsync/apps/backend/manager"
)

type AuthenticationService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*db.User, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.TokenPair, error)
	Logout(ctx context.Context, refreshToken string) error
	LogoutAllSessions(ctx context.Context, userID string) error
	RefreshToken(ctx context.Context, refreshToken string) (*dto.TokenPair, error)

	VerifyEmail(ctx context.Context, token string) error
	ResendVerificationEmail(ctx context.Context, email string) error
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token string, newPassword string) error

	ValidateAccessToken(ctx context.Context, accessToken string) (*dto.Claims, error)
}

type DefaultAuthenticationService struct {
	AuthenticationService
	queries       *db.Queries
	secretManager manager.SecretManager
	cryptoManager manager.CryptoManager
	emailManager  manager.EmailManager
}

func NewAuthenticationService(queries *db.Queries, secretManager manager.SecretManager, cryptoManager manager.CryptoManager, emailManager manager.EmailManager) AuthenticationService {
	return &DefaultAuthenticationService{
		queries:       queries,
		secretManager: secretManager,
		cryptoManager: cryptoManager,
		emailManager:  emailManager,
	}
}

func (s *DefaultAuthenticationService) Register(ctx context.Context, req *dto.RegisterRequest) (*db.User, error) {
	hashedPassword, err := s.cryptoManager.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.queries.InsertUser(ctx, db.InsertUserParams{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *DefaultAuthenticationService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.TokenPair, error) {
	user, err := s.queries.GetUserByUsernameOrEmail(ctx, req.Login)
	if err != nil {
		return nil, consts.ErrUserNotFound
	}

	if user.Banned {
		return nil, consts.ErrUserBanned
	}

	match, err := s.cryptoManager.VerifyPassword(req.Password, user.PasswordHash)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, consts.ErrInvalidCredentials
	}

	// Generate token pair (access and refresh tokens)
	jwtToken, err := s.cryptoManager.GenerateJwtToken(&user)
	if err != nil {
		return nil, err
	}

	refreshToken, _, err := s.cryptoManager.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &dto.TokenPair{
		AccessToken:  jwtToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Hour),
	}, nil
}

func (s *DefaultAuthenticationService) Logout(ctx context.Context, refreshToken string) error {
	tokenHash := s.cryptoManager.HashToken(refreshToken)

	err := s.queries.DeleteUserRefreshTokenByRefreshTokenHash(ctx, tokenHash)
	if err != nil {
		return err
	}

	return nil
}

func (s *DefaultAuthenticationService) LogoutAllSessions(ctx context.Context, userID string) error {
	var pgUid pgtype.UUID
	if err := pgUid.Scan(userID); err != nil {
		return err
	}

	err := s.queries.DeleteUserRefreshTokensByUserID(ctx, pgUid)
	if err != nil {
		return err
	}

	return nil
}

func (s *DefaultAuthenticationService) RefreshToken(ctx context.Context, refreshToken string) (*dto.TokenPair, error) {
	tokenHash := s.cryptoManager.HashToken(refreshToken)

	dbSession, err := s.queries.GetUserRefreshTokenByRefreshTokenHash(ctx, tokenHash)
	if err != nil {
		return nil, consts.ErrInvalidRefreshToken
	}

	if time.Now().After(dbSession.ExpiresAt.Time) {
		_ = s.queries.DeleteUserRefreshTokenByRefreshTokenHash(ctx, tokenHash)
		return nil, consts.ErrRefreshTokenExpired
	}

	user, err := s.queries.GetUserByID(ctx, dbSession.UserID)
	if err != nil {
		return nil, consts.ErrUserNotFound
	}

	accessToken, err := s.cryptoManager.GenerateJwtToken(&user)
	if err != nil {
		return nil, err
	}

	rawRefreshToken, hashedRefreshToken, err := s.cryptoManager.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	// Store the hashed refresh token in the database
	_, err = s.queries.InsertUserRefreshToken(ctx, db.InsertUserRefreshTokenParams{
		UserID:    user.ID,
		TokenHash: hashedRefreshToken,
		ExpiresAt: pgtype.Timestamptz{
			Time:  expiresAt,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &dto.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: rawRefreshToken,
		ExpiresAt:    time.Now().Add(time.Hour),
	}, nil

}

func (s *DefaultAuthenticationService) VerifyEmail(ctx context.Context, rawToken string) error {
	if enabled, err := s.secretManager.GetEmailVerificationEnabled(); err != nil || !enabled {
		return consts.ErrEmailVerificationDisabled
	}

	tokenHash := s.cryptoManager.HashToken(rawToken)

	verificationToken, err := s.queries.GetUserVerificationTokenByHash(ctx, tokenHash)
	if err != nil {
		return err
	}

	if time.Now().After(verificationToken.ExpiresAt.Time) {
		return consts.ErrVerificationTokenExpired
	}

	if err := s.queries.VerifyUserEmail(ctx, verificationToken.UserID); err != nil {
		return err
	}

	if err := s.queries.DeleteUserVerificationTokensByUserID(ctx, verificationToken.UserID); err != nil {
		return err
	}

	return nil
}

func (s *DefaultAuthenticationService) ResendVerificationEmail(ctx context.Context, email string) error {
	if enabled, err := s.secretManager.GetEmailVerificationEnabled(); err != nil || !enabled {
		return consts.ErrEmailVerificationDisabled
	}

	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return consts.ErrUserNotFound
	}

	if user.EmailVerified {
		return consts.ErrEmailAlreadyVerified
	}

	rawToken, hashedToken, err := s.cryptoManager.GenerateVerificationToken()
	if err != nil {
		return err
	}

	err = s.queries.DeleteUserVerificationTokensByUserID(ctx, user.ID)
	if err != nil {
		return err
	}

	expiresAt := time.Now().Add(30 * time.Minute)
	_, err = s.queries.InsertUserVerificationToken(ctx, db.InsertUserVerificationTokenParams{
		UserID:    user.ID,
		TokenHash: hashedToken,
		ExpiresAt: pgtype.Timestamptz{
			Time:  expiresAt,
			Valid: true,
		},
	})
	if err != nil {
		return err
	}

	go func() {
		bgCtx := context.Background()

		body := fmt.Sprintf(consts.EmailVerificationBody, rawToken, rawToken)
		if err := s.emailManager.SendEmailVerification(bgCtx, user.Email, consts.EmailVerificationSubject, body); err != nil {
			// Log the error, but don't return it since this is a background operation
			fmt.Printf("Failed to send verification email: %v\n", err)
		}
	}()

	return nil
}

func (s *DefaultAuthenticationService) ForgotPassword(ctx context.Context, email string) error {
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil // Do not reveal that the email does not exist
		}
		return err
	}

	rawToken, hashedToken, err := s.cryptoManager.GenerateResetPasswordToken()
	if err != nil {
		return err
	}

	_ = s.queries.DeleteUserPasswordResetsByUserID(ctx, user.ID)

	expiresAt := time.Now().Add(15 * time.Minute)
	_, err = s.queries.InsertUserPasswordReset(ctx, db.InsertUserPasswordResetParams{
		UserID:    user.ID,
		TokenHash: hashedToken,
		ExpiresAt: pgtype.Timestamptz{
			Time:  expiresAt,
			Valid: true,
		},
	})
	if err != nil {
		return err
	}

	go func() {
		bgCtx := context.Background()

		body := fmt.Sprintf(consts.EmailResetPasswordBody, rawToken, rawToken)
		if err := s.emailManager.SendEmailResetPassword(bgCtx, user.Email, consts.EmailResetPasswordSubject, body); err != nil {
			// Log the error, but don't return it since this is a background operation
			fmt.Printf("Failed to send reset password email: %v\n", err)
		}
	}()

	return nil
}

func (s *DefaultAuthenticationService) ResetPassword(ctx context.Context, token, newPassword string) error {
	tokenHash := s.cryptoManager.HashToken(token)

	resetToken, err := s.queries.GetUserPasswordResetByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return consts.ErrResetPasswordTokenInvalid
		}
		return err
	}

	if time.Now().After(resetToken.ExpiresAt.Time) {
		_ = s.queries.DeleteUserPasswordResetsByUserID(ctx, resetToken.UserID)
		return consts.ErrResetPasswordTokenExpired
	}

	hashedPassword, err := s.cryptoManager.EncryptPassword(newPassword)
	if err != nil {
		return err
	}

	err = s.queries.UpdateUserPassword(ctx, db.UpdateUserPasswordParams{
		ID:           resetToken.UserID,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		return err
	}

	_ = s.queries.DeleteUserPasswordResetsByUserID(ctx, resetToken.UserID)
	_ = s.queries.DeleteUserRefreshTokensByUserID(ctx, resetToken.UserID)

	return nil
}

func (s *DefaultAuthenticationService) ValidateAccessToken(ctx context.Context, accessToken string) (*dto.Claims, error) {
	claims, err := s.cryptoManager.ParseJwtToken(accessToken)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
