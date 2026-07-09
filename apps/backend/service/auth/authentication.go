package auth

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mennaverse/dotsync/apps/backend/db"
	"github.com/mennaverse/dotsync/apps/backend/manager"
)

type AuthenticationService interface {
	Register(ctx context.Context, req RegisterRequest) (*db.User, error)
	Login(ctx context.Context, req LoginRequest) (*manager.TokenPair, error)
	Logout(ctx context.Context, refreshToken string) error
	RefreshToken(ctx context.Context, refreshToken string) (*manager.TokenPair, error)

	VerifyEmail(ctx context.Context, token string) error
	ResendVerificationEmail(ctx context.Context, email string) error
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token string, newPassword string) error

	ValidateAccessToken(ctx context.Context, accessToken string) (*Claims, error)
}

type DefaultAuthenticationService struct {
	AuthenticationService
	connection    *pgx.Tx
	cryptoManager manager.CryptoManager
}

func NewAuthenticationService(conn *pgx.Tx, cryptoManager manager.CryptoManager) AuthenticationService {
	return &DefaultAuthenticationService{
		connection:    conn,
		cryptoManager: cryptoManager,
	}
}

func (s *DefaultAuthenticationService) Register(ctx context.Context, req RegisterRequest) (*db.User, error) {
	queries := db.New(*s.connection)

	hashedPassword, err := s.cryptoManager.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user, err := queries.InsertUser(ctx, db.InsertUserParams{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *DefaultAuthenticationService) Login(ctx context.Context, req LoginRequest) (*manager.TokenPair, error) {
	queries := db.New(*s.connection)

	user, err := queries.FindUserByUsernameOrEmail(ctx, req.Login)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if user.Banned {
		return nil, ErrUserBanned
	}

	match, err := s.cryptoManager.VerifyPassword(req.Password, user.PasswordHash)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, ErrInvalidCredentials
	}

	// Generate token pair (access and refresh tokens)
	tokenPair, err := s.cryptoManager.GenerateTokenPair(user)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}
