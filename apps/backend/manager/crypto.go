package manager

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mennaverse/dotsync/apps/backend/consts"
	"github.com/mennaverse/dotsync/apps/backend/db"
	"github.com/mennaverse/dotsync/apps/backend/dto"
)

type CryptoManager interface {
	EncryptPassword(password string) (string, error)
	VerifyPassword(password, hashedPassword string) (bool, error)
	ParseJwtToken(tokenString string) (*dto.Claims, error)
	GenerateJwtToken(user *db.User) (string, error)
	GenerateRefreshToken() (string, string, error)
	GenerateVerificationToken() (string, string, error)
	GenerateResetPasswordToken() (string, string, error)
	HashToken(token string) string
}

type DefaultCryptoManager struct {
	CryptoManager
	secretManager SecretManager
}

func NewCryptoManager(secretManager SecretManager) CryptoManager {
	return &DefaultCryptoManager{
		secretManager: secretManager,
	}
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

func (c *DefaultCryptoManager) ParseJwtToken(tokenString string) (*dto.Claims, error) {
	secret, err := c.secretManager.GetPasswordSecret()
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, consts.ErrTokenInvalid
}

func (c *DefaultCryptoManager) GenerateJwtToken(user *db.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID.String(),
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})

	secret, err := c.secretManager.GetPasswordSecret()
	if err != nil {
		return "", err
	}

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (c *DefaultCryptoManager) GenerateRefreshToken() (string, string, error) {
	rawToken, hashedToken, err := c.generateSecureToken()
	if err != nil {
		return "", "", err
	}

	return rawToken, hashedToken, nil
}

func (c *DefaultCryptoManager) GenerateVerificationToken() (string, string, error) {
	rawToken, hashedToken, err := c.generateSecureToken()
	if err != nil {
		return "", "", err
	}

	return rawToken, hashedToken, nil
}

func (c *DefaultCryptoManager) GenerateResetPasswordToken() (string, string, error) {
	rawToken, hashedToken, err := c.generateSecureToken()
	if err != nil {
		return "", "", err
	}

	return rawToken, hashedToken, nil
}

func (c *DefaultCryptoManager) HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func (c *DefaultCryptoManager) generateSecureToken() (rawToken string, hashedToken string, err error) {
	tokenBytes := make([]byte, 32)
	if _, err = rand.Read(tokenBytes); err != nil {
		return "", "", err
	}

	rawToken = hex.EncodeToString(tokenBytes)

	hashedToken = c.HashToken(rawToken)

	return rawToken, hashedToken, nil
}
