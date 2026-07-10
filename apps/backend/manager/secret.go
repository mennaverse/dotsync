package manager

import (
	"os"
	"strconv"

	"github.com/mennaverse/dotsync/apps/backend/consts"
	"github.com/spf13/viper"
)

type SecretManager interface {
	GetFrontendURL() (string, error)
	GetDatabaseDSN() (string, error)
	GetEmailVerificationEnabled() (bool, error)
	GetSmtpEnabled() (bool, error)
	GetSmtpHost() (string, error)
	GetSmtpPort() (int, error)
	GetSmtpUsername() (string, error)
	GetSmtpPassword() (string, error)
	GetPasswordSecret() (string, error)
}

type DefaultSecretManager struct {
	SecretManager
}

func NewSecretManager() SecretManager {
	return &DefaultSecretManager{}
}

func (s *DefaultSecretManager) GetFrontendURL() (string, error) {
	secret := viper.Get(consts.FrontendURLEnvVar)
	if secret == nil {
		return "", os.ErrNotExist
	}

	return secret.(string), nil
}

func (s *DefaultSecretManager) GetDatabaseDSN() (string, error) {
	secret := viper.Get(consts.DatabaseDSNEnvVar)
	if secret == nil {
		return "", os.ErrNotExist
	}

	return secret.(string), nil
}
func (s *DefaultSecretManager) GetEmailVerificationEnabled() (bool, error) {
	secret := viper.GetBool(consts.EmailVerificationEnabledEnvVar)
	return secret, nil
}

func (s *DefaultSecretManager) GetSmtpEnabled() (bool, error) {
	secret := viper.GetBool(consts.SmtpEnabledEnvVar)
	return secret, nil
}

func (s *DefaultSecretManager) GetSmtpHost() (string, error) {
	secret := viper.Get(consts.SmtpHostEnvVar)
	if secret == nil {
		return "", os.ErrNotExist
	}

	return secret.(string), nil
}

func (s *DefaultSecretManager) GetSmtpPort() (int, error) {
	secret := viper.Get(consts.SmtpPortEnvVar)
	if secret == nil {
		return 0, os.ErrNotExist
	}

	// Convert the secret to int if it's not already
	switch v := secret.(type) {
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case string:
		port, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return port, nil
	default:
		return 0, os.ErrInvalid
	}
}

func (s *DefaultSecretManager) GetSmtpUsername() (string, error) {
	secret := viper.Get(consts.SmtpUsernameEnvVar)
	if secret == nil {
		return "", os.ErrNotExist
	}

	return secret.(string), nil
}

func (s *DefaultSecretManager) GetSmtpPassword() (string, error) {
	secret := viper.Get(consts.SmtpPasswordEnvVar)
	if secret == nil {
		return "", os.ErrNotExist
	}

	return secret.(string), nil
}

func (s *DefaultSecretManager) GetPasswordSecret() (string, error) {
	secret := viper.Get(consts.PasswordSecretEnvVar)
	if secret == nil {
		return "", os.ErrNotExist
	}

	return secret.(string), nil
}
