package service

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/google/uuid"
	"github.com/Sakshamyadav19/emailtracker/config"
)

type EmailData struct {
	To       []string
	Cc       []string
	Subject  string
	Body     string
	IsHTML   bool
	Tracking map[string]string // Maps email addresses to tracking IDs
}

func SendEmail(cfg *config.Config, emailData EmailData) error {
	toRecipients := strings.Join(emailData.To, ", ")
	ccRecipients := strings.Join(emailData.Cc, ", ")

	headers := map[string]string{
		"From":    cfg.AuthEmail,
		"To":      toRecipients,
		"Cc":      ccRecipients,
		"Subject": emailData.Subject,
	}

	var emailBody strings.Builder
	for key, value := range headers {
		emailBody.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	if emailData.IsHTML {
		emailBody.WriteString("MIME-Version: 1.0\r\n")
		emailBody.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n")
	} else {
		emailBody.WriteString("Content-Type: text/plain; charset=\"UTF-8\"\r\n\r\n")
	}

	for trackingID := range emailData.Tracking {
		trackingPixel := fmt.Sprintf(`<img src=\"%s/track/%s\" style=\"display:none;\" alt=\"\">`, cfg.BaseURL, trackingID)
		emailData.Body += trackingPixel
	}

	emailBody.WriteString(emailData.Body + "\r\n")

	recipients := append(emailData.To, emailData.Cc...)

	auth := smtp.PlainAuth("", cfg.AuthEmail, cfg.AuthPassword, cfg.SMTPHost)
	return smtp.SendMail(cfg.SMTPHost+":"+cfg.SMTPPort, auth, cfg.AuthEmail, recipients, []byte(emailBody.String()))
}

func GenerateTrackingIDs(recipients []string) map[string]string {
	tracking := make(map[string]string)
	for _, recipient := range recipients {
		tracking[recipient] = uuid.NewString()
	}
	return tracking
}
