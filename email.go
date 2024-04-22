package main

import (
	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"log/slog"
	"os"
)

func sendEmail(toName, toEmail, fromName, subject, message string) (*rest.Response, error) {
	from := mail.NewEmail(fromName, os.Getenv("FROM_EMAIL"))
	to := mail.NewEmail(toName, toEmail)

	// NATODO do we need to clean message
	email := mail.NewSingleEmail(from, subject, to, message, "")

	encryptedKey := os.Getenv("SENDGRID_API_KEY")
	decryptedKey, err := decrypt(encryptedKey)
	if err != nil {
		return nil, err
	}

	client := sendgrid.NewSendClient(decryptedKey)

	response, err := client.Send(email)
	if err != nil {
		return nil, err
	}

	slog.Debug("Return from sendgrid send", "response", response)

	return response, nil
}
