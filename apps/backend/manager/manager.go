package manager

type Managers struct {
	Crypto CryptoManager
	Email  EmailManager
	Secret SecretManager
}

func DefaultManagers() *Managers {
	secretManager := NewSecretManager()

	return &Managers{
		Crypto: NewCryptoManager(secretManager),
		Email:  NewEmailManager(secretManager),
		Secret: secretManager,
	}
}
