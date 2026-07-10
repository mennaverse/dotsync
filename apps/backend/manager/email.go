package manager

import (
	"context"

	"github.com/wneessen/go-mail"
)

type EmailManager interface {
	SendEmailVerification(ctx context.Context, to string, subject string, body string) error
	SendEmailResetPassword(ctx context.Context, to string, subject string, body string) error
	SendEmail(ctx context.Context, to string, subject string, body string) error
}

type DefaultEmailManager struct {
	EmailManager
	secretManager SecretManager
}

func NewEmailManager(secretManager SecretManager) EmailManager {
	return &DefaultEmailManager{
		secretManager: secretManager,
	}
}

type EmailConfig struct {
	SMTPEnabled  bool
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

func (e *DefaultEmailManager) SendEmailVerification(ctx context.Context, to string, subject string, body string) error {
	return e.SendEmail(ctx, to, subject, body)
}

func (e *DefaultEmailManager) SendEmailResetPassword(ctx context.Context, to string, subject string, body string) error {
	return e.SendEmail(ctx, to, subject, body)
}

func (e *DefaultEmailManager) SendEmail(ctx context.Context, to string, subject string, body string) error {
	c, err := e.getEmailConfig()
	if err != nil {
		return err
	}

	message := mail.NewMsg()

	if err := message.From(c.SMTPUsername); err != nil {
		return err
	}

	if err := message.To(to); err != nil {
		return err
	}

	message.Subject(subject)
	message.SetBodyString(mail.TypeTextHTML, body)

	client, err := mail.NewClient(
		c.SMTPHost,
		mail.WithPort(c.SMTPPort),
		mail.WithUsername(c.SMTPUsername),
		mail.WithPassword(c.SMTPPassword),
		mail.WithTLSPolicy(mail.TLSOpportunistic),
	)
	if err != nil {
		return err
	}

	if err := client.DialAndSendWithContext(ctx, message); err != nil {
		return err
	}
	defer client.Close()

	return nil
}

func (e *DefaultEmailManager) getEmailConfig() (*EmailConfig, error) {
	smtpEnabled, err := e.secretManager.GetSmtpEnabled()
	if err != nil {
		return nil, err
	}

	smtpHost, err := e.secretManager.GetSmtpHost()
	if err != nil {
		return nil, err
	}

	smtpPort, err := e.secretManager.GetSmtpPort()
	if err != nil {
		return nil, err
	}

	smtpUsername, err := e.secretManager.GetSmtpUsername()
	if err != nil {
		return nil, err
	}

	smtpPassword, err := e.secretManager.GetSmtpPassword()
	if err != nil {
		return nil, err
	}

	return &EmailConfig{
		SMTPEnabled:  smtpEnabled,
		SMTPHost:     smtpHost,
		SMTPPort:     smtpPort,
		SMTPUsername: smtpUsername,
		SMTPPassword: smtpPassword,
	}, nil
}
