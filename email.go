package main

import (
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"log/slog"
	"os"
)

func sendEmail(name, toEmail, subject, message string) error {
	from := mail.NewEmail("Nick", os.Getenv("FROM_EMAIL"))
	to := mail.NewEmail(name, toEmail)

	// NATODO do we need to clean message
	email := mail.NewSingleEmail(from, subject, to, message, "")

	encryptedKey := os.Getenv("SENDGRID_API_KEY")
	decryptedKey, err := decrypt(encryptedKey)
	if err != nil {
		return err
	}

	client := sendgrid.NewSendClient(decryptedKey)

	response, err := client.Send(email)
	if err != nil {
		return err
	}

	slog.Debug("email response: %v", response)

	return nil
}
