package manager

import (
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mennaverse/dotsync/apps/backend/db"
)

type TokenPair struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type CryptoManager interface {
	EncryptPassword(password string) (string, error)
	VerifyPassword(password, hashedPassword string) (bool, error)
	GenerateTokenPair(user db.User) (*TokenPair, error)
}

type DefaultCryptoManager struct {
	CryptoManager
}

func NewCryptoManager() CryptoManager {
	return &DefaultCryptoManager{}
}

func (c *DefaultCryptoManager) EncryptPassword(password string) (string, error) {
	hashedPassword, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}

func (c *DefaultCryptoManager) VerifyPassword(password, hashedPassword string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hashedPassword)
	if err != nil {
		return false, err
	}
	return match, nil
}

func (c *DefaultCryptoManager) GenerateTokenPair(user db.User) (*TokenPair, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
	})
	signedToken, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return nil, err
	}
	return &TokenPair{
		AccessToken: signedToken,
		ExpiresAt:   time.Now().Add(time.Hour), // Set the expiration time as needed
	}, nil
}
