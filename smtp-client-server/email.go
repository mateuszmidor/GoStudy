package main

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/pkg/errors"
	"github.com/wneessen/go-mail"
)

// Error definitions
var (
	ErrDependencyFailed = errors.New("dependency failed")
	ErrCorrupted        = errors.New("corrupted data")
)

// ServerConfig holds the SMTP server configuration to be used for sending emails
type ServerConfig struct {
	Host       string
	Port       int
	UseTLS     bool          // if true, TLS will be used for the connection
	VerifyCert bool          // if true, the server's certificate will be verified. Set it to false for self-signed certificates
	Timeout    time.Duration // server connection timeout; set to 0 to use default timeout
}

// AuthMode is a function that returns the authentication options for the SMTP server
type AuthMode = func() []mail.Option

// AuthNone adds support for non-authenticated email sending, e.g. localhost Postfix server with default configuration can accept this
func AuthNone() AuthMode {
	return func() []mail.Option {
		return []mail.Option{}
	}
}

// AuthPlainUserPass adds support for SMTP AUTH PLAIN
func AuthPlainUserPass(user, pass string) AuthMode {
	return func() []mail.Option {
		return []mail.Option{
			mail.WithSMTPAuth(mail.SMTPAuthPlain),
			mail.WithUsername(user),
			mail.WithPassword(pass),
		}
	}
}

// AuthLoginUserPass adds support for SMTP AUTH LOGIN
func AuthLoginUserPass(user, pass string) AuthMode {
	return func() []mail.Option {
		return []mail.Option{
			mail.WithSMTPAuth(mail.SMTPAuthLogin),
			mail.WithUsername(user),
			mail.WithPassword(pass),
		}
	}
}

// SendEmail sends an email using specified SMTP server configuration and auth mode
func SendEmail(ctx context.Context, from, to, subject, body string, server ServerConfig, auth AuthMode) error {
	// prepare client config
	options := []mail.Option{
		mail.WithPort(server.Port),
	}

	// add timeout config
	if server.Timeout > 0 {
		options = append(options, mail.WithTimeout(server.Timeout))
	}

	// add TLS connection config
	if server.UseTLS {
		tlsConfig := &tls.Config{InsecureSkipVerify: !server.VerifyCert}
		options = append(options, mail.WithTLSPolicy(mail.TLSMandatory), mail.WithTLSConfig(tlsConfig))
	} else {
		options = append(options, mail.WithTLSPolicy(mail.NoTLS))
	}

	// add auth config
	options = append(options, auth()...)

	// create the client
	client, err := mail.NewClient(server.Host, options...)
	if err != nil {
		return errors.Wrapf(ErrDependencyFailed, "failed to create SMTP client for host %s:%d (TLS: %v): %s", server.Host, server.Port, server.UseTLS, err.Error())
	}

	// send the email
	return sendEmailWithClient(ctx, from, to, subject, body, client)
}

// sendEmailWithClient sends an email using the provided mail client
func sendEmailWithClient(ctx context.Context, from, to, subject, body string, client *mail.Client) error {
	m := mail.NewMsg()
	if err := m.From(from); err != nil {
		return errors.Wrapf(ErrCorrupted, "failed to set From address: %s", err.Error())
	}
	if err := m.To(to); err != nil {
		return errors.Wrapf(ErrCorrupted, "failed to set To address: %s", err.Error())
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextPlain, body)

	if err := client.DialAndSendWithContext(ctx, m); err != nil {
		return errors.Wrapf(ErrDependencyFailed, "failed to send mail: %s", err.Error())
	}

	return nil
}
