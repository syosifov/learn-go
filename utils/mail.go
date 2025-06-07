package utils

import (
	"crypto/tls"
	"fmt"
	"os"

	"gopkg.in/mail.v2"
)

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	FromName string
	FromAddr string
}

// LoadEmailConfig loads email configuration from environment variables
func LoadEmailConfig() EmailConfig {
	return EmailConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     getEnvAsInt("SMTP_PORT", 587),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
		FromName: os.Getenv("EMAIL_FROM_NAME"),
		FromAddr: os.Getenv("EMAIL_FROM_ADDR"),
	}
}

// Helper function to get an env var as int
func getEnvAsInt(name string, defaultVal int) int {
	valStr := os.Getenv(name)
	if valStr == "" {
		return defaultVal
	}
	val := 0
	fmt.Sscanf(valStr, "%d", &val)
	return val
}

// SendPasswordResetEmail sends a password reset email to the user
func SendPasswordResetEmail(to string, resetToken string, config EmailConfig) error {
	m := mail.NewMessage()

	// Set sender
	m.SetHeader("From", fmt.Sprintf("%s <%s>", config.FromName, config.FromAddr))

	// Set recipient
	m.SetHeader("To", to)

	// Set subject
	m.SetHeader("Subject", "Password Reset Request")

	// Set body with HTML content
	resetLink := fmt.Sprintf("https://yourapp.com/reset-password?token=%s", resetToken)
	body := fmt.Sprintf(`
        <h1>Password Reset</h1>
        <p>You have requested to reset your password. Please click the link below to set a new password:</p>
        <p><a href="%s">Reset Password</a></p>
        <p>If you did not request a password reset, please ignore this email.</p>
        <p>This link will expire in 30 minutes.</p>
    `, resetLink)

	m.SetBody("text/html", body)

	// Setup dialer with SMTP server settings
	d := mail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	return d.DialAndSend(m)
}
