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
	allRecipients := append(emailData.To, emailData.Cc...)
	emailData.Tracking = GenerateTrackingIDs(allRecipients)

	auth := smtp.PlainAuth("", cfg.AuthEmail, cfg.AuthPassword, cfg.SMTPHost)

	// Send a personalized email to each recipient
	for _, recipient := range allRecipients {
		headers := map[string]string{
			"From":    cfg.AuthEmail,
			"To":      recipient,
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

		// Add the personalized tracking pixel
		trackingID := emailData.Tracking[recipient]
		trackingPixel := fmt.Sprintf(`<img src="%s/track/%s" style="display:none;" alt="">`, cfg.BaseURL, trackingID)

		// Build the email body
		personalizedBody := emailData.Body + trackingPixel
		emailBody.WriteString(personalizedBody + "\r\n")

		// Send the email
		err := smtp.SendMail(cfg.SMTPHost+":"+cfg.SMTPPort, auth, cfg.AuthEmail, []string{recipient}, []byte(emailBody.String()))
		if err != nil {
			return fmt.Errorf("failed to send email to %s: %w", recipient, err)
		}
	}

	return nil
}

func GenerateTrackingIDs(recipients []string) map[string]string {
	tracking := make(map[string]string)
	for _, recipient := range recipients {
		tracking[recipient] = uuid.NewString()
	}
	return tracking
}
