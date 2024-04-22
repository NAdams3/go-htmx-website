package main

import (
	"os"
	"testing"
)

func TestSendEmail(t *testing.T) {

	devEmail, err := decrypt(os.Getenv("DEV_EMAIL"))
	if err != nil {
		t.Error("error decrypting DEV_EMAIL")
	}

	expectedStatusCode := 202
	got, _ := sendEmail("Nick", devEmail, "test", "this is a test", "this is a test")
	if got.StatusCode != expectedStatusCode {
		t.Errorf("expected %v, got %v", expectedStatusCode, got)
	}
}
