package service

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/google/uuid"
)

type EmailData struct {
	To       []string
	Cc       []string
	Subject  string
	Body     string
	IsHTML   bool
	Tracking map[string]string // Maps email addresses to tracking IDs
}

func SendEmail(smtpHost, smtpPort, authEmail, authPassword string, emailData EmailData) error {
	toRecipients := strings.Join(emailData.To, ", ")
	ccRecipients := strings.Join(emailData.Cc, ", ")

	headers := map[string]string{
		"From":    authEmail,
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

	for recipient := range emailData.Tracking {
		fmt.Printf(`http://localhost:8080/track/%s`, emailData.Tracking[recipient])
		trackingPixel := fmt.Sprintf(`<img src="http://localhost:8080/track/%s" style="display:none;" alt="">`, emailData.Tracking[recipient])
		emailData.Body += trackingPixel
	}

	emailBody.WriteString(emailData.Body + "\r\n")

	recipients := append(emailData.To, emailData.Cc...)

	fmt.Println("I am inside")

	auth := smtp.PlainAuth("", authEmail, authPassword, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, authEmail, recipients, []byte(emailBody.String()))
}

func GenerateTrackingIDs(recipients []string) map[string]string {
	tracking := make(map[string]string)
	for _, recipient := range recipients {
		tracking[recipient] = uuid.NewString()
	}
	return tracking
}
